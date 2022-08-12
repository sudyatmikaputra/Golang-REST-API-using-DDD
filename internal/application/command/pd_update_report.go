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

type UpdateReportForPatientDoctorCommand struct {
	reportService          reportDomainService.ReportServiceInterface
	reportParameterService reportParameterDomainService.ReportParameterServiceInterface
}

func NewUpdateReportForPatientDoctorCommand(
	reportService reportDomainService.ReportServiceInterface,
	reportParameterService reportParameterDomainService.ReportParameterServiceInterface,
) UpdateReportForPatientDoctorCommand {
	return UpdateReportForPatientDoctorCommand{
		reportService:          reportService,
		reportParameterService: reportParameterService,
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

	reportParameter, _ := r.reportParameterService.GetReportParameterByReportType(ctx, internal.ParameterType(report.ReportType), string(internal.BahasaIndonesia))
	if reportParameter != nil {
		report.ReportParameter = public.ReportParameterResponse{
			ID:           reportParameter.ID,
			ReportType:   reportParameter.ReportType,
			Name:         reportParameter.Name,
			LanguageCode: reportParameter.LanguageCode,
			IsDefault:    reportParameter.IsDefault,
		}
	}

	return report, nil
}
