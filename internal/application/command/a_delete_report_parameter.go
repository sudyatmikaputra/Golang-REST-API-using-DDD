package command

import (
	"context"

	reportParameterDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/report_parameter"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
)

type DeleteReportParameterForAdminCommand struct {
	reportParameterService reportParameterDomainService.ReportParameterServiceInterface
}

func NewDeleteReportParameterForAdminCommand(
	reportParameterService reportParameterDomainService.ReportParameterServiceInterface,
) DeleteReportParameterForAdminCommand {
	return DeleteReportParameterForAdminCommand{
		reportParameterService: reportParameterService,
	}
}

func (r DeleteReportParameterForAdminCommand) Execute(ctx context.Context, params public.DeleteReportParameterRequest) error {

	err := r.reportParameterService.DeleteReportParameter(ctx, &params)
	if err != nil {
		return err
	}

	return nil
}
