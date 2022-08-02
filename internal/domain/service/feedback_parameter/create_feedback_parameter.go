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

	existingFeedbackParameter, err := s.repository.FindFeedbackParameterByParameterType(ctx, internal.ParameterType(params.FeedbackType), params.LanguageCode)
	if err != nil {
		return nil, err
	}
	if existingFeedbackParameter != nil {
		return nil, libError.New(internal.ErrLanguageCodeAlreadyExists, http.StatusBadRequest, internal.ErrLanguageCodeAlreadyExists.Error())
	}

	feedbackParameter := &domain.FeedbackParameter{
		FeedbackType: internal.ParameterType(params.FeedbackType),
		Name:         params.Name,
		LanguageCode: internal.LanguageCode(params.LanguageCode),
		IsDefault:    params.IsDefault,
		CreatedBy:    userLoggedIn["uuid"].(uuid.UUID),
		UpdatedBy:    userLoggedIn["uuid"].(uuid.UUID),
	}

	feedbackParameterRepo := feedbackParameter.ToRepositoryModel()

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
