package reportcategory

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/domain"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

// ListReportCategories is listing all report categories
func (s *ReportCategoryService) ListReportCategories(ctx context.Context, params *public.ListReportCategoryRequest) ([]public.ReportCategoryResponse, error) {
	reportRepo, err := s.repository.FindAllReportCategories(ctx, params)
	if err != nil {
		return nil, err
	}
	if reportRepo == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	result := []public.ReportCategoryResponse{}
	for _, _report := range reportRepo {
		report := &domain.ReportCategory{}
		report.FromRepositoryModel(_report)

		reportPublicModel := report.ToPublicModel()
		result = append(result, *reportPublicModel)
	}

	return result, nil
}
