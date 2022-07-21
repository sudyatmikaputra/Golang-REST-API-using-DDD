package transport

import (
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/config"
	"github.com/medicplus-inc/medicplus-feedback/lib/net/http/middleware"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
	kitHttp "github.com/go-kit/kit/transport/http"
)

func CompileRoute(
	r *chi.Mux,
	opts []kitHttp.ServerOption,
) http.Handler {
	jwtAuth := jwtauth.New("HS256", []byte(config.GetEnv(config.JWT_SECRET)), nil)

	// For general use
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(jwtAuth))
		r.Use(jwtauth.Authenticator)

	})

	// Accessed by admin only
	r.Group(func(r chi.Router) {
		r.Use(middleware.AdminVerifier(jwtAuth))

	})

	// Accessed by doctor only
	r.Group(func(r chi.Router) {
		r.Use(middleware.DoctorVerifier(jwtAuth))

	})

	// Accessed by patient only
	r.Group(func(r chi.Router) {
		r.Use(middleware.PatientVerifier(jwtAuth))

	})

	return r
}
