package command

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	reportParameterDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/report_parameter"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

type UpdateReportParameterForAdminCommand struct {
	reportParameterService reportParameterDomainService.ReportParameterServiceInterface
}

func NewUpdateReportParameterForAdminCommand(
	reportParameterService reportParameterDomainService.ReportParameterServiceInterface,
) UpdateReportParameterForAdminCommand {
	return UpdateReportParameterForAdminCommand{
		reportParameterService: reportParameterService,
	}
}

func (r UpdateReportParameterForAdminCommand) Execute(ctx context.Context, params public.UpdateReportParameterRequest) (*public.ReportParameterResponse, error) {
	parameter, err := r.reportParameterService.UpdateReportParameter(ctx, &params)
	if err != nil {
		return nil, err
	}
	if parameter == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	return parameter, nil
}
