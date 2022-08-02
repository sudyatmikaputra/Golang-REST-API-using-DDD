package query

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	reportDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/report"
	reportParameterDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/report_parameter"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

type GetReportQuery struct {
	reportService          reportDomainService.ReportServiceInterface
	reportParameterService reportParameterDomainService.ReportParameterServiceInterface
}

func NewGetReportQuery(
	reportService reportDomainService.ReportServiceInterface,
	reportParameterService reportParameterDomainService.ReportParameterServiceInterface,
) GetReportQuery {
	return GetReportQuery{
		reportService:          reportService,
		reportParameterService: reportParameterService,
	}
}

func (r GetReportQuery) Execute(ctx context.Context, params public.GetReportRequest) (*public.ReportResponse, error) {
	report, err := r.reportService.GetReport(ctx, params.ID)
	if err != nil {
		return nil, err
	}
	if report == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	reportParameter, err := r.reportParameterService.GetReportParameterByReportType(ctx, internal.ParameterType(report.ReportType), params.LanguageCode)
	if err != nil {
		return nil, err
	}
	report.ReportParameter = *reportParameter

	return report, nil
}
