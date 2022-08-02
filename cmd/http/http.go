package http

import (
	"net/http"
	"time"

	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/medicplus-inc/medicplus-kit/net/http/encoding"
	"github.com/rs/cors"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	feedback "github.com/medicplus-inc/medicplus-feedback/internal/transport"
	medicplusMiddleware "github.com/medicplus-inc/medicplus-feedback/lib/net/http/middleware"
)

func CompileRoute(
	r *chi.Mux,
) http.Handler {
	opts := []kitHttp.ServerOption{
		kitHttp.ServerErrorEncoder(encoding.EncodeError),
	}

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.Compress(5))

	r.Use(middleware.Timeout(60 * time.Second))

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Access-Token", "X-Requested-With"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(cors.Handler)
	r.Use(medicplusMiddleware.Authenticate)

	feedback.CompileRoute(r, opts)

	return r
}
