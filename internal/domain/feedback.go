package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/infrastructure/repository"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	"github.com/medicplus-inc/medicplus-kit/encoding"
)

type Feedback struct {
	ID                uuid.UUID             `json:"id"`
	FeedbackTo        internal.ReceiverType `json:"feedback_to"`
	FeedbackToID      uuid.UUID             `json:"feedback_to_id"`
	FeedbackFromID    uuid.UUID             `json:"feedback_from_id"`
	FeedbackParamID   uuid.UUID             `json:"feedback_param_id"`
	FeedbackParameter *FeedbackParameter    `json:"feedback_parameter"`
	FeedbackValue     int                   `json:"feedback_value"`
	Notes             string                `json:"notes"`
	CreatedAt         time.Time             `json:"created_at"`
	UpdatedAt         time.Time             `json:"updated_at"`
}

func (a *Feedback) FromPublicModel(feedbackPublic interface{}) {
	_ = encoding.TransformObject(feedbackPublic, a)
}

func (a *Feedback) ToPublicModel() *public.FeedbackResponse {
	feedbackPublic := &public.FeedbackResponse{}
	_ = encoding.TransformObject(a, feedbackPublic)
	return feedbackPublic
}

func (a *Feedback) FromRepositoryModel(feedbackRepo interface{}) {
	_ = encoding.TransformObject(feedbackRepo, a)
}

func (a *Feedback) ToRepositoryModel() *repository.Feedback {
	feedbackRepo := &repository.Feedback{}
	_ = encoding.TransformObject(a, feedbackRepo)
	return feedbackRepo
}
