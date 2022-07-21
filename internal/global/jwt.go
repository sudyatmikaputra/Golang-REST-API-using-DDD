package global

import (
	"context"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/medicplus-inc/medicplus-feedback/config"
)

func GenerateJWTToken(claims map[string]interface{}) (*string, error) {
	jwtauth.SetExpiry(claims, time.Now().Add(time.Hour*48))
	tokenAuth := jwtauth.New("HS256", []byte(config.GetValue(config.JWT_SECRET)), nil)

	_, jwtToken, err := tokenAuth.Encode(claims)
	if err != nil {
		return nil, err
	}

	return &jwtToken, nil
}

func GetClaimsFromContext(ctx context.Context) (map[string]interface{}, error) {
	_, claims, err := jwtauth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	return claims, nil
}
