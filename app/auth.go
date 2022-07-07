package app

import (
	"net/http"
	"strings"

	"github.com/LeKovr/go-kit/oauth2"

	"github.com/go-logr/logr"
)

type Authenticator struct {
	config oauth2.Config
	prefix string

	// New creates service
	service *oauth2.Service
}

func NewAuthenticator(cfg oauth2.Config, prefix string) Authenticator {
	cfg.Do401 = true // we need redirect
	as := oauth2.New(cfg)
	return Authenticator{config: cfg, prefix: prefix, service: as}
}

func (au Authenticator) SetupRoutes(mux *http.ServeMux) {
	mux.Handle("/auth", au.service.AuthHandler())
	mux.Handle(au.config.CallBackURL, au.service.Stage2Handler())
	mux.Handle("/401/", au.service.Stage1Handler())
	mux.Handle("/my/logout", au.service.LogoutHandler())
}

func (au Authenticator) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logr.FromContextOrDiscard(r.Context())
		if strings.HasPrefix(r.URL.Path, au.prefix) {
			log.V(1).Info("Page is protected", "url", r.URL.Path)
			// TODO: do we need it?
			scheme := "http"
			if r.TLS != nil {
				scheme += "s"
			}
			r.Header.Set("X-Forwarded-Proto", scheme)
			r.Header.Set("X-Forwarded-Host", r.Host)
			r.Header.Set("X-Forwarded-Uri", r.RequestURI)

			if !au.service.AuthIsOK(w, r) {
				return
			}
			w.Header().Set("Last-Modified", "")
		}
		next.ServeHTTP(w, r)
	})
}
