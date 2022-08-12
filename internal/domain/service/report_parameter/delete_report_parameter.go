package report_parameter

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/global"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

//DeleteReportParameter deleting report category
func (s *ReportParameterService) DeleteReportParameter(ctx context.Context, params *public.DeleteReportParameterRequest) error {
	userLoggedIn, _ := global.GetClaimsFromContext(ctx)
	userLoggedInID := uuid.MustParse(userLoggedIn["uuid"].(string))

	reportParameterRepo, err := s.repository.FindReportParameterByID(ctx, params.ID)
	if err != nil {
		return err
	}
	if reportParameterRepo == nil {
		return libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	reportParameterRepo.DeletedBy = &userLoggedInID

	err = s.repository.DeleteReportParameter(ctx, reportParameterRepo)
	if err != nil {
		return err
	}

	return nil
}
