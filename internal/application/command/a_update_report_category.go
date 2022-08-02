package command

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	reportCategoryDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/report_category"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

// UpdateReportCategoryCommand encapsulate process for updating report category in Command
type UpdateReportCategoryCommand struct {
	categoryService reportCategoryDomainService.ReportCategoryServiceInterface
}

// NewUpdateReportCategoryCommand build an Command for updating report category
func NewUpdateReportCategoryCommand(
	categoryService reportCategoryDomainService.ReportCategoryServiceInterface,
) UpdateReportCategoryCommand {
	return UpdateReportCategoryCommand{
		categoryService: categoryService,
	}
}

func (r UpdateReportCategoryCommand) Execute(ctx context.Context, params public.UpdateReportCategoryRequest) (*public.ReportCategoryResponse, error) {
	category, err := r.categoryService.UpdateReportCategory(ctx, &params)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	return category, nil
}
