package feedback_parameter

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/global"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

func (s *FeedbackParameterService) DeleteFeedbackParameter(ctx context.Context, params *public.DeleteFeedbackParameterRequest) error {
	userLoggedIn, _ := global.GetClaimsFromContext(ctx)
	userLoggedInID := uuid.MustParse(userLoggedIn["uuid"].(string))

	feedbackParameter, err := s.repository.FindFeedbackParameterByID(ctx, params.ID)
	if err != nil {
		return err
	}
	if feedbackParameter == nil {
		return libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	feedbackParameter.DeletedBy = &userLoggedInID
	updatedFeedbackParameter, err := s.repository.UpdateFeedbackParameter(ctx, feedbackParameter)
	if err != nil {
		return err
	}
	if updatedFeedbackParameter == nil {
		return libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	err = s.repository.DeleteFeedbackParameter(ctx, updatedFeedbackParameter)
	if err != nil {
		return err
	}

	return nil
}
