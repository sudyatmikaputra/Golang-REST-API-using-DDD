package reportcategory

import (
	"context"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/infrastructure/repository"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
)

// ReportCategoryServiceInterface represents the report category service interface
type ReportCategoryServiceInterface interface {
	ListReportCategories(ctx context.Context, params *public.ListReportCategoryRequest) ([]public.ReportCategoryResponse, error)
	GetReportCategory(ctx context.Context, categoryID uuid.UUID) (*public.ReportCategoryResponse, error)
	GetReportCategoryByReportType(ctx context.Context, reportType internal.ParameterType, languageCode string) (*public.ReportCategoryResponse, error)
	CreateReportCategory(ctx context.Context, params *public.CreateReportCategoryRequest) (*public.ReportCategoryResponse, error)
	UpdateReportCategory(ctx context.Context, params *public.UpdateReportCategoryRequest) (*public.ReportCategoryResponse, error)
	DeleteReportCategory(ctx context.Context, params *public.DeleteReportCategoryRequest) error
}

// ReportCategoryService is the domain logic implementation of report category service interface
type ReportCategoryService struct {
	repository repository.ReportCategoryRepository
}

// NewReportCategoryService creates a new report category domain service
func NewReportCategoryService(
	repository repository.ReportCategoryRepository,
) ReportCategoryServiceInterface {
	return &ReportCategoryService{
		repository: repository,
	}
}
