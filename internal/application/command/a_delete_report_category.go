package command

import (
	"context"

	reportCategoryDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/report_category"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
)

// DeleteReportCategoryCommand encapsulate process for deleting report category in Command
type DeleteReportCategoryCommand struct {
	categoryService reportCategoryDomainService.ReportCategoryServiceInterface
}

// NewDeleteReportCategoryCommand build an Command for deleting report category
func NewDeleteReportCategoryCommand(
	categoryService reportCategoryDomainService.ReportCategoryServiceInterface,
) DeleteReportCategoryCommand {
	return DeleteReportCategoryCommand{
		categoryService: categoryService,
	}
}

func (r DeleteReportCategoryCommand) Execute(ctx context.Context, params public.DeleteReportCategoryRequest) error {

	err := r.categoryService.DeleteReportCategory(ctx, &params)
	if err != nil {
		return err
	}

	return nil
}
