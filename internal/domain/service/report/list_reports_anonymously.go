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
func (s *ReportService) ListReportsAnonymously(ctx context.Context, params *public.ListReportRequest) ([]public.AnonymousReportResponse, error) {
	reportRepo, err := s.repository.FindAllReports(ctx, params)
	if err != nil {
		return nil, err
	}
	if reportRepo == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	result := []public.AnonymousReportResponse{}
	for _, _report := range reportRepo {
		report := &domain.Report{}
		report.FromRepositoryModel(_report)
		anonymousReport := &public.AnonymousReportResponse{
			ID:         report.ID,
			ReportType: string(report.ReportType),
			ReportToID: report.ReportToID,
			Context:    string(report.Context),
			ContextID:  report.ContextID,
			Notes:      report.Notes,
		}
		result = append(result, *anonymousReport)
	}

	return result, nil
}
