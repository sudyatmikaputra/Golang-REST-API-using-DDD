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
	ID            uuid.UUID              `json:"feedback_parameter"`
	ParameterType internal.ParameterType `json:"parameter_type"`
	Name          string                 `json:"name"`
	LanguageCode  internal.LanguageCode  `json:"language_code"`
	IsDefault     bool                   `json:"is_default"`
	CreatedBy     uuid.UUID              `json:"created_by"`
	CreatedAt     time.Time              `json:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at"`
	UpdatedBy     uuid.UUID              `json:"updated_by"`
	DeletedAt     *time.Time             `json:"deleted_at"`
	DeletedBy     *uuid.UUID             `json:"deleted_by"`
}

func (a *FeedbackParameter) FromPublicModel(parameterPublic interface{}) {
	_ = encoding.TransformObject(parameterPublic, a)
}

func (a *FeedbackParameter) ToPublicModel() *public.FeedbackParameterResponse {
	parameterPublic := &public.FeedbackParameterResponse{}
	_ = encoding.TransformObject(a, parameterPublic)
	return parameterPublic
}

func (a *FeedbackParameter) FromRepositoryModel(feedbackRepo interface{}) {
	_ = encoding.TransformObject(feedbackRepo, a)
}

func (a *FeedbackParameter) ToRepositoryModel() *repository.FeedbackParameter {
	feedbackRepo := &repository.FeedbackParameter{}
	_ = encoding.TransformObject(a, feedbackRepo)
	return feedbackRepo
}
