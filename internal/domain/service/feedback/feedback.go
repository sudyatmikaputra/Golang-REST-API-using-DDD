package feedback

import (
	"context"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal/infrastructure/repository"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
)

// FeedbackServiceInterface represents the feedback service interface
type FeedbackServiceInterface interface {
	ListFeedbacks(ctx context.Context, params *public.ListFeedbackRequest) ([]public.FeedbackResponse, error)
	ListFeedbacksAnonymously(ctx context.Context, params *public.ListFeedbackRequest) ([]public.AnonymousFeedbackResponse, error)
	GetFeedback(ctx context.Context, feedbackID uuid.UUID) (*public.FeedbackResponse, error)
	CreateFeedback(ctx context.Context, params *public.CreateFeedbackRequest) (*public.FeedbackResponse, error)
	UpdateFeedback(ctx context.Context, params *public.UpdateFeedbackRequest) (*public.FeedbackResponse, error)
	DeleteFeedback(ctx context.Context, params *public.DeleteFeedbackRequest) error
}

// FeedbackService is the domain logic implementation of feedback service interface
type FeedbackService struct {
	repository repository.FeedbackRepository
}

// NewFeedbackService creates a new feedback domain service
func NewFeedbackService(
	repository repository.FeedbackRepository,
) FeedbackServiceInterface {
	return &FeedbackService{
		repository: repository,
	}
}
