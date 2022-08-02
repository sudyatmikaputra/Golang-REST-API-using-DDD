package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
)

type FeedbackParameter struct {
	ID            uuid.UUID              `json:"feedback_parameter" gorm:"primaryKey,not null"`
	ParameterType internal.ParameterType `json:"parameter_type" gorm:"not null"`
	Name          string                 `json:"name" gorm:"index:idx_feedback_name,not null"`
	LanguageCode  internal.LanguageCode  `json:"language_code" gorm:"index:idx_feedback_language_code,not null"`
	IsDefault     bool                   `json:"is_default" gorm:"not null"`
	CreatedBy     uuid.UUID              `json:"created_by" gorm:"not null"`
	CreatedAt     time.Time              `json:"created_at" gorm:"not null,autoCreateTime"`
	UpdatedAt     time.Time              `json:"updated_at" gorm:"not null,autoUpdateTime"`
	UpdatedBy     uuid.UUID              `json:"updated_by" gorm:"not null"`
	DeletedAt     *time.Time             `json:"deleted_at"`
	DeletedBy     *uuid.UUID             `json:"deleted_by"`
}

type FeedbackParameterRepository interface {
	FindAllFeedbackParameters(ctx context.Context, params *public.ListFeedbackParameterRequest) ([]FeedbackParameter, error)
	FindFeedbackParameterByParameterType(ctx context.Context, parameterType internal.ParameterType, languageCode string) (*FeedbackParameter, error)
	FindFeedbackParameterByID(ctx context.Context, feedbackParameterID uuid.UUID) (*FeedbackParameter, error)
	InsertFeedbackParameter(ctx context.Context, feedbackParameter *FeedbackParameter) (*FeedbackParameter, error)
	UpdateFeedbackParameter(ctx context.Context, feedbackParameter *FeedbackParameter) (*FeedbackParameter, error)
	DeleteFeedbackParameter(ctx context.Context, feedbackParameter *FeedbackParameter) error
}
