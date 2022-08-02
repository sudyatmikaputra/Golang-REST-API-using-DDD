package command

import (
	"context"

	feedbackParameterDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/feedback_parameter"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
)

// DeleteFeedbackParameterCommand encapsulate process for deleting feedback parameter in Command
type DeleteFeedbackParameterCommand struct {
	parameterService feedbackParameterDomainService.FeedbackParameterServiceInterface
}

// NewDeleteFeedbackParameterCommand build an Command for deleting feedback parameter
func NewDeleteFeedbackParameterCommand(
	parameterService feedbackParameterDomainService.FeedbackParameterServiceInterface,
) DeleteFeedbackParameterCommand {
	return DeleteFeedbackParameterCommand{
		parameterService: parameterService,
	}
}

func (r DeleteFeedbackParameterCommand) Execute(ctx context.Context, params public.DeleteFeedbackParameterRequest) error {

	err := r.parameterService.DeleteFeedbackParameter(ctx, &params)
	if err != nil {
		return err
	}

	return nil
}
