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
	feedbackService          feedbackDomainService.FeedbackServiceInterface
	feedbackParameterService feedbackParameterDomainService.FeedbackParameterServiceInterface
}

func NewListFeedbacksForAdminQuery(
	feedbackService feedbackDomainService.FeedbackServiceInterface,
	feedbackParameterService feedbackParameterDomainService.FeedbackParameterServiceInterface,
) ListFeedbacksForAdminQuery {
	return ListFeedbacksForAdminQuery{
		feedbackService:          feedbackService,
		feedbackParameterService: feedbackParameterService,
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

	feedbackParameter, _ := r.feedbackParameterService.GetFeedbackParameterByParameterType(ctx, internal.ParameterType(params.FeedbackType), params.LanguageCode)
	if feedbackParameter != nil {
		for i := range feedbacks {
			feedbacks[i].FeedbackParameter = public.FeedbackParameterResponse{
				ID:           feedbackParameter.ID,
				FeedbackType: feedbackParameter.FeedbackType,
				Name:         feedbackParameter.Name,
				LanguageCode: feedbackParameter.LanguageCode,
				IsDefault:    feedbackParameter.IsDefault,
			}
		}
	}

	return feedbacks, nil
}
