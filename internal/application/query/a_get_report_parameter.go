package query

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	reportParameterDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/report_parameter"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

type GetReportParameterForAdminQuery struct {
	reportParameterService reportParameterDomainService.ReportParameterServiceInterface
}

func NewGetReportParameterForAdminQuery(
	reportParameterService reportParameterDomainService.ReportParameterServiceInterface,
) GetReportParameterForAdminQuery {
	return GetReportParameterForAdminQuery{
		reportParameterService: reportParameterService,
	}
}

func (r GetReportParameterForAdminQuery) Execute(ctx context.Context, params public.GetReportParameterRequest) (*public.ReportParameterResponse, error) {
	reportParameter, err := r.reportParameterService.GetReportParameter(ctx, params.ID)
	if err != nil {
		return nil, err
	}
	if reportParameter == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	return reportParameter, nil
}
