package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/infrastructure/repository"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	"github.com/medicplus-inc/medicplus-kit/encoding"
)

type FeedbackParameter struct {
	ID           uuid.UUID              `json:"id"`
	FeedbackType internal.ParameterType `json:"feedback_type"`
	Name         string                 `json:"name"`
	LanguageCode internal.LanguageCode  `json:"language_code"`
	IsDefault    bool                   `json:"is_default"`
	CreatedBy    uuid.UUID              `json:"created_by"`
	CreatedAt    time.Time              `json:"created_at"`
	UpdatedAt    time.Time              `json:"updated_at"`
	UpdatedBy    uuid.UUID              `json:"updated_by"`
	DeletedAt    *time.Time             `json:"deleted_at"`
	DeletedBy    *uuid.UUID             `json:"deleted_by"`
}

func (a *FeedbackParameter) FromPublicModel(feedbackParameterPublic interface{}) {
	_ = encoding.TransformObject(feedbackParameterPublic, a)
}

func (a *FeedbackParameter) ToPublicModel() *public.FeedbackParameterResponse {
	feedbackParameterPublic := &public.FeedbackParameterResponse{}
	_ = encoding.TransformObject(a, feedbackParameterPublic)
	return feedbackParameterPublic
}

func (a *FeedbackParameter) FromRepositoryModel(feedbackParameterRepo interface{}) {
	_ = encoding.TransformObject(feedbackParameterRepo, a)
}

func (a *FeedbackParameter) ToRepositoryModel() *repository.FeedbackParameter {
	feedbackParameterRepo := &repository.FeedbackParameter{}
	_ = encoding.TransformObject(a, feedbackParameterRepo)
	return feedbackParameterRepo
}
