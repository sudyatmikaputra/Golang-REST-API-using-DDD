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

// CreateReportCategory creates a new report catgeory
func (s *ReportCategoryService) CreateReportCategory(ctx context.Context, params *public.CreateReportCategoryRequest) (*public.ReportCategoryResponse, error) {
	userLoggedIn, _ := global.GetClaimsFromContext(ctx)

	existingReportCategory, err := s.repository.FindReportCategoryByReportType(ctx, internal.ParameterType(params.ReportType), params.LanguageCode)
	if err != nil {
		return nil, err
	}
	if existingReportCategory != nil {
		return nil, libError.New(internal.ErrLanguageCodeAlreadyExists, http.StatusBadRequest, internal.ErrLanguageCodeAlreadyExists.Error())
	}

	category := &domain.ReportCategory{
		ReportType:   internal.ParameterType(params.ReportType),
		Name:         params.Name,
		LanguageCode: internal.LanguageCode(params.LanguageCode),
		IsDefault:    params.IsDefault,
	}

	categoryRepo := category.ToRepositoryModel()
	userLoggedInID := userLoggedIn["uuid"].(uuid.UUID)
	categoryRepo.CreatedBy = userLoggedInID
	categoryRepo.UpdatedBy = userLoggedInID

	insertedCategory, err := s.repository.InsertReportCategory(ctx, categoryRepo)
	if err != nil {
		return nil, err
	}
	if insertedCategory == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	category.FromRepositoryModel(insertedCategory)

	return category.ToPublicModel(), nil
}
