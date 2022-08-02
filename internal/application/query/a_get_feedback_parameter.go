package query

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	feedbackParameterDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/feedback_parameter"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

type GetFeedbackParameterForAdminQuery struct {
	feedbackParameterService feedbackParameterDomainService.FeedbackParameterServiceInterface
}

func NewGetFeedbackParameterForAdminQuery(
	feedbackParameterService feedbackParameterDomainService.FeedbackParameterServiceInterface,
) GetFeedbackParameterForAdminQuery {
	return GetFeedbackParameterForAdminQuery{
		feedbackParameterService: feedbackParameterService,
	}
}

func (r GetFeedbackParameterForAdminQuery) Execute(ctx context.Context, params public.GetFeedbackParameterRequest) (*public.FeedbackParameterResponse, error) {
	feedbackParameter, err := r.feedbackParameterService.GetFeedbackParameter(ctx, params.ID)
	if err != nil {
		return nil, err
	}
	if feedbackParameter == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	return feedbackParameter, nil
}
