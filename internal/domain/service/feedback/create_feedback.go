package feedback

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/domain"
	"github.com/medicplus-inc/medicplus-feedback/internal/global"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

// CreateFeedback creates a new feedback
func (s *FeedbackService) CreateFeedback(ctx context.Context, params *public.CreateFeedbackRequest) (*public.FeedbackResponse, error) {
	userLoggedIn, _ := global.GetClaimsFromContext(ctx)
	userLoggedInID := uuid.MustParse(userLoggedIn["uuid"].(string))

	feedbackDomain := &domain.Feedback{
		FeedbackType:   internal.ReceiverType(params.FeedbackType),
		FeedbackToID:   params.FeedbackToID,
		FeedbackFromID: userLoggedInID,
		FeedbackValue:  params.FeedbackValue,
		Notes:          params.Notes,
	}

	feedbackRepo := feedbackDomain.ToRepositoryModel()

	insertedFeedback, err := s.repository.InsertFeedback(ctx, feedbackRepo)
	if err != nil {
		return nil, err
	}
	if insertedFeedback == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	feedbackDomain.FromRepositoryModel(insertedFeedback)

	return feedbackDomain.ToPublicModel(), nil
}
