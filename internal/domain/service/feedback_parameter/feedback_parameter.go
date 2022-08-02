package feedback_parameter

import (
	"context"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/infrastructure/repository"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
)

// FeedbackParameterServiceInterface represents the feedback parameter service interface
type FeedbackParameterServiceInterface interface {
	ListFeedbackParameters(ctx context.Context, params *public.ListFeedbackParameterRequest) ([]public.FeedbackParameterResponse, error)
	GetFeedbackParameter(ctx context.Context, feedbackParameterID uuid.UUID) (*public.FeedbackParameterResponse, error)
	GetFeedbackParameterByParameterType(ctx context.Context, feedbackType internal.ParameterType, languageCode string) (*public.FeedbackParameterResponse, error)
	CreateFeedbackParameter(ctx context.Context, params *public.CreateFeedbackParameterRequest) (*public.FeedbackParameterResponse, error)
	UpdateFeedbackParameter(ctx context.Context, params *public.UpdateFeedbackParameterRequest) (*public.FeedbackParameterResponse, error)
	DeleteFeedbackParameter(ctx context.Context, params *public.DeleteFeedbackParameterRequest) error
}

// FeedbackParameterService is the domain logic implementation of feedback parameter service interface
type FeedbackParameterService struct {
	repository repository.FeedbackParameterRepository
}

// NewFeedbackParameterService creates a new feedback parameter domain service
func NewFeedbackParameterService(
	repository repository.FeedbackParameterRepository,
) FeedbackParameterServiceInterface {
	return &FeedbackParameterService{
		repository: repository,
	}
}
