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
	if params.ReportType != string(internal.ToMerchant) {
		return nil, libError.New(internal.ErrInvalidParameterType, http.StatusBadRequest, internal.ErrInvalidParameterType.Error())
	}
	if params.Context != string(internal.Purchase) {
		return nil, libError.New(internal.ErrInvalidContext, http.StatusBadRequest, internal.ErrInvalidContext.Error())
	}
	report, err := r.reportService.CreateReport(ctx, &public.CreateReportRequest{
		ReportType: params.ReportType,
		ReportToID: params.ReportToID,
		Context:    params.Context,
		ContextID:  params.ContextID,
		Notes:      params.Notes,
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
