package command

import (
	"context"

	reportParameterDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/report_parameter"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
)

type CreateReportParameterForAdminCommand struct {
	reportParameterService reportParameterDomainService.ReportParameterServiceInterface
}

func NewCreateReportParameterForAdminCommand(
	reportParameterService reportParameterDomainService.ReportParameterServiceInterface,
) CreateReportParameterForAdminCommand {
	return CreateReportParameterForAdminCommand{
		reportParameterService: reportParameterService,
	}
}

func (r CreateReportParameterForAdminCommand) Execute(ctx context.Context, params public.CreateReportParameterRequest) (*public.ReportParameterResponse, error) {
	reportParameter, err := r.reportParameterService.CreateReportParameter(ctx, &params)
	if err != nil {
		return nil, err
	}

	return reportParameter, nil
}
