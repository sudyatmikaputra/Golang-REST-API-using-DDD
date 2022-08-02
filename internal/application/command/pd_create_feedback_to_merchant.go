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

type CreateFeedbackForPatientDoctorToMerchantCommand struct {
	feedbackService          feedbackDomainService.FeedbackServiceInterface
	feedbackParameterService feedbackParameterDomainService.FeedbackParameterServiceInterface
}

func NewCreateFeedbackForPatientDoctorToMerchantCommand(
	feedbackService feedbackDomainService.FeedbackServiceInterface,
	feedbackParameterService feedbackParameterDomainService.FeedbackParameterServiceInterface,
) CreateFeedbackForPatientDoctorToMerchantCommand {
	return CreateFeedbackForPatientDoctorToMerchantCommand{
		feedbackService:          feedbackService,
		feedbackParameterService: feedbackParameterService,
	}
}

func (r CreateFeedbackForPatientDoctorToMerchantCommand) Execute(ctx context.Context, params public.CreateFeedbackRequest) (*public.FeedbackResponse, error) {
	feedback, err := r.feedbackService.CreateFeedback(ctx, &public.CreateFeedbackRequest{
		FeedbackType:   string(internal.ToMerchant),
		FeedbackToID:   params.FeedbackToID,
		FeedbackFromID: params.FeedbackFromID,
		FeedbackValue:  params.FeedbackValue,
		Notes:          params.Notes,
	})
	if err != nil {
		return nil, err
	}
	if feedback == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	feedbackParameter, err := r.feedbackParameterService.GetFeedbackParameterByParameterType(ctx, internal.ParameterType(params.FeedbackType), string(internal.BahasaIndonesia))
	if err != nil {
		return nil, err
	}

	feedback.FeedbackParameter = *feedbackParameter

	return feedback, nil
}
