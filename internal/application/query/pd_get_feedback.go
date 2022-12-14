package query

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	feedbackDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/feedback"
	feedbackParameterDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/feedback_parameter"
	"github.com/medicplus-inc/medicplus-feedback/internal/global"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

type GetFeedbackForPatientDoctorQuery struct {
	feedbackService          feedbackDomainService.FeedbackServiceInterface
	feedbackParameterService feedbackParameterDomainService.FeedbackParameterServiceInterface
}

func NewGetFeedbackForPatientDoctorQuery(
	feedbackService feedbackDomainService.FeedbackServiceInterface,
	feedbackParameterService feedbackParameterDomainService.FeedbackParameterServiceInterface,
) GetFeedbackForPatientDoctorQuery {
	return GetFeedbackForPatientDoctorQuery{
		feedbackService:          feedbackService,
		feedbackParameterService: feedbackParameterService,
	}
}

func (r GetFeedbackForPatientDoctorQuery) Execute(ctx context.Context, params public.GetFeedbackRequest) (*public.FeedbackResponse, error) {
	userLoggedIn, _ := global.GetClaimsFromContext(ctx)
	userLoggedInID := uuid.MustParse(userLoggedIn["uuid"].(string))

	feedback, err := r.feedbackService.GetFeedback(ctx, params.ID)
	if err != nil {
		return nil, err
	}
	if feedback == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}
	if userLoggedInID != feedback.FeedbackFromID && userLoggedInID != feedback.FeedbackToID {
		return nil, libError.New(internal.ErrNotAuthorized, http.StatusUnauthorized, internal.ErrNotAuthorized.Error())
	}

	feedbackParameter, _ := r.feedbackParameterService.GetFeedbackParameterByParameterType(ctx, internal.ParameterType(feedback.FeedbackType), string(internal.BahasaIndonesia))
	if feedbackParameter != nil {
		feedback.FeedbackParameter = public.FeedbackParameterResponse{
			ID:           feedbackParameter.ID,
			FeedbackType: feedbackParameter.FeedbackType,
			Name:         feedbackParameter.Name,
			LanguageCode: feedbackParameter.Name,
			IsDefault:    feedbackParameter.IsDefault,
		}
	}

	return feedback, nil
}
