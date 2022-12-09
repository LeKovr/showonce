package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	"time"

	"github.com/LeKovr/go-kit/config"
	"github.com/LeKovr/go-kit/logger"
	"github.com/dopos/narra"
	"github.com/felixge/httpsnoop"
	"golang.org/x/sync/errgroup"

	"SELF/service"
	"SELF/static"
)

const (
	AppName = "showonce"
)

// Config holds all config vars
type Config struct {
	Listen      string        `long:"listen" default:":8080" description:"Addr and port which server listens at"`
	Root        string        `long:"root" env:"ROOT" default:""  description:"Static files root directory"`
	PrivPrefix  string        `long:"priv" default:"/my/" description:"URI prefix for pages which requires auth"`
	GracePeriod time.Duration `long:"grace" default:"1m" description:"Stop grace period"`

	Logger     logger.Config         `group:"Logging Options" namespace:"log" env-namespace:"LOG"`
	AuthServer narra.Config          `group:"Auth Service Options" namespace:"as" env-namespace:"AS"`
	Storage    service.StorageConfig `group:"Storage Options" namespace:"db" env-namespace:"DB"`
}

// Run app and exit via given exitFunc
func Run(version string, exitFunc func(code int)) {
	// Load config
	var cfg Config
	err := config.Open(&cfg)
	defer func() { config.Close(err, exitFunc) }()
	if err != nil {
		return
	}

	// Example: Disable oauth2 debug
	// oauth2.DL = 3

	// Setup logger
	log := logger.New(cfg.Logger, nil)

	log.Info(AppName, "version", version)

	// static pages server
	hfs, _ := static.New(cfg.Root)
	fileServer := http.FileServer(hfs)

	mux := http.NewServeMux()
	mux.Handle("/", fileServer)

	db := service.NewStorage(cfg.Storage)
	api := service.NewAPIService(cfg.AuthServer.UserHeader, db)
	api.SetupRoutes(mux, cfg.PrivPrefix)

	cfg.AuthServer.Do401 = true // we need redirect instead status 401
	// ? cfg.AuthServer.CallBackURL
	auth := narra.New(cfg.AuthServer)
	auth.SetupRoutes(mux, cfg.PrivPrefix)
	re := regexp.MustCompile("^" + cfg.PrivPrefix)
	hh := auth.ProtectMiddleware(mux, re)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	ctx = logger.NewContext(ctx, log)

	srv := &http.Server{
		Addr:    cfg.Listen,
		Handler: withReqLogger(hh),
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		log.V(1).Info("Start", "addr", cfg.Listen)
		return srv.ListenAndServe()
	})
	g.Go(func() error {
		<-gCtx.Done()
		log.V(1).Info("Shutdown")
		stop()
		timedCtx, cancel := context.WithTimeout(context.Background(), cfg.GracePeriod)
		defer cancel()
		return srv.Shutdown(timedCtx)
	})
	if er := g.Wait(); er != nil && er != http.ErrServerClosed {
		err = er
	}
	log.V(1).Info("Exit")
}

func LoggerMiddleware(ctx context.Context, next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// TODO: add client IP here
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}

// withReqLogger prints HTTP request log
func withReqLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		m := httpsnoop.CaptureMetrics(handler, writer, request)
		fmt.Printf("http[%d]-- %s -- %s\n", m.Code, m.Duration, request.RequestURI)
	})
}
