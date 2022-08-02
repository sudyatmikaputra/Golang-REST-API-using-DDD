package public

import (
	"github.com/google/uuid"
)

type FeedbackResponse struct {
	ID                uuid.UUID                 `json:"id"`
	FeedbackType      string                    `json:"feedback_type"`
	FeedbackToID      uuid.UUID                 `json:"feedback_to_id"`
	FeedbackFromID    uuid.UUID                 `json:"feedback_from_id"`
	FeedbackParameter FeedbackParameterResponse `json:"feedback_parameter"`
	FeedbackValue     int                       `json:"feedback_value"`
	Notes             string                    `json:"notes"`
}

// add for anonymous
type AnonymousFeedbackResponse struct {
	ID                uuid.UUID                 `json:"id"`
	FeedbackType      string                    `json:"feedback_type"`
	FeedbackToID      uuid.UUID                 `json:"feedback_to_id"`
	FeedbackParameter FeedbackParameterResponse `json:"feedback_parameter"`
	FeedbackValue     int                       `json:"feedback_value"`
	Notes             string                    `json:"notes"`
}

type CreateFeedbackRequest struct {
	FeedbackType   string    `json:"feedback_type" validate:"required"`
	FeedbackToID   uuid.UUID `json:"feedback_to_id" validate:"required"`
	FeedbackFromID uuid.UUID `json:"feedback_from_id" validate:"required"`
	FeedbackValue  int       `json:"feedback_value" validate:"required"`
	Notes          string    `json:"notes" validate:"required"`
}

type UpdateFeedbackRequest struct {
	ID            uuid.UUID `json:"id" validate:"required"`
	FeedbackValue int       `json:"feedback_value" validate:"required"`
	Notes         string    `json:"notes" validate:"required"`
}

type ListFeedbackRequest struct {
	Search         string    `qs:"search"` //values and notes
	Page           int       `qs:"page"`
	Limit          int       `qs:"limit"`
	FeedbackToID   uuid.UUID `qs:"feedback_to_id"`
	FeedbackFromID uuid.UUID `qs:"feedback_from_id"`
	FeedbackType   string    `qs:"feedback_type" validate:"required"`
	LanguageCode   string    `qs:"language_code" validate:"required,oneof=id en"`
}

type GetFeedbackRequest struct {
	ID           uuid.UUID `qs:"id" validate:"required"`
	LanguageCode string    `qs:"language_code" validate:"required,oneof=id en"`
}

type DeleteFeedbackRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}
