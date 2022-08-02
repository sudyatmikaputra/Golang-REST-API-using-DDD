package report_parameter

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/domain"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

func (s *ReportParameterService) GetReportParameterByReportType(ctx context.Context, reportType internal.ParameterType, languageCode string) (*public.ReportParameterResponse, error) {
	reportParameterRepo, err := s.repository.FindReportParameterByReportType(ctx, reportType, languageCode)
	if err != nil {
		return nil, err
	}
	if reportParameterRepo == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}
	reportParameterDomain := &domain.ReportParameter{}

	reportParameterDomain.FromRepositoryModel(reportParameterRepo)

	return reportParameterDomain.ToPublicModel(), nil
}
