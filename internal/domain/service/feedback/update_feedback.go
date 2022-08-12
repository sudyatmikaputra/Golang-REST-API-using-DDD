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

// UpdateFeedback updates feedback data
func (s *FeedbackService) UpdateFeedback(ctx context.Context, params *public.UpdateFeedbackRequest) (*public.FeedbackResponse, error) {
	userLoggedIn, _ := global.GetClaimsFromContext(ctx)
	userLoggedInID := uuid.MustParse(userLoggedIn["uuid"].(string))

	updatedFeedbackRepo, err := s.repository.FindFeedbackByID(ctx, params.ID)
	if err != nil {
		return nil, err
	}
	if updatedFeedbackRepo == nil {
		return nil, libError.New(internal.ErrNotFound, http.StatusNotFound, internal.ErrNotFound.Error())
	}

	if userLoggedInID != updatedFeedbackRepo.FeedbackFromID {
		return nil, libError.New(internal.ErrNotAuthorized, http.StatusUnauthorized, internal.ErrNotAuthorized.Error())
	}
	updatedFeedbackDomain := &domain.Feedback{}
	updatedFeedbackDomain.FromRepositoryModel(updatedFeedbackRepo)

	if params.FeedbackValue != 0 {
		updatedFeedbackDomain.FeedbackValue = params.FeedbackValue
	}
	if params.Notes != "" {
		updatedFeedbackDomain.Notes = params.Notes
	}

	updatedFeedbackRepo, err = s.repository.UpdateFeedback(ctx, updatedFeedbackDomain.ToRepositoryModel())
	if err != nil {
		return nil, err
	}
	updatedFeedbackDomain.FromRepositoryModel(updatedFeedbackRepo)

	return updatedFeedbackDomain.ToPublicModel(), nil
}
