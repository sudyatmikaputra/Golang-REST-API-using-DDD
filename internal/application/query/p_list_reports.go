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
	userLoggedInID := uuid.MustParse(userLoggedIn["uuid"].(string))

	if userLoggedInID != params.ReportFromID {
		return nil, libError.New(internal.ErrNotAuthorized, http.StatusUnauthorized, internal.ErrNotAuthorized.Error())
	}

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
