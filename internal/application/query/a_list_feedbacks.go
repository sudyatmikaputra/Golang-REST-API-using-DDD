package query

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	feedbackDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/feedback"
	feedbackParameterDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/feedback_parameter"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

type ListFeedbacksForAdminQuery struct {
	feedbackService  feedbackDomainService.FeedbackServiceInterface
	parameterService feedbackParameterDomainService.FeedbackParameterServiceInterface
}

func NewListFeedbacksForAdminQuery(
	feedbackService feedbackDomainService.FeedbackServiceInterface,
	parameterService feedbackParameterDomainService.FeedbackParameterServiceInterface,
) ListFeedbacksForAdminQuery {
	return ListFeedbacksForAdminQuery{
		feedbackService:  feedbackService,
		parameterService: parameterService,
	}
}

func (r ListFeedbacksForAdminQuery) Execute(ctx context.Context, params public.ListFeedbackRequest) ([]public.FeedbackResponse, error) {
	feedbacks, err := r.feedbackService.ListFeedbacks(ctx, &params)
	if err != nil {
		return nil, err
	}
	if feedbacks == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	feedbackParameter, err := r.parameterService.GetFeedbackParameterByParameterType(ctx, internal.ParameterType(params.FeedbackType), params.LanguageCode)
	if err != nil {
		return nil, err
	}
	if feedbackParameter != nil {
		for _, feedback := range feedbacks {
			feedback.FeedbackParameter = *feedbackParameter
		}
	}

	return feedbacks, nil
}
