package reportcategory

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/domain"
	"github.com/medicplus-inc/medicplus-feedback/internal/global"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

// UpdateReportCategory updates report category data
func (s *ReportCategoryService) UpdateReportCategory(ctx context.Context, params *public.UpdateReportCategoryRequest) (*public.ReportCategoryResponse, error) {
	userLoggedIn, _ := global.GetClaimsFromContext(ctx)
	if userLoggedIn["uuid"].(uuid.UUID) == uuid.Nil && userLoggedIn["role"].(string) != string(internal.Admin) {
		return nil, libError.New(internal.ErrNotAuthorized, http.StatusUnauthorized, internal.ErrNotAuthorized.Error())
	}
	updatedReport := &domain.ReportCategory{}
	updatedReportRepo, err := s.repository.FindReportCategoryByID(ctx, params.ID)
	if err != nil {
		return nil, err
	}
	if updatedReportRepo == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	updatedReport.FromRepositoryModel(updatedReportRepo)
	if params.ReportType != "" {
		updatedReport.ReportType = internal.ParameterType(params.ReportType)
	}
	if params.Name != "" {
		updatedReport.Name = params.Name
	}
	if params.LanguageCode != "" {
		updatedReport.LanguageCode = internal.LanguageCode(params.LanguageCode)
	}
	updatedReport.IsDefault = params.IsDefault

	userLoggedInID := userLoggedIn["uuid"].(uuid.UUID)
	updatedReportRepo.UpdatedBy = userLoggedInID

	updatedReportRepo, err = s.repository.UpdateReportCategory(ctx, updatedReportRepo)
	if err != nil {
		return nil, err
	}
	updatedReport.FromRepositoryModel(updatedReportRepo)

	return updatedReport.ToPublicModel(), nil
}
