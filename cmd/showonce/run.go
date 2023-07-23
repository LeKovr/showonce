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

	"github.com/felixge/httpsnoop"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"

	// "google.golang.org/grpc/credentials/insecure"

	"github.com/LeKovr/go-kit/config"
	"github.com/LeKovr/go-kit/logger"
	"github.com/dopos/narra"

	// importing generated stubs
	gen "github.com/LeKovr/showonce/zgen/go/proto"
	// importing implementation
	app "github.com/LeKovr/showonce"
	"github.com/LeKovr/showonce/static"
)

// Config holds all config vars
type Config struct {
	Listen      string        `long:"listen" default:":8080" description:"Addr and port which server listens at"`
	Root        string        `long:"root" env:"ROOT" default:""  description:"Static files root directory"`
	PrivPrefix  string        `long:"priv" default:"/my/" description:"URI prefix for pages which requires auth"`
	GracePeriod time.Duration `long:"grace" default:"1m" description:"Stop grace period"`

	Logger     logger.Config     `group:"Logging Options" namespace:"log" env-namespace:"LOG"`
	AuthServer narra.Config      `group:"Auth Service Options" namespace:"as" env-namespace:"AS"`
	Storage    app.StorageConfig `group:"Storage Options" namespace:"db" env-namespace:"DB"`
}

const (
	application = "showonce"
)

// Actual main.version value will be set at build time
var version = "0.0-dev"

// Run app and exit via given exitFunc
func Run(exitFunc func(code int)) {
	// Load config
	var cfg Config
	err := config.Open(&cfg)
	defer func() { config.Close(err, exitFunc) }()
	if err != nil {
		return
	}
	log := logger.New(cfg.Logger, nil)
	log.Info(application, "version", version)

	db := app.NewStorage(cfg.Storage)

	// create new gRPC server
	grpcSever := grpc.NewServer()
	// Register reflection service on gRPC server.
	reflection.Register(grpcSever)
	// register the PublicServiceServerImpl on the gRPC server
	gen.RegisterPublicServiceServer(grpcSever, app.NewPublicService(db))
	// creating mux for gRPC gateway. This will multiplex or route request different gRPC service
	mux := runtime.NewServeMux(
		// handle incoming headers
		runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
			header := request.Header.Get("Authorization")
			// send all the headers received from the client
			md := metadata.Pairs("auth", header)
			return md
		}),
	)
	clientAddr := chooseClientAddr(cfg.Listen)
	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	// setting up a dial up for gRPC service by specifying endpoint/target url
	err = gen.RegisterPublicServiceHandlerFromEndpoint(context.Background(), mux, clientAddr, dialOpts)
	if err != nil {
		return
	}

	// static pages server
	hfs, _ := static.New(cfg.Root)
	fileServer := http.FileServer(hfs)
	muxH := http.NewServeMux()
	muxH.Handle("/", fileServer)

	//	api := app.NewAPIService(cfg.AuthServer.UserHeader, db)
	//	api.SetupRoutes(mux, cfg.PrivPrefix)

	cfg.AuthServer.Do401 = true // we need redirect instead status 401
	// ? cfg.AuthServer.CallBackURL
	auth := narra.New(&cfg.AuthServer)
	auth.SetupRoutes(muxH, cfg.PrivPrefix)
	re := regexp.MustCompile("^" + cfg.PrivPrefix)
	hh := auth.ProtectMiddleware(muxH, re)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	ctx = logger.NewContext(ctx, log)

	// Creating a normal HTTP server
	srv := &http.Server{
		Addr:    cfg.Listen,
		Handler: withReqLogger(withGW(mux, hh)),
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}

	// creating a listener for server
	var l net.Listener
	l, err = net.Listen("tcp", cfg.Listen)
	if err != nil {
		return
	}
	m := cmux.New(l)

	// a different listener for HTTP1
	httpL := m.Match(cmux.HTTP1Fast())

	// a different listener for HTTP2 since gRPC uses HTTP2
	grpcL := m.Match(cmux.HTTP2())
	// start server

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return srv.Serve(httpL)
	})
	g.Go(func() error {
		return grpcSever.Serve(grpcL)
	})
	g.Go(func() error {
		log.V(1).Info("Start", "addr", cfg.Listen)
		return m.Serve()
	})
	g.Go(func() error {
		<-gCtx.Done()
		log.V(1).Info("Shutdown")
		stop()
		timedCtx, cancel := context.WithTimeout(ctx, cfg.GracePeriod)
		defer cancel()
		return srv.Shutdown(timedCtx)
	})
	if er := g.Wait(); er != nil && !errors.Is(er, net.ErrClosed) { //er != http.ErrServerClosed {
		err = er
	}
	log.Info("Exit")
}

// withReqLogger prints HTTP request log
func withReqLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		m := httpsnoop.CaptureMetrics(handler, writer, request)
		fmt.Printf("http[%d]-- %s -- %s\n", m.Code, m.Duration, request.RequestURI)
	})
}

func withGW(gwmux *runtime.ServeMux, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api") {
			gwmux.ServeHTTP(w, r)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

// chooseClientAddr chooses localhost if server listens any ip
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

app,err := jast.Setup(cfg,
  jast.logger(),
  jast.UseHTTP(true),
  jast.GRPC("/app",pubService),
  jast.GRPC("/my/app",privService),
  jast.Static(openapi),
  jast.Static(openapiUI),
//...
  jast.Handlers()


)
  if err == nil {
    app.Run(exitFunc)
  }
}
*/
