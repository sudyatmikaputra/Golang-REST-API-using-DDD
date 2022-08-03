package middleware

import (
	"errors"
	"log"
	"net/http"

	kitErr "github.com/medicplus-inc/medicplus-kit/error"

	"github.com/medicplus-inc/medicplus-feedback/cmd/container"
	"github.com/medicplus-inc/medicplus-kit/net/authentication/csrf"
	"github.com/medicplus-inc/medicplus-kit/net/http/encoding"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authentication := csrf.NewCSRFService(container.Injector().Repository.Redis())
		if r.Method == http.MethodGet {
			csrfToken, err := authentication.GenerateToken()
			if err != nil {
				log.Println("Errr while generating CSRF Token: ", err)
			}
			log.Println("Token", csrfToken)

			w.Header().Add("X-CSRF-Token", csrfToken)
		} else {
			if !authentication.ValidateToken(r.Header.Get("X-CSRF-Token")) {
				encoding.EncodeError(r.Context(), kitErr.New(
					errors.New("access-denied"),
					http.StatusForbidden,
					"Access Denied",
				), w)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
