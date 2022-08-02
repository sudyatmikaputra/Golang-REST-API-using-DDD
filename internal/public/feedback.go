package public

import (
	"github.com/google/uuid"
)

type FeedbackResponse struct {
	ID              uuid.UUID                 `json:"id"`
	FeedbackTo      string                    `json:"feedback_to"`
	FeedbackParamID uuid.UUID                 `json:"feedback_param_id"`
	FeedbackToID    uuid.UUID                 `json:"feedback_to_id"`
	FeedbackFromID  uuid.UUID                 `json:"feedback_from_id"`
	FeedbackParam   FeedbackParameterResponse `json:"feedback_param"`
	FeedbackValue   int                       `json:"feedback_value"`
	Notes           string                    `json:"notes"`
}

// add for anonymous
type AnonymousFeedbackResponse struct {
	ID              uuid.UUID                 `json:"id"`
	FeedbackTo      string                    `json:"feedback_to"`
	FeedbackParamID uuid.UUID                 `json:"feedback_param_id"`
	FeedbackToID    uuid.UUID                 `json:"feedback_to_id"`
	FeedbackParam   FeedbackParameterResponse `json:"feedback_param"`
	FeedbackValue   int                       `json:"feedback_value"`
	Notes           string                    `json:"notes"`
}

type CreateFeedbackRequest struct {
	FeedbackTo      string    `json:"feedback_to" validate:"required"`
	FeedbackParamID uuid.UUID `json:"feedback_param_id" validate:"required"`
	FeedbackToID    uuid.UUID `json:"feedback_to_id" validate:"required"`
	FeedbackFromID  uuid.UUID `json:"feedback_from_id" validate:"required"`
	FeedbackValue   int       `json:"feedback_value" validate:"required"`
	Notes           string    `json:"notes" validate:"required"`
}

type UpdateFeedbackRequest struct {
	ID            uuid.UUID `json:"id" validate:"required"`
	FeedbackValue int       `json:"feedback_value" validate:"required"`
	Notes         string    `json:"notes" validate:"required"`
}

type ListFeedbackRequest struct {
	Search          string    `qs:"search"` //values and notes
	Page            int       `qs:"page"`
	Limit           int       `qs:"limit"`
	FeedbackParamID uuid.UUID `qs:"feedback_param_id" validate:"required"`
	FeedbackToID    uuid.UUID `qs:"feedback_to_id"`
	FeedbackFromID  uuid.UUID `qs:"feedback_from_id"`
	FeedbackTo      string    `qs:"feedback_to" validate:"required"`
}

type GetFeedbackRequest struct {
	FeedbackID uuid.UUID `url_param:"feedback_id" validate:"required"`
}

type DeleteFeedbackRequest struct {
	FeedbackID uuid.UUID `json:"feedback_id" validate:"required"`
}
