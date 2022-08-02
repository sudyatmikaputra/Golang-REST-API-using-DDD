package command

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	feedbackParameterDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/feedback_parameter"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

// UpdateFeedbackParameterCommand encapsulate process for updating feedback parameter in Command
type UpdateFeedbackParameterCommand struct {
	parameterService feedbackParameterDomainService.FeedbackParameterServiceInterface
}

// NewUpdateFeedbackParameterCommand build an Command for updating feedback parameter
func NewUpdateFeedbackParameterCommand(
	parameterService feedbackParameterDomainService.FeedbackParameterServiceInterface,
) UpdateFeedbackParameterCommand {
	return UpdateFeedbackParameterCommand{
		parameterService: parameterService,
	}
}

func (r UpdateFeedbackParameterCommand) Execute(ctx context.Context, params public.UpdateFeedbackParameterRequest) (*public.FeedbackParameterResponse, error) {
	parameter, err := r.parameterService.UpdateFeedbackParameter(ctx, &params)
	if err != nil {
		return nil, err
	}
	if parameter == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	return parameter, nil
}
