package feedback

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/domain"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

// CreateFeedback creates a new feedback
func (s *FeedbackService) CreateFeedback(ctx context.Context, params *public.CreateFeedbackRequest) (*public.FeedbackResponse, error) {

	feedback := &domain.Feedback{
		FeedbackTo:      internal.ReceiverType(params.FeedbackTo),
		FeedbackParamID: params.FeedbackParamID,
		FeedbackToID:    params.FeedbackToID,
		FeedbackFromID:  params.FeedbackFromID,
		FeedbackValue:   params.FeedbackValue,
		Notes:           params.Notes,
	}

	feedbackRepo := feedback.ToRepositoryModel()

	insertedFeedback, err := s.repository.InsertFeedback(ctx, feedbackRepo)
	if err != nil {
		return nil, err
	}
	if insertedFeedback == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	feedback.FromRepositoryModel(insertedFeedback)

	return feedback.ToPublicModel(), nil
}
