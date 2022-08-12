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
	userLoggedInID := uuid.MustParse(userLoggedIn["uuid"].(string))

	if params.FeedbackToID != uuid.Nil {
		if params.FeedbackType == string(internal.DoctorParameterType) && userLoggedInID != params.FeedbackToID {
			return nil, libError.New(internal.ErrNotAuthorized, http.StatusUnauthorized, internal.ErrNotAuthorized.Error())
		}
	}
	if params.FeedbackFromID != uuid.Nil {
		if params.FeedbackType != string(internal.DoctorParameterType) && userLoggedInID != params.FeedbackFromID {
			return nil, libError.New(internal.ErrNotAuthorized, http.StatusUnauthorized, internal.ErrNotAuthorized.Error())
		}
	}

	feedbacks, err := r.feedbackService.ListFeedbacksAnonymously(ctx, &params)
	if err != nil {
		return nil, err
	}
	if feedbacks == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	feedbackParameter, _ := r.feedbackParameterService.GetFeedbackParameterByParameterType(ctx, internal.ParameterType(params.FeedbackType), params.LanguageCode)
	if feedbackParameter != nil {
		for i := range feedbacks {
			feedbacks[i].FeedbackParameter = public.FeedbackParameterResponse{
				ID:           feedbackParameter.ID,
				FeedbackType: feedbackParameter.FeedbackType,
				Name:         feedbackParameter.Name,
				LanguageCode: feedbackParameter.LanguageCode,
				IsDefault:    feedbackParameter.IsDefault,
			}
		}
	}

	return feedbacks, nil
}
