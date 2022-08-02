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

// ListFeedbacksQuery encapsulate process for list feedbacks in query
type ListFeedbacksQuery struct {
	feedbackService  feedbackDomainService.FeedbackServiceInterface
	parameterService feedbackParameterDomainService.FeedbackParameterServiceInterface
}

// NewListFeedbacksQuery build an query for list feedbacks
func NewListFeedbacksQuery(
	feedbackService feedbackDomainService.FeedbackServiceInterface,
	parameterService feedbackParameterDomainService.FeedbackParameterServiceInterface,
) ListFeedbacksQuery {
	return ListFeedbacksQuery{
		feedbackService:  feedbackService,
		parameterService: parameterService,
	}
}

func (r ListFeedbacksQuery) ExecutePatient(ctx context.Context, params public.ListFeedbackRequest) ([]public.FeedbackResponse, error) {
	feedbacks, err := r.feedbackService.ListFeedbacks(ctx, &params)
	if err != nil {
		return nil, err
	}
	if feedbacks == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	// parameterIDs := []uuid.UUID{}
	// feedbackParameterMaps := map[uuid.UUID]*public.FeedbackResponse{}
	// for _, feedback := range feedbacks {
	// 	// feedbackParameterMaps[feedback.FeedbackParam.ID] = feedback
	// 	parameterIDs = append(parameterIDs, feedback.FeedbackParam.ID)
	// }

	// allParameters, err := r.parameterService.ListFeedbackParameters(ctx, &public.ListFeedbackParameterRequest{
	// 	IDs: parameterIDs,
	// })
	// if err != nil {
	// 	return nil, err
	// }

	// for _, parameter := range allParameters {
	// 	feedbackParameterMaps[parameter.ID].FeedbackParam = parameter
	// }

	// for _, feedback := range feedbacks {
	// 	feedback.FeedbackParam = feedbackParameterMaps[feedback.FeedbackParam.ID].FeedbackParam
	// }

	return feedbacks, nil
}
