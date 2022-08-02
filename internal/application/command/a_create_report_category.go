package command

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	reportCategoryDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/report_category"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

// CreateReportCategoryCommand encapsulate process for creating report category in Command
type CreateReportCategoryCommand struct {
	categoryService reportCategoryDomainService.ReportCategoryServiceInterface
}

// NewCreateReportCategoryCommand build an Command for creating report category
func NewCreateReportCategoryCommand(
	categoryService reportCategoryDomainService.ReportCategoryServiceInterface,
) CreateReportCategoryCommand {
	return CreateReportCategoryCommand{
		categoryService: categoryService,
	}
}

func (r CreateReportCategoryCommand) Execute(ctx context.Context, params public.CreateReportCategoryRequest) (*public.ReportCategoryResponse, error) {
	category, err := r.categoryService.CreateReportCategory(ctx, &params)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	return category, nil
}
