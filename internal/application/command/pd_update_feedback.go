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

type UpdateFeedbackForPatientDoctorCommand struct {
	feedbackService          feedbackDomainService.FeedbackServiceInterface
	feedbackParameterService feedbackParameterDomainService.FeedbackParameterServiceInterface
}

func NewUpdateFeedbackForPatientDoctorCommand(
	feedbackService feedbackDomainService.FeedbackServiceInterface,
	feedbackParameterService feedbackParameterDomainService.FeedbackParameterServiceInterface,
) UpdateFeedbackForPatientDoctorCommand {
	return UpdateFeedbackForPatientDoctorCommand{
		feedbackService:          feedbackService,
		feedbackParameterService: feedbackParameterService,
	}
}

func (r UpdateFeedbackForPatientDoctorCommand) Execute(ctx context.Context, params public.UpdateFeedbackRequest) (*public.FeedbackResponse, error) {
	feedback, err := r.feedbackService.UpdateFeedback(ctx, &params)
	if err != nil {
		return nil, err
	}
	if feedback == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	feedbackParameter, _ := r.feedbackParameterService.GetFeedbackParameterByParameterType(ctx, internal.ParameterType(feedback.FeedbackType), string(internal.BahasaIndonesia))
	if feedbackParameter != nil {
		feedback.FeedbackParameter = public.FeedbackParameterResponse{
			ID:           feedbackParameter.ID,
			FeedbackType: feedbackParameter.FeedbackType,
			Name:         feedbackParameter.Name,
			LanguageCode: feedbackParameter.LanguageCode,
			IsDefault:    feedbackParameter.IsDefault,
		}
	}

	return feedback, nil
}
