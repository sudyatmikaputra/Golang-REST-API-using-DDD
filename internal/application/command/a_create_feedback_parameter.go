package command

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	feedbackParameterDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/feedback_parameter"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

type CreateFeedbackParameterForAdminCommand struct {
	parameterService feedbackParameterDomainService.FeedbackParameterServiceInterface
}

func NewCreateFeedbackParameterForAdminCommand(
	parameterService feedbackParameterDomainService.FeedbackParameterServiceInterface,
) CreateFeedbackParameterForAdminCommand {
	return CreateFeedbackParameterForAdminCommand{
		parameterService: parameterService,
	}
}

func (r CreateFeedbackParameterForAdminCommand) Execute(ctx context.Context, params public.CreateFeedbackParameterRequest) (*public.FeedbackParameterResponse, error) {
	parameter, err := r.parameterService.CreateFeedbackParameter(ctx, &params)
	if err != nil {
		return nil, err
	}
	if parameter == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	return parameter, nil
}
