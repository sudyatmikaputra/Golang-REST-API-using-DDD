package report_parameter

import (
	"context"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/infrastructure/repository"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
)

// ReportParameterServiceInterface represents the report parameter service interface
type ReportParameterServiceInterface interface {
	ListReportParameters(ctx context.Context, params *public.ListReportParameterRequest) ([]public.ReportParameterResponse, error)
	GetReportParameter(ctx context.Context, reportParameterID uuid.UUID) (*public.ReportParameterResponse, error)
	GetReportParameterByReportType(ctx context.Context, reportType internal.ParameterType, languageCode string) (*public.ReportParameterResponse, error)
	CreateReportParameter(ctx context.Context, params *public.CreateReportParameterRequest) (*public.ReportParameterResponse, error)
	UpdateReportParameter(ctx context.Context, params *public.UpdateReportParameterRequest) (*public.ReportParameterResponse, error)
	DeleteReportParameter(ctx context.Context, params *public.DeleteReportParameterRequest) error
}

// ReportParameterService is the domain logic implementation of report parameter service interface
type ReportParameterService struct {
	repository repository.ReportParameterRepository
}

// NewReportParameterService creates a new report parameter domain service
func NewReportParameterService(
	repository repository.ReportParameterRepository,
) ReportParameterServiceInterface {
	return &ReportParameterService{
		repository: repository,
	}
}
