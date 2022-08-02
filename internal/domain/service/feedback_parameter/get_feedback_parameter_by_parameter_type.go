package feedback_parameter

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/domain"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"

	libError "github.com/medicplus-inc/medicplus-kit/error"
)

func (s *FeedbackParameterService) GetFeedbackParameterByParameterType(ctx context.Context, feedbackType internal.ParameterType, languageCode string) (*public.FeedbackParameterResponse, error) {
	feedbackRepo, err := s.repository.FindFeedbackParameterByParameterType(ctx, feedbackType, languageCode)
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
