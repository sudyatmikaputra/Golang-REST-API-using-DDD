package query

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	reportDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/report"
	reportParameterDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/report_parameter"
	"github.com/medicplus-inc/medicplus-feedback/internal/global"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

type ListReportsForPatientQuery struct {
	reportService          reportDomainService.ReportServiceInterface
	reportParameterService reportParameterDomainService.ReportParameterServiceInterface
}

func NewListReportsForPatientQuery(
	reportService reportDomainService.ReportServiceInterface,
	reportParameterService reportParameterDomainService.ReportParameterServiceInterface,
) ListReportsForPatientQuery {
	return ListReportsForPatientQuery{
		reportService:          reportService,
		reportParameterService: reportParameterService,
	}
}

func (r ListReportsForPatientQuery) Execute(ctx context.Context, params public.ListReportRequest) ([]public.ReportResponse, error) {
	userLoggedIn, _ := global.GetClaimsFromContext(ctx)
	if userLoggedIn["uuid"].(uuid.UUID) != params.ReportFromID {
		return nil, libError.New(internal.ErrNotAuthorized, http.StatusUnauthorized, internal.ErrNotAuthorized.Error())
	}

	reports, err := r.reportService.ListReports(ctx, &params)
	if err != nil {
		return nil, err
	}
	if reports == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	reportParameter, err := r.reportParameterService.GetReportParameterByReportType(ctx, internal.ParameterType(params.ReportType), params.LanguageCode)
	if err != nil {
		return nil, err
	}
	if reportParameter != nil {
		for _, report := range reports {
			report.ReportParameter = *reportParameter
		}
	}

	return reports, nil
}
