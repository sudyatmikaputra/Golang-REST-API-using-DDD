package middleware

import (
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

func AdminVerifier(jwt *jwtauth.JWTAuth) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			token, err := jwtauth.VerifyRequest(jwt, r, jwtauth.TokenFromHeader, jwtauth.TokenFromCookie)
			if err == nil {
				claims := token.PrivateClaims()
				if claims["roles"] != "admin" {
					err = jwtauth.ErrUnauthorized
				}
			}
			ctx = jwtauth.NewContext(ctx, token, err)
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(hfn)
	}
}

func PatientVerifier(jwt *jwtauth.JWTAuth) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			token, err := jwtauth.VerifyRequest(jwt, r, jwtauth.TokenFromHeader, jwtauth.TokenFromCookie)
			if err == nil {
				claims := token.PrivateClaims()
				if claims["roles"] != "patient" {
					err = jwtauth.ErrUnauthorized
				}
			}
			ctx = jwtauth.NewContext(ctx, token, err)
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(hfn)
	}
}

func DoctorVerifier(jwt *jwtauth.JWTAuth) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			token, err := jwtauth.VerifyRequest(jwt, r, jwtauth.TokenFromHeader, jwtauth.TokenFromCookie)
			if err == nil {
				claims := token.PrivateClaims()
				if claims["roles"] != "doctor" {
					err = jwtauth.ErrUnauthorized
				}
			}
			ctx = jwtauth.NewContext(ctx, token, err)
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(hfn)
	}
}
