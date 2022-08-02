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
	updatedFeedback := &domain.FeedbackParameter{}
	updatedFeedbackRepo, err := s.repository.FindFeedbackParameterByID(ctx, params.ID)
	if err != nil {
		return nil, err
	}
	if updatedFeedbackRepo == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	updatedFeedback.FromRepositoryModel(updatedFeedbackRepo)
	if params.ParameterType != "" {
		updatedFeedback.ParameterType = internal.ParameterType(params.ParameterType)
	}
	if params.Name != "" {
		updatedFeedback.Name = params.Name
	}
	if params.LanguageCode != "" {
		updatedFeedback.LanguageCode = internal.LanguageCode(params.LanguageCode)
	}
	updatedFeedback.IsDefault = params.IsDefault

	userLoggedInID := userLoggedIn["uuid"].(uuid.UUID)
	updatedFeedbackRepo.UpdatedBy = userLoggedInID

	updatedFeedbackRepo, err = s.repository.UpdateFeedbackParameter(ctx, updatedFeedbackRepo)
	if err != nil {
		return nil, err
	}
	updatedFeedback.FromRepositoryModel(updatedFeedbackRepo)

	return updatedFeedback.ToPublicModel(), nil
}
