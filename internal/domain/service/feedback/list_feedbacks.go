package feedback

import (
	"context"
	"net/http"

	libError "github.com/medicplus-inc/medicplus-kit/error"

	"github.com/medicplus-inc/medicplus-feedback/internal"

	"github.com/medicplus-inc/medicplus-feedback/internal/domain"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
)

// ListFeedbacks is listing all feedbacks
func (s *FeedbackService) ListFeedbacks(ctx context.Context, params *public.ListFeedbackRequest) ([]public.FeedbackResponse, error) {
	feedbackRepo, err := s.repository.FindAllFeedbacks(ctx, params)
	if err != nil {
		return nil, err
	}
	if feedbackRepo == nil {
		return nil, libError.New(internal.ErrNotFound, http.StatusNotFound, internal.ErrNotFound.Error())
	}

	result := []public.FeedbackResponse{}
	for _, _feedback := range feedbackRepo {
		feedback := &domain.Feedback{}
		feedback.FromRepositoryModel(_feedback)

		feedbackPublicMode := feedback.ToPublicModel()
		result = append(result, *feedbackPublicMode)
	}

	return result, nil
}
