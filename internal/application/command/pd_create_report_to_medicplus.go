package command

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	reportDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/report"
	reportCategoryDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/report_category"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

// CreateReportCommand encapsulate process for creating report in Command
type CreateReportCommand struct {
	reportService   reportDomainService.ReportServiceInterface
	categoryService reportCategoryDomainService.ReportCategoryServiceInterface
}

// NewCreateReportCommand build an Command for creating report
func NewCreateReportCommand(
	reportService reportDomainService.ReportServiceInterface,
	categoryService reportCategoryDomainService.ReportCategoryServiceInterface,
) CreateReportCommand {
	return CreateReportCommand{
		reportService:   reportService,
		categoryService: categoryService,
	}
}

func (r CreateReportCommand) ExecuteToMedicplus(ctx context.Context, params public.CreateReportRequest) (*public.ReportResponse, error) {
	// category, err := r.categoryService.GetReportCategory(ctx, params.ReportCategory.ID)
	// if err != nil {
	// 	return nil, err
	// }
	// if category == nil {
	// 	return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	// }

	report, err := r.reportService.CreateReport(ctx, &public.CreateReportRequest{
		ReportTo:         string(internal.ToMedicplus),
		ReportCategoryID: params.ReportCategoryID,
		ReportToID:       params.ReportToID,
		ReportFromID:     params.ReportFromID,
		Context:          string(internal.Purchase),
		ContextID:        params.ContextID,
		Notes:            params.Notes,
	})
	if err != nil {
		return nil, err
	}
	if report == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	// report.ReportCategory = *category

	return report, nil
}
