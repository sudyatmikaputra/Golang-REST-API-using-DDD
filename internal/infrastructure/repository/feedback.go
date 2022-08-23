package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
)

type Feedback struct {
	ID                  uuid.UUID             `json:"id" gorm:"primaryKey,not null"`
	FeedbackType        internal.ReceiverType `json:"feedback_type" gorm:"index:idx_feedback_feedback_type, not null"`
	FeedbackToID        uuid.UUID             `json:"feedback_to_id" gorm:"index:idx_feedback_feedback_to_id, not null"`
	FeedbackFromID      uuid.UUID             `json:"feedback_from_id" gorm:"index:idx_feedback_feedback_from_id, not null"`
	FeedbackParameterID uuid.UUID             `json:"feedback_parameter_id" gorm:"index:idx_feedback_feedback_parameter_id, not null"`
	FeedbackValue       int                   `json:"feedback_value" gorm:"not null"`
	Notes               string                `json:"notes"`
	CreatedAt           time.Time             `json:"created_at" gorm:"not null,autoCreateTime"`
	UpdatedAt           time.Time             `json:"updated_at" gorm:"not null,autoUpdateTime"`
}

type FeedbackRepository interface {
	FindAllFeedbacks(ctx context.Context, params *public.ListFeedbackRequest) ([]Feedback, error)
	FindFeedbackByID(ctx context.Context, feedbackID uuid.UUID) (*Feedback, error)
	InsertFeedback(ctx context.Context, feedback *Feedback) (*Feedback, error)
	UpdateFeedback(ctx context.Context, feedback *Feedback) (*Feedback, error)
	DeleteFeedback(ctx context.Context, feedback *Feedback) error
}
