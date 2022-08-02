package query

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

func (r ListFeedbacksQuery) ExecuteDoctor(ctx context.Context, params public.ListFeedbackRequest) ([]public.AnonymousFeedbackResponse, error) {
	feedbacks, err := r.feedbackService.ListFeedbacksAnonymously(ctx, &params)
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
