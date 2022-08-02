package feedback

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

//DeleteFeedback deleting feedback
func (s *FeedbackService) DeleteFeedback(ctx context.Context, params *public.DeleteFeedbackRequest) error {

	feedback, err := s.repository.FindFeedbackByID(ctx, params.ID)
	if err != nil {
		return err
	}
	if feedback == nil {
		return libError.New(internal.ErrNotFound, http.StatusNotFound, internal.ErrNotFound.Error())
	}

	err = s.repository.DeleteFeedback(ctx, feedback)
	if err != nil {
		return err
	}

	return nil
}
