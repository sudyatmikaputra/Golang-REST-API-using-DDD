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

type ListReportsForAdminQuery struct {
	reportService          reportDomainService.ReportServiceInterface
	reportParameterService reportParameterDomainService.ReportParameterServiceInterface
}

func NewListReportsForAdminQuery(
	reportService reportDomainService.ReportServiceInterface,
	reportParameterService reportParameterDomainService.ReportParameterServiceInterface,
) ListReportsForAdminQuery {
	return ListReportsForAdminQuery{
		reportService:          reportService,
		reportParameterService: reportParameterService,
	}
}

func (r ListReportsForAdminQuery) Execute(ctx context.Context, params public.ListReportRequest) ([]public.ReportResponse, error) {
	reports, err := r.reportService.ListReports(ctx, &params)
	if err != nil {
		return nil, err
	}
	if reports == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	reportParameter, _ := r.reportParameterService.GetReportParameterByReportType(ctx, internal.ParameterType(params.ReportType), params.LanguageCode)
	if reportParameter != nil {
		for i := range reports {
			reports[i].ReportParameter = public.ReportParameterResponse{
				ID:           reportParameter.ID,
				ReportType:   reportParameter.ReportType,
				Name:         reportParameter.Name,
				LanguageCode: reportParameter.LanguageCode,
				IsDefault:    reportParameter.IsDefault,
			}
		}
	}

	return reports, nil
}
