package feedback

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/domain"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

// GetFeedback get feedback by its id
func (s *FeedbackService) GetFeedback(ctx context.Context, feedbackID uuid.UUID) (*public.FeedbackResponse, error) {
	feedbackDomain := &domain.Feedback{}
	feedbackRepo, err := s.repository.FindFeedbackByID(ctx, feedbackID)
	if err != nil {
		return nil, err
	}
	if feedbackRepo == nil {
		return nil, libError.New(internal.ErrNotFound, http.StatusNotFound, internal.ErrNotFound.Error())
	}
	feedbackDomain.FromRepositoryModel(feedbackRepo)

	return feedbackDomain.ToPublicModel(), nil
}
