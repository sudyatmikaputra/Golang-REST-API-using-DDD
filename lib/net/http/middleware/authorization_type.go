package middleware

import (
	"context"

	"github.com/google/uuid"
)

type AuthorizationKey string

const (
	AuthorizationUserID    AuthorizationKey = "UserId"
	AuthorizationSessionID AuthorizationKey = "SessionId"
)

func GetAuthorizedUserID(ctx context.Context) uuid.UUID {
	if id := ctx.Value(AuthorizationUserID); id != nil {
		return id.(uuid.UUID)
	}

	return uuid.Nil
}
