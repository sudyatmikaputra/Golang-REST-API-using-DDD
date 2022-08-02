package query

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	reportDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/report"
	reportCategoryDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/report_category"
	"github.com/medicplus-inc/medicplus-feedback/internal/global"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

// ListReportsQuery encapsulate process for listing report in query
type ListReportsQuery struct {
	reportService   reportDomainService.ReportServiceInterface
	categoryService reportCategoryDomainService.ReportCategoryServiceInterface
}

// NewListReportsQuery build an query for listing report
func NewListReportsQuery(
	reportService reportDomainService.ReportServiceInterface,
	categoryService reportCategoryDomainService.ReportCategoryServiceInterface,
) ListReportsQuery {
	return ListReportsQuery{
		reportService:   reportService,
		categoryService: categoryService,
	}
}

func (r ListReportsQuery) Execute(ctx context.Context, params public.ListReportRequest) ([]public.ReportResponse, error) {
	userLoggedIn, _ := global.GetClaimsFromContext(ctx)
	userLoggedInID := userLoggedIn["uuid"].(uuid.UUID)
	userLoggedInRole := userLoggedIn["role"].(string)

	if userLoggedInID != params.ReportFromID {
		if userLoggedInRole != string(internal.Admin) {
			return nil, libError.New(internal.ErrNotAuthorized, http.StatusUnauthorized, internal.ErrNotAuthorized.Error())
		}
	}

	reports, err := r.reportService.ListReports(ctx, &params)
	if err != nil {
		return nil, err
	}
	if reports == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	// categoryIDs := []uuid.UUID{}
	// reportCategoryMaps := map[uuid.UUID]*public.ReportResponse{}
	// for _, report := range reports {
	// 	reportCategoryMaps[report.ReportCategory.ID] = &report
	// 	categoryIDs = append(categoryIDs, report.ReportCategory.ID)
	// }

	// allCategories, err := r.categoryService.ListReportCategories(ctx, &public.ListReportCategoryRequest{
	// 	IDs: categoryIDs,
	// })
	// if err != nil {
	// 	return nil, err
	// }

	// for _, category := range allCategories {
	// 	reportCategoryMaps[category.ID].ReportCategory = category
	// }

	// for _, report := range reports {
	// 	report.ReportCategory = reportCategoryMaps[report.ReportCategory.ID].ReportCategory
	// }

	return reports, nil
}
