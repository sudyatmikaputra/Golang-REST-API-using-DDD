package query

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	reportCategoryDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/report_category"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

// ListReportCategoriesQuery encapsulate process for listing report categories in query
type ListReportCategoriesQuery struct {
	categoryService reportCategoryDomainService.ReportCategoryServiceInterface
}

// NewListReportCategoriesQuery build an query for listing report categories
func NewListReportCategoriesQuery(
	categoryService reportCategoryDomainService.ReportCategoryServiceInterface,
) ListReportCategoriesQuery {
	return ListReportCategoriesQuery{
		categoryService: categoryService,
	}
}

func (r ListReportCategoriesQuery) Execute(ctx context.Context, params public.ListReportCategoryRequest) ([]public.ReportCategoryResponse, error) {
	categories, err := r.categoryService.ListReportCategories(ctx, &params)
	if err != nil {
		return nil, err
	}
	if categories == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	return categories, nil
}
