package query

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	reportParameterDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/report_parameter"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

type ListReportParametersQuery struct {
	reportParameterService reportParameterDomainService.ReportParameterServiceInterface
}

func NewListReportParametersQuery(
	reportParameterService reportParameterDomainService.ReportParameterServiceInterface,
) ListReportParametersQuery {
	return ListReportParametersQuery{
		reportParameterService: reportParameterService,
	}
}

func (r ListReportParametersQuery) Execute(ctx context.Context, params public.ListReportParameterRequest) ([]public.ReportParameterResponse, error) {
	categories, err := r.reportParameterService.ListReportParameters(ctx, &params)
	if err != nil {
		return nil, err
	}
	if categories == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	return categories, nil
}
