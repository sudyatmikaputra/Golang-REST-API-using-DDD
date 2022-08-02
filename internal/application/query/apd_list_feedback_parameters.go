package query

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	feedbackParameterDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/feedback_parameter"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

type ListFeedbackParametersQuery struct {
	parameterService feedbackParameterDomainService.FeedbackParameterServiceInterface
}

func NewListFeedbackParametersQuery(
	parameterService feedbackParameterDomainService.FeedbackParameterServiceInterface,
) ListFeedbackParametersQuery {
	return ListFeedbackParametersQuery{
		parameterService: parameterService,
	}
}

func (r ListFeedbackParametersQuery) Execute(ctx context.Context, params public.ListFeedbackParameterRequest) ([]public.FeedbackParameterResponse, error) {
	parameters, err := r.parameterService.ListFeedbackParameters(ctx, &params)
	if err != nil {
		return nil, err
	}
	if parameters == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	return parameters, nil
}
