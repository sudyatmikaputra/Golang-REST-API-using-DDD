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

func (s *FeedbackParameterService) UpdateFeedbackParameter(ctx context.Context, params *public.UpdateFeedbackParameterRequest) (*public.FeedbackParameterResponse, error) {
	userLoggedIn, _ := global.GetClaimsFromContext(ctx)
	userLoggedInID := uuid.MustParse(userLoggedIn["uuid"].(string))

	updatedFeedbackParameterRepo, err := s.repository.FindFeedbackParameterByID(ctx, params.ID)
	if err != nil {
		return nil, err
	}
	if updatedFeedbackParameterRepo == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	updatedFeedbackParameterDomain := &domain.FeedbackParameter{}
	updatedFeedbackParameterDomain.FromRepositoryModel(updatedFeedbackParameterRepo)
	if params.FeedbackType != "" {
		updatedFeedbackParameterDomain.FeedbackType = internal.ParameterType(params.FeedbackType)
	}
	if params.Name != "" {
		updatedFeedbackParameterDomain.Name = params.Name
	}
	if params.LanguageCode != "" {
		updatedFeedbackParameterDomain.LanguageCode = internal.LanguageCode(params.LanguageCode)
	}
	updatedFeedbackParameterDomain.IsDefault = params.IsDefault
	updatedFeedbackParameterRepo.UpdatedBy = userLoggedInID

	updatedFeedbackParameterRepo, err = s.repository.UpdateFeedbackParameter(ctx, updatedFeedbackParameterDomain.ToRepositoryModel())
	if err != nil {
		return nil, err
	}
	updatedFeedbackParameterDomain.FromRepositoryModel(updatedFeedbackParameterRepo)

	return updatedFeedbackParameterDomain.ToPublicModel(), nil
}
