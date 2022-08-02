package report

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/domain"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

// GetReport get report by its id
func (s *ReportService) GetReport(ctx context.Context, reportID uuid.UUID) (*public.ReportResponse, error) {
	report := &domain.Report{}
	reportRepo, err := s.repository.FindReportByID(ctx, reportID)
	if err != nil {
		return nil, err
	}
	if reportRepo == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}
	report.FromRepositoryModel(reportRepo)

	return report.ToPublicModel(), nil
}
