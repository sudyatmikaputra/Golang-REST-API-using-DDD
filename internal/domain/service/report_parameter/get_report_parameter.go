package report_parameter

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/domain"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

// GetReportParameter get report category by its id
func (s *ReportParameterService) GetReportParameter(ctx context.Context, reportParameterID uuid.UUID) (*public.ReportParameterResponse, error) {
	reportParameterDomain := &domain.ReportParameter{}
	reportParameterRepo, err := s.repository.FindReportParameterByID(ctx, reportParameterID)
	if err != nil {
		return nil, err
	}
	if reportParameterRepo == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}
	reportParameterDomain.FromRepositoryModel(reportParameterRepo)

	return reportParameterDomain.ToPublicModel(), nil
}
