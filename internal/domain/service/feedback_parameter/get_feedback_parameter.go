package feedback_parameter

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/domain"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

// GetFeedbackParameter get feedback parameter by its id
func (s *FeedbackParameterService) GetFeedbackParameter(ctx context.Context, parameterID uuid.UUID) (*public.FeedbackParameterResponse, error) {
	feedbackRepo, err := s.repository.FindFeedbackParameterByID(ctx, parameterID)
	if err != nil {
		return nil, err
	}
	if feedbackRepo == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}
	feedback := &domain.FeedbackParameter{}

	feedback.FromRepositoryModel(feedbackRepo)

	return feedback.ToPublicModel(), nil
}
