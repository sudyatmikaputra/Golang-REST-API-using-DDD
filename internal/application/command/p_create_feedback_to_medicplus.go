package command

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	feedbackDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/feedback"
	feedbackParameterDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/feedback_parameter"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

// CreateFeedbackCommand encapsulate process for creating feedback in Command
type CreateFeedbackCommand struct {
	feedbackService  feedbackDomainService.FeedbackServiceInterface
	parameterService feedbackParameterDomainService.FeedbackParameterServiceInterface
}

// NewCreateFeedbackCommand build an Command for creating feedback
func NewCreateFeedbackCommand(
	feedbackService feedbackDomainService.FeedbackServiceInterface,
	parameterService feedbackParameterDomainService.FeedbackParameterServiceInterface,
) CreateFeedbackCommand {
	return CreateFeedbackCommand{
		feedbackService:  feedbackService,
		parameterService: parameterService,
	}
}

func (r CreateFeedbackCommand) ExecuteToMedicplus(ctx context.Context, params public.CreateFeedbackRequest) (*public.FeedbackResponse, error) {

	feedback, err := r.feedbackService.CreateFeedback(ctx, &public.CreateFeedbackRequest{
		FeedbackTo:      string(internal.ToMedicplus),
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
