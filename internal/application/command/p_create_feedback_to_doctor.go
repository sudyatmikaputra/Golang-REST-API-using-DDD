package command

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

func (r CreateFeedbackCommand) ExecuteToDoctor(ctx context.Context, params public.CreateFeedbackRequest) (*public.FeedbackResponse, error) {

	feedback, err := r.feedbackService.CreateFeedback(ctx, &public.CreateFeedbackRequest{
		FeedbackTo:      string(internal.ToDoctor),
		FeedbackParamID: params.FeedbackParamID,
		FeedbackToID:    params.FeedbackToID,
		FeedbackFromID:  params.FeedbackFromID,
		FeedbackValue:   params.FeedbackValue,
		Notes:           params.Notes,
	})
	if err != nil {
		return nil, err
	}
	if feedback == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	return feedback, nil
}
