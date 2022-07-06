package app

import (
	"context"
	"net/http"
	"os"
	"time"

	"KIT/config"
	"KIT/logger"
	"KIT/oauth2"
	"SELF/service"
	"SELF/static"
)

const (
	AppName = "showonce"
)

// Config holds all config vars
type Config struct {
	config.ConfigWithVersion
	Listen      string        `long:"listen" default:":8080" description:"Addr and port which server listens at"`
	Root        string        `long:"root" env:"ROOT" default:""  description:"Static files root directory"`
	ExecTimeout time.Duration `long:"timeout" default:"1m" description:"Overall exec timeout"`

	AuthServer oauth2.Config         `group:"Auth Service Options" namespace:"as" env-namespace:"AS"`
	Storage    service.StorageConfig `group:"Storage Options" namespace:"db" env-namespace:"DB"`
}

// Run app and exit via given exitFunc
func Run(version string, exitFunc func(code int)) {
	// Load config
	var cfg Config
	err := config.OpenWithVersion(&cfg)
	defer func() { config.Close(exitFunc, err, os.Stdout, version) }()
	if err != nil {
		return
	}

	// Disable oauth2 debug
	oauth2.DL = 3

	// Setup logger
	log := logger.New(os.Stdout, cfg.Debug) //.With().Str("app", (AppName))

	// Setup conext
	ctx, cancel := context.WithTimeout(context.Background(), cfg.ExecTimeout)
	defer cancel()
	ctx = logger.NewContext(ctx, log)

	log.Info(AppName, "version ", version)

	// static pages server
	hfs, _ := static.New(cfg.Root)
	fileServer := http.FileServer(hfs)

	mux := http.NewServeMux()
	mux.Handle("/", fileServer)

	db := service.NewStorage(cfg.Storage)
	api := service.NewAPIService(cfg.AuthServer.UserHeader, db)
	api.SetupRoutes(mux)

	auth := NewAuthenticator(cfg.AuthServer, "/my/")
	auth.SetupRoutes(mux)
	m := auth.Middleware(mux)
	m = LoggerMiddleware(ctx, m)
	srv := &http.Server{
		Addr:    cfg.Listen,
		Handler: m,
	}

	err = srv.ListenAndServe()

	// Shutdown
	ctx.Done()
	log.V(1).Info("Exit")
}

func LoggerMiddleware(ctx context.Context, next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}
