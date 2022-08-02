package query

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	feedbackDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/feedback"
	feedbackParameterDomainSerivce "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/feedback_parameter"
	"github.com/medicplus-inc/medicplus-feedback/internal/global"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

type ListFeedbacksForDoctorAnonymouslyQuery struct {
	feedbackService          feedbackDomainService.FeedbackServiceInterface
	feedbackParameterService feedbackParameterDomainSerivce.FeedbackParameterServiceInterface
}

func NewListFeedbacksForDoctorAnonymouslyQuery(
	feedbackService feedbackDomainService.FeedbackServiceInterface,
	feedbackParameterService feedbackParameterDomainSerivce.FeedbackParameterServiceInterface,
) ListFeedbacksForDoctorAnonymouslyQuery {
	return ListFeedbacksForDoctorAnonymouslyQuery{
		feedbackService:          feedbackService,
		feedbackParameterService: feedbackParameterService,
	}
}

func (r ListFeedbacksForDoctorAnonymouslyQuery) Execute(ctx context.Context, params public.ListFeedbackRequest) ([]public.AnonymousFeedbackResponse, error) {
	userLoggedIn, _ := global.GetClaimsFromContext(ctx)
	if params.FeedbackType == string(internal.DoctorParameterType) && userLoggedIn["uuid"].(uuid.UUID) != params.FeedbackToID {
		return nil, libError.New(internal.ErrNotAuthorized, http.StatusUnauthorized, internal.ErrNotAuthorized.Error())
	}
	if params.FeedbackType != string(internal.DoctorParameterType) && userLoggedIn["uuid"].(uuid.UUID) != params.FeedbackFromID {
		return nil, libError.New(internal.ErrNotAuthorized, http.StatusUnauthorized, internal.ErrNotAuthorized.Error())
	}

	feedbacks, err := r.feedbackService.ListFeedbacksAnonymously(ctx, &params)
	if err != nil {
		return nil, err
	}
	if feedbacks == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	feedbackParameter, err := r.feedbackParameterService.GetFeedbackParameterByParameterType(ctx, internal.ParameterType(params.FeedbackType), params.LanguageCode)
	if err != nil {
		return nil, err
	}
	if feedbackParameter != nil {
		for _, feedback := range feedbacks {
			feedback.FeedbackParameter = *feedbackParameter
		}
	}

	return feedbacks, nil
}
