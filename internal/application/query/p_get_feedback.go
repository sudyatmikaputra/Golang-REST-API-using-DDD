package query

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	feedbackDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/feedback"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

// GetFeedbackQuery encapsulate process for getting the feedback in query
type GetFeedbackQuery struct {
	feedbackService feedbackDomainService.FeedbackServiceInterface
}

// NewGetFeedbackQuery build an query for getting the feedback
func NewGetFeedbackQuery(
	feedbackService feedbackDomainService.FeedbackServiceInterface,
) GetFeedbackQuery {
	return GetFeedbackQuery{
		feedbackService: feedbackService,
	}
}

func (r GetFeedbackQuery) Execute(ctx context.Context, params public.GetFeedbackRequest) (*public.FeedbackResponse, error) {
	feedback, err := r.feedbackService.GetFeedback(ctx, params.FeedbackID)
	if err != nil {
		return nil, err
	}
	if feedback == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	return feedback, nil
}
