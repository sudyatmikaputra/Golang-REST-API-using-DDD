package reportcategory

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/domain"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

func (s *ReportCategoryService) GetReportCategoryByReportType(ctx context.Context, reportType internal.ParameterType, languageCode string) (*public.ReportCategoryResponse, error) {
	reportRepo, err := s.repository.FindReportCategoryByReportType(ctx, reportType, languageCode)
	if err != nil {
		return nil, err
	}
	if reportRepo == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}
	report := &domain.ReportCategory{}

	report.FromRepositoryModel(reportRepo)

	return report.ToPublicModel(), nil
}
