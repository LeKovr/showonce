package main

// See also: https://blog.logrocket.com/guide-to-grpc-gateway/

import (
	"context"
	"fmt"
	"io/fs"
	"log/slog"
	"net"
	"net/http"
	"regexp"
	"strings"

	"github.com/LeKovr/go-kit/config"
	"github.com/LeKovr/go-kit/server"
	"github.com/LeKovr/go-kit/slogger"
	"github.com/LeKovr/go-kit/ver"

	"github.com/dopos/narra"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"

	// importing implementation.
	app "github.com/LeKovr/showonce"
	// importing static files.
	"github.com/LeKovr/showonce/static"
	// importing storage implementation.
	storage "github.com/LeKovr/showonce/storage/cache"
	// importing generated stubs.
	gen "github.com/LeKovr/showonce/zgen/go/proto"
)

// Config holds all config vars.
type Config struct {
	Listen     string `default:":8080" description:"Addr and port which server listens at"          long:"listen"`
	ListenGRPC string `default:":8081" description:"Addr and port which GRPC pub server listens at" long:"listen_grpc"`
	Root       string `default:""      description:"Static files root directory"                    env:"ROOT"         long:"root"`
	HTMLPath   string `default:"html"  description:"Static site subdirectory"                       long:"html"`
	PrivPrefix string `default:"/my/"  description:"URI prefix for pages which requires auth"       long:"priv"`

	Logger     slogger.Config `env-namespace:"LOG" group:"Logging Options"      namespace:"log"`
	AuthServer narra.Config   `env-namespace:"AS"  group:"Auth Service Options" namespace:"as"`
	Storage    storage.Config `env-namespace:"DB"  group:"Storage Options"      namespace:"db"`
	Server     server.Config  `env-namespace:"SRV" group:"Server Options"       namespace:"srv"`
}

const (
	application = "showonce"
)

var (
	// App version, actual value will be set at build time.
	version = "0.0-dev"

	// Repository address, actual value will be set at build time.
	repo = "repo.git"
)

// Run app and exit via given exitFunc.
func Run(ctx context.Context, exitFunc func(code int)) {
	// Load config
	var cfg Config
	err := config.Open(&cfg)
	defer func() {
		if r := recover(); r != nil {
			slog.Error("Recovered panic", "err", r)
		}
		config.Close(err, exitFunc)
	}()
	if err != nil {
		return
	}
	err = slogger.Setup(cfg.Logger, nil)
	if err != nil {
		return
	}
	slog.Info(application, "version", version)
	go ver.Check(repo, version)
	db := storage.New(cfg.Storage)

	Interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Inject logger.
		return handler(ctx, req)
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(Interceptor),
	}

	// Public GRPC Service
	// Доступен извне, отдельный порт
	grpcPubServer := grpc.NewServer(opts...)
	gen.RegisterPublicServiceServer(grpcPubServer, app.NewPublicService(db))
	reflection.Register(grpcPubServer)
	muxPub := runtime.NewServeMux()
	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// Private GRPC Service
	// Доступен только через HTTP
	// Авторизацию делает HTTP Handler
	grpcPrivServer := grpc.NewServer(opts...) // TODO: UnaryInterceptor: md[app.MDUserKey]!=""
	gen.RegisterPrivateServiceServer(grpcPrivServer, app.NewPrivateService(db))
	mux := runtime.NewServeMux(
		runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
			userName := request.Header.Get(cfg.AuthServer.UserHeader)
			slog.Info("Got PrivGRPC", "user", userName)
			/*
				   	type key string
				   	const myCustomKey key = "key"
				   	func f(ctx context.Context) {
				       ctx = context.WithValue(ctx, myCustomKey, "foo")
					}
			*/

			md := metadata.Pairs(app.MDUserKey, userName)
			return md
		}),
	)

	// static pages server
	var root, hfs fs.FS
	if root, err = static.New(cfg.Root); err != nil {
		return
	}
	if hfs, err = fs.Sub(root, cfg.HTMLPath); err != nil {
		return
	}
	srv := server.New(cfg.Server).WithStatic(hfs).WithVersion(version)

	// Setup OAuth
	cfg.AuthServer.Do401 = true // we need redirect instead status 401
	auth := narra.New(&cfg.AuthServer)
	auth.SetupRoutes(srv.ServeMux(), cfg.PrivPrefix)
	re := regexp.MustCompile("^" + cfg.PrivPrefix)
	srv.Use(func(handler http.Handler) http.Handler {
		return auth.ProtectMiddleware(withGW(mux, muxPub, srv.ServeMux()), re)
	})
	err = gen.RegisterPublicServiceHandlerFromEndpoint(ctx, muxPub, cfg.ListenGRPC, dialOpts)
	if err != nil {
		return
	}
	clientAddr := chooseClientAddr(cfg.Listen)
	err = gen.RegisterPrivateServiceHandlerFromEndpoint(ctx, mux, clientAddr, dialOpts)
	if err != nil {
		return
	}

	// creating a listener for GRPC server
	var listenerPub net.Listener
	listenerPub, err = net.Listen("tcp", cfg.ListenGRPC)
	if err != nil {
		return
	}

	// creating a listener for HTTP server
	var listener net.Listener
	listener, err = net.Listen("tcp", cfg.Listen)
	if err != nil {
		return
	}
	m := cmux.New(listener)

	// a different listener for HTTP1
	httpL := m.Match(cmux.HTTP1Fast())

	// a different listener for HTTP2 since gRPC uses HTTP2
	grpcL := m.Match(cmux.HTTP2())

	srv.Use(cfg.Server.WithAccessLog)
	srv.WithListener(httpL)
	srv.WithShutdown(
		func(_ context.Context) error {
			grpcPrivServer.GracefulStop()
			grpcPubServer.GracefulStop()
			return nil
		})

	err = srv.Run(ctx,
		func(_ context.Context) error {
			return grpcPrivServer.Serve(grpcL)
		},
		func(_ context.Context) error {
			slog.Info("Start GRPC service", "addr", cfg.ListenGRPC)
			return grpcPubServer.Serve(listenerPub)
		},
		func(_ context.Context) error {
			slog.Info("Start HTTP service", "addr", cfg.Listen)
			return m.Serve()
		},
	)
}

func withGW(gwmux, gwmuxPub *runtime.ServeMux, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api") {
			gwmuxPub.ServeHTTP(w, r)
			return
		}
		if strings.HasPrefix(r.URL.Path, "/my/api") {
			gwmux.ServeHTTP(w, r)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

// chooseClientAddr chooses localhost if server listens any ip.
func chooseClientAddr(addr string) string {
	parts := strings.SplitN(addr, ":", 2)
	if parts[0] == "0.0.0.0" || parts[0] == "" {
		return fmt.Sprintf("%s:%s", "localhost", parts[1])
	}
	return addr
}
