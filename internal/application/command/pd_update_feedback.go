package command

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	feedbackDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/feedback"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

type UpdateFeedbackForPatientDoctorCommand struct {
	feedbackService feedbackDomainService.FeedbackServiceInterface
}

func NewUpdateFeedbackForPatientDoctorCommand(
	feedbackService feedbackDomainService.FeedbackServiceInterface,
) UpdateFeedbackForPatientDoctorCommand {
	return UpdateFeedbackForPatientDoctorCommand{
		feedbackService: feedbackService,
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

	return feedback, nil
}
