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

type ListReportsForDoctorAnonymouslyQuery struct {
	reportService          reportDomainService.ReportServiceInterface
	reportParameterService reportParameterDomainService.ReportParameterServiceInterface
}

func NewListReportsForDoctorAnonymouslyQuery(
	reportService reportDomainService.ReportServiceInterface,
	reportParameterService reportParameterDomainService.ReportParameterServiceInterface,
) ListReportsForDoctorAnonymouslyQuery {
	return ListReportsForDoctorAnonymouslyQuery{
		reportService:          reportService,
		reportParameterService: reportParameterService,
	}
}

func (r ListReportsForDoctorAnonymouslyQuery) Execute(ctx context.Context, params public.ListReportRequest) ([]public.AnonymousReportResponse, error) {
	userLoggedIn, _ := global.GetClaimsFromContext(ctx)
	if params.ReportType == string(internal.DoctorParameterType) && userLoggedIn["uuid"].(uuid.UUID) != params.ReportToID {
		return nil, libError.New(internal.ErrNotAuthorized, http.StatusUnauthorized, internal.ErrNotAuthorized.Error())
	}
	if params.ReportType != string(internal.DoctorParameterType) && userLoggedIn["uuid"].(uuid.UUID) != params.ReportFromID {
		return nil, libError.New(internal.ErrNotAuthorized, http.StatusUnauthorized, internal.ErrNotAuthorized.Error())
	}

	reports, err := r.reportService.ListReportsAnonymously(ctx, &params)
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
