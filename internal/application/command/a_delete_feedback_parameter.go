package command

import (
	"context"

	feedbackParameterDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/feedback_parameter"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
)

type DeleteFeedbackParameterForAdminCommand struct {
	parameterService feedbackParameterDomainService.FeedbackParameterServiceInterface
}

func NewDeleteFeedbackParameterForAdminCommand(
	parameterService feedbackParameterDomainService.FeedbackParameterServiceInterface,
) DeleteFeedbackParameterForAdminCommand {
	return DeleteFeedbackParameterForAdminCommand{
		parameterService: parameterService,
	}
}

func (r DeleteFeedbackParameterForAdminCommand) Execute(ctx context.Context, params public.DeleteFeedbackParameterRequest) error {

	err := r.parameterService.DeleteFeedbackParameter(ctx, &params)
	if err != nil {
		return err
	}

	return nil
}
