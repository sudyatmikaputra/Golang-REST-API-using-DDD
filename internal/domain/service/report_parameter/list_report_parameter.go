package report_parameter

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/domain"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

// ListReportParameters is listing all report categories
func (s *ReportParameterService) ListReportParameters(ctx context.Context, params *public.ListReportParameterRequest) ([]public.ReportParameterResponse, error) {
	reportParameterRepo, err := s.repository.FindAllReportParameters(ctx, params)
	if err != nil {
		return nil, err
	}
	if reportParameterRepo == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	result := []public.ReportParameterResponse{}
	for _, _report := range reportParameterRepo {
		report := &domain.ReportParameter{}
		report.FromRepositoryModel(_report)

		reportPublicModel := report.ToPublicModel()
		result = append(result, *reportPublicModel)
	}

	return result, nil
}
