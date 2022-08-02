package command

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	reportDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/report"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

type UpdateReportForPatientDoctorCommand struct {
	reportService reportDomainService.ReportServiceInterface
}

func NewUpdateReportForPatientDoctorCommand(
	reportService reportDomainService.ReportServiceInterface,
) UpdateReportForPatientDoctorCommand {
	return UpdateReportForPatientDoctorCommand{
		reportService: reportService,
	}
}

func (r UpdateReportForPatientDoctorCommand) Execute(ctx context.Context, params public.UpdateReportRequest) (*public.ReportResponse, error) {
	report, err := r.reportService.UpdateReport(ctx, &params)
	if err != nil {
		return nil, err
	}
	if report == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	return report, nil
}
