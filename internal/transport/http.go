package transport

import (
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/lib/net/http/middleware"

	"github.com/go-chi/chi"
	kitHttp "github.com/go-kit/kit/transport/http"
)

func CompileRoute(
	r *chi.Mux,
	opts []kitHttp.ServerOption,
) http.Handler {

	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthorizedOnly)

	})

	return r
}
