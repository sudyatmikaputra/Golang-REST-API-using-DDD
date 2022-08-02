package query

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	reportDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/report"
	reportCategoryDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/report_category"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

// GetReportQuery encapsulate process for getting a report in query
type GetReportQuery struct {
	reportService   reportDomainService.ReportServiceInterface
	categoryService reportCategoryDomainService.ReportCategoryServiceInterface
}

// NewGetReportQuery build an query for getting a report
func NewGetReportQuery(
	reportService reportDomainService.ReportServiceInterface,
	categoryService reportCategoryDomainService.ReportCategoryServiceInterface,
) GetReportQuery {
	return GetReportQuery{
		reportService:   reportService,
		categoryService: categoryService,
	}
}

func (r GetReportQuery) Execute(ctx context.Context, params public.GetReportRequest) (*public.ReportResponse, error) {
	report, err := r.reportService.GetReport(ctx, params.ReportID)
	if err != nil {
		return nil, err
	}
	if report == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	category, err := r.categoryService.GetReportCategory(ctx, report.ReportCategory.ID)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	report.ReportCategory = *category

	return report, nil
}
