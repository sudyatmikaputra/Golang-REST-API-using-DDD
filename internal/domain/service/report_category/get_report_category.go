package reportcategory

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/domain"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

// GetReportCategory get report category by its id
func (s *ReportCategoryService) GetReportCategory(ctx context.Context, categoryID uuid.UUID) (*public.ReportCategoryResponse, error) {
	report := &domain.ReportCategory{}
	reportRepo, err := s.repository.FindReportCategoryByID(ctx, categoryID)
	if err != nil {
		return nil, err
	}
	if reportRepo == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}
	report.FromRepositoryModel(reportRepo)

	return report.ToPublicModel(), nil
}
