package feedback_parameter

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

// CreateFeedbackParameter creates a new feedback parameter
func (s *FeedbackParameterService) CreateFeedbackParameter(ctx context.Context, params *public.CreateFeedbackParameterRequest) (*public.FeedbackParameterResponse, error) {
	userLoggedIn, _ := global.GetClaimsFromContext(ctx)

	existingFeedbackParameter, err := s.repository.FindFeedbackParameterByParameterType(ctx, internal.ParameterType(params.ParameterType), params.LanguageCode)
	if err != nil {
		return nil, err
	}
	if existingFeedbackParameter != nil {
		return nil, libError.New(internal.ErrLanguageCodeAlreadyExists, http.StatusBadRequest, internal.ErrLanguageCodeAlreadyExists.Error())
	}

	feedbackParameter := &domain.FeedbackParameter{
		ParameterType: internal.ParameterType(params.ParameterType),
		Name:          params.Name,
		LanguageCode:  internal.LanguageCode(params.LanguageCode),
		IsDefault:     params.IsDefault,
	}

	feedbackParameterRepo := feedbackParameter.ToRepositoryModel()
	feedbackParameterRepo.CreatedBy = userLoggedIn["uuid"].(uuid.UUID)
	feedbackParameterRepo.UpdatedBy = userLoggedIn["uuid"].(uuid.UUID)

	insertedFeedback, err := s.repository.InsertFeedbackParameter(ctx, feedbackParameterRepo)
	if err != nil {
		return nil, err
	}
	if insertedFeedback == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	feedbackParameter.FromRepositoryModel(insertedFeedback)

	return feedbackParameter.ToPublicModel(), nil
}
