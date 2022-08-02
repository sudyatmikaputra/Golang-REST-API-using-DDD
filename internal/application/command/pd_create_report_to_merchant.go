package command

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	reportDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/report"
	reportParameterDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/report_parameter"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

type CreateReportForPatientDoctorToMerchantCommand struct {
	reportService          reportDomainService.ReportServiceInterface
	reportParameterService reportParameterDomainService.ReportParameterServiceInterface
}

func NewCreateReportForPatientDoctorToMerchantCommand(
	reportService reportDomainService.ReportServiceInterface,
	reportParameterService reportParameterDomainService.ReportParameterServiceInterface,
) CreateReportForPatientDoctorToMerchantCommand {
	return CreateReportForPatientDoctorToMerchantCommand{
		reportService:          reportService,
		reportParameterService: reportParameterService,
	}
}

func (r CreateReportForPatientDoctorToMerchantCommand) Execute(ctx context.Context, params public.CreateReportRequest) (*public.ReportResponse, error) {
	report, err := r.reportService.CreateReport(ctx, &public.CreateReportRequest{
		ReportType:   string(internal.ToMerchant),
		ReportToID:   params.ReportToID,
		ReportFromID: params.ReportFromID,
		Context:      string(internal.Purchase),
		ContextID:    params.ContextID,
		Notes:        params.Notes,
	})
	if err != nil {
		return nil, err
	}
	if report == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	reportParameter, err := r.reportParameterService.GetReportParameterByReportType(ctx, internal.ParameterType(params.ReportType), string(internal.BahasaIndonesia))
	if err != nil {
		return nil, err
	}

	report.ReportParameter = *reportParameter

	return report, nil
}
