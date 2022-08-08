package feedback_parameter

import (
	"context"
	"fmt"
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
	userLoggedInID := uuid.MustParse(userLoggedIn["uuid"].(string))
	userLoggedInName := userLoggedIn["name"].(string)
	userLoggedInRole := userLoggedIn["role"].(string)

	fmt.Println(params.FeedbackType)
	fmt.Println(params.Name)
	fmt.Println(params.LanguageCode)
	fmt.Println(params.IsDefault)
	fmt.Println("before domain")

	fmt.Println(userLoggedInID)
	fmt.Println(userLoggedInName)
	fmt.Println(userLoggedInRole)

	feedbackParameter := &domain.FeedbackParameter{
		FeedbackType: internal.ParameterType(params.FeedbackType),
		Name:         params.Name,
		LanguageCode: internal.LanguageCode(params.LanguageCode),
		IsDefault:    params.IsDefault,
		CreatedBy:    userLoggedInID,
		UpdatedBy:    userLoggedInID,
	}

	fmt.Println("after domain")

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
