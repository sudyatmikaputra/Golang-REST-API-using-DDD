package feedback_parameter

import (
	"context"
	"net/http"

	libError "github.com/medicplus-inc/medicplus-kit/error"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/domain"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
)

// ListFeedbacksParameters is listing all feedback parameters
func (s *FeedbackParameterService) ListFeedbackParameters(ctx context.Context, params *public.ListFeedbackParameterRequest) ([]public.FeedbackParameterResponse, error) {
	feedbackRepo, err := s.repository.FindAllFeedbackParameters(ctx, params)
	if err != nil {
		return nil, err
	}
	if feedbackRepo == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	result := []public.FeedbackParameterResponse{}
	for _, _feedback := range feedbackRepo {
		feedback := &domain.FeedbackParameter{}
		feedback.FromRepositoryModel(_feedback)

		feedbackPublicModel := feedback.ToPublicModel()
		result = append(result, *feedbackPublicModel)
	}

	return result, nil
}
