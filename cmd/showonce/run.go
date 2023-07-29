package main

// See also: https://blog.logrocket.com/guide-to-grpc-gateway/

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/LeKovr/go-kit/config"

	"github.com/LeKovr/go-kit/logger"

	// importing implementation.
	app "github.com/LeKovr/showonce"
	// importing storage implementation.
	"github.com/LeKovr/showonce/static"
	storage "github.com/LeKovr/showonce/storage/cache"

	// importing generated stubs.
	gen "github.com/LeKovr/showonce/zgen/go/proto"
	"github.com/dopos/narra"
	"github.com/felixge/httpsnoop"
	"github.com/go-logr/logr"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

// Config holds all config vars.
type Config struct {
	Listen      string        `long:"listen" default:":8080" description:"Addr and port which server listens at"`
	ListenGRPC  string        `long:"listen_grpc" default:":8081" description:"Addr and port which GRPC pub server listens at"`
	Root        string        `long:"root" env:"ROOT" default:""  description:"Static files root directory"`
	PrivPrefix  string        `long:"priv" default:"/my/" description:"URI prefix for pages which requires auth"`
	GracePeriod time.Duration `long:"grace" default:"10s" description:"Stop grace period"`

	Logger     logger.Config  `group:"Logging Options" namespace:"log" env-namespace:"LOG"`
	AuthServer narra.Config   `group:"Auth Service Options" namespace:"as" env-namespace:"AS"`
	Storage    storage.Config `group:"Storage Options" namespace:"db" env-namespace:"DB"`
}

const (
	application = "showonce"
)

// Actual main.version value will be set at build time.
var version = "0.0-dev"

// Run app and exit via given exitFunc.
func Run(ctx context.Context, exitFunc func(code int)) {
	// Load config
	var cfg Config
	err := config.Open(&cfg)
	defer func() { config.Close(err, exitFunc) }()
	if err != nil {
		return
	}
	log := logger.New(cfg.Logger, nil)
	log.Info(application, "version", version)
	ctx = logr.NewContext(ctx, log)
	db := storage.New(cfg.Storage)

	Interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Inject logger.
		return handler(logr.NewContext(ctx, log), req)
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
	grpcPrivServer := grpc.NewServer(opts...) // TODO: UnaryInterceptor: md["user"]!=""
	gen.RegisterPrivateServiceServer(grpcPrivServer, app.NewPrivateService(db))
	mux := runtime.NewServeMux(
		runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
			userName := request.Header.Get(cfg.AuthServer.UserHeader)
			log.Info("Got GRPC", "user", userName)
			md := metadata.Pairs("user", userName)
			return md
		}),
	)

	// static pages server
	hfs, _ := static.New(cfg.Root)
	fileServer := http.FileServer(hfs)
	muxHTTP := http.NewServeMux()
	muxHTTP.Handle("/", fileServer)

	// Setup OAuth
	cfg.AuthServer.Do401 = true // we need redirect instead status 401
	auth := narra.New(&cfg.AuthServer)
	auth.SetupRoutes(muxHTTP, cfg.PrivPrefix)
	re := regexp.MustCompile("^" + cfg.PrivPrefix)
	hh := auth.ProtectMiddleware(withGW(mux, muxPub, muxHTTP), re)

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	err = gen.RegisterPublicServiceHandlerFromEndpoint(ctx, muxPub, cfg.ListenGRPC, dialOpts)
	if err != nil {
		return
	}
	clientAddr := chooseClientAddr(cfg.Listen)
	err = gen.RegisterPrivateServiceHandlerFromEndpoint(ctx, mux, clientAddr, dialOpts)
	if err != nil {
		return
	}
	// Creating a normal HTTP server
	srv := &http.Server{
		Addr:    cfg.Listen,
		Handler: withReqLogger(hh),
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// creating a listener for server
	var listenerPub net.Listener
	listenerPub, err = net.Listen("tcp", cfg.ListenGRPC)
	if err != nil {
		return
	}
	// creating a listener for server
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

	// start servers
	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return srv.Serve(httpL)
	})
	g.Go(func() error {
		return grpcPrivServer.Serve(grpcL)
	})
	g.Go(func() error {
		log.V(1).Info("Start GRPC service", "addr", cfg.ListenGRPC)
		return grpcPubServer.Serve(listenerPub)
	})
	g.Go(func() error {
		log.V(1).Info("Start HTTP service", "addr", cfg.Listen)
		return m.Serve()
	})
	g.Go(func() error {
		<-gCtx.Done()
		log.V(1).Info("Shutdown")
		stop()
		timedCtx, cancel := context.WithTimeout(ctx, cfg.GracePeriod)
		defer cancel()
		grpcPrivServer.GracefulStop()
		grpcPubServer.GracefulStop()
		return srv.Shutdown(timedCtx)
	})
	if er := g.Wait(); er != nil && !errors.Is(er, net.ErrClosed) {
		err = er
	}
	log.Info("Exit")
}

// withReqLogger prints HTTP request log.
func withReqLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		m := httpsnoop.CaptureMetrics(handler, writer, request)
		fmt.Printf("http[%d]-- %s -- %s\n", m.Code, m.Duration, request.RequestURI)
	})
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

/*

type JAST interface {
  SetField(name string, value interface{}) error
  SetFields(name string, values []interface{}) error
}

func Setup(options ..Option) (JAST, error) {
}

main() {

app := jast.Setup(cfg).
  Logger(log).
  UseHTTP(true).
  GRPC("/app",pubService).
  GRPC("/my/app",privService).
  Static(openapi).
  Static(openapiUI)
//...
err = app.Serve()


)
  if err == nil {
    app.Run(exitFunc)
  }
}
*/
