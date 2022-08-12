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

type CreateReportForPatientToDoctorCommand struct {
	reportService          reportDomainService.ReportServiceInterface
	reportParameterService reportParameterDomainService.ReportParameterServiceInterface
}

// NewCreateReportCommand build an Command for creating report
func NewCreateReportForPatientToDoctorCommand(
	reportService reportDomainService.ReportServiceInterface,
	reportParameterService reportParameterDomainService.ReportParameterServiceInterface,
) CreateReportForPatientToDoctorCommand {
	return CreateReportForPatientToDoctorCommand{
		reportService:          reportService,
		reportParameterService: reportParameterService,
	}
}

func (r CreateReportForPatientToDoctorCommand) Execute(ctx context.Context, params public.CreateReportRequest) (*public.ReportResponse, error) {
	if params.ReportType != string(internal.ToDoctor) {
		return nil, libError.New(internal.ErrInvalidParameterType, http.StatusBadRequest, internal.ErrInvalidParameterType.Error())
	}
	if params.Context != string(internal.Consultation) {
		return nil, libError.New(internal.ErrInvalidContext, http.StatusBadRequest, internal.ErrInvalidContext.Error())
	}
	report, err := r.reportService.CreateReport(ctx, &public.CreateReportRequest{
		ReportType:   params.ReportType,
		ReportToID:   params.ReportToID,
		ReportFromID: params.ReportFromID,
		Context:      params.Context,
		ContextID:    params.ContextID,
		Notes:        params.Notes,
	})
	if err != nil {
		return nil, err
	}
	if report == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	reportParameter, _ := r.reportParameterService.GetReportParameterByReportType(ctx, internal.ParameterType(params.ReportType), string(internal.BahasaIndonesia))
	if reportParameter != nil {
		report.ReportParameter = public.ReportParameterResponse{
			ID:           reportParameter.ID,
			ReportType:   reportParameter.ReportType,
			Name:         reportParameter.Name,
			LanguageCode: reportParameter.LanguageCode,
			IsDefault:    reportParameter.IsDefault,
		}
	} else {
		return nil, libError.New(internal.ErrParameterNotFound, http.StatusBadRequest, internal.ErrParameterNotFound.Error())
	}

	return report, nil
}
