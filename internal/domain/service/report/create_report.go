package report

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/domain"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

// CreateReport creates a new report
func (s *ReportService) CreateReport(ctx context.Context, params *public.CreateReportRequest) (*public.ReportResponse, error) {
	report := &domain.Report{
		ReportType:   internal.ReceiverType(params.ReportType),
		ReportToID:   params.ReportToID,
		ReportFromID: params.ReportFromID,
		Context:      internal.ReportContext(params.Context),
		ContextID:    params.ContextID,
	}
	reportRepo := report.ToRepositoryModel()
	insertedRepo, err := s.repository.InsertReport(ctx, reportRepo)
	if err != nil {
		return nil, err
	}
	if insertedRepo == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	report.FromRepositoryModel(insertedRepo)

	return report.ToPublicModel(), nil
}
