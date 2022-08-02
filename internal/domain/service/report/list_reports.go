package report

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/domain"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

// ListReports is listing all reports
func (s *ReportService) ListReports(ctx context.Context, params *public.ListReportRequest) ([]public.ReportResponse, error) {
	reportRepo, err := s.repository.FindAllReports(ctx, params)
	if err != nil {
		return nil, err
	}
	if reportRepo == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	result := []public.ReportResponse{}
	for _, _report := range reportRepo {
		report := &domain.Report{}
		report.FromRepositoryModel(_report)
		reportPublicModel := report.ToPublicModel()
		result = append(result, *reportPublicModel)
	}

	return result, nil
}
