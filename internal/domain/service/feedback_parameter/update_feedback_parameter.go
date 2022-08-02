package feedback_parameter

import (
	"context"
	"net/http"

	libError "github.com/medicplus-inc/medicplus-kit/error"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/domain"
	"github.com/medicplus-inc/medicplus-feedback/internal/global"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
)

// UpdateFeedbackParameter updates feedback parameter data
func (s *FeedbackParameterService) UpdateFeedbackParameter(ctx context.Context, params *public.UpdateFeedbackParameterRequest) (*public.FeedbackParameterResponse, error) {
	userLoggedIn, _ := global.GetClaimsFromContext(ctx)
	updatedFeedbackDomain := &domain.FeedbackParameter{}
	updatedFeedbackRepo, err := s.repository.FindFeedbackParameterByID(ctx, params.ID)
	if err != nil {
		return nil, err
	}
	if updatedFeedbackRepo == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	updatedFeedbackDomain.FromRepositoryModel(updatedFeedbackRepo)
	if params.FeedbackType != "" {
		updatedFeedbackDomain.FeedbackType = internal.ParameterType(params.FeedbackType)
	}
	if params.Name != "" {
		updatedFeedbackDomain.Name = params.Name
	}
	if params.LanguageCode != "" {
		updatedFeedbackDomain.LanguageCode = internal.LanguageCode(params.LanguageCode)
	}
	updatedFeedbackDomain.IsDefault = params.IsDefault

	userLoggedInID := userLoggedIn["uuid"].(uuid.UUID)
	updatedFeedbackRepo.UpdatedBy = userLoggedInID

	updatedFeedbackRepo, err = s.repository.UpdateFeedbackParameter(ctx, updatedFeedbackRepo)
	if err != nil {
		return nil, err
	}
	updatedFeedbackDomain.FromRepositoryModel(updatedFeedbackRepo)

	return updatedFeedbackDomain.ToPublicModel(), nil
}
