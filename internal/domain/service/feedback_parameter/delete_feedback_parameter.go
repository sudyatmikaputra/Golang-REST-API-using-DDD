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

//DeleteFeedbackParameter deleting feedback parameter
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
	err = s.repository.DeleteFeedbackParameter(ctx, feedbackParameter)
	if err != nil {
		return err
	}

	return nil
}
