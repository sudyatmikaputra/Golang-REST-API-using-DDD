package report

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

//DeleteReport deleting report
func (s *ReportService) DeleteReport(ctx context.Context, params *public.DeleteReportRequest) error {
	reportRepo, err := s.repository.FindReportByID(ctx, params.ReportID)
	if err != nil {
		return err
	}
	if reportRepo == nil {
		return libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	err = s.repository.DeleteReport(ctx, reportRepo)
	if err != nil {
		return err
	}

	return nil
}
