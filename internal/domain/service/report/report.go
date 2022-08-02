package report

import (
	"context"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal/infrastructure/repository"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
)

// ReportServiceInterface represents the report service interface
type ReportServiceInterface interface {
	ListReports(ctx context.Context, params *public.ListReportRequest) ([]public.ReportResponse, error)
	GetReport(ctx context.Context, reportID uuid.UUID) (*public.ReportResponse, error)
	CreateReport(ctx context.Context, params *public.CreateReportRequest) (*public.ReportResponse, error)
	UpdateReport(ctx context.Context, params *public.UpdateReportRequest) (*public.ReportResponse, error)
	DeleteReport(ctx context.Context, params *public.DeleteReportRequest) error
}

// ReportService is the domain logic implementation of report service interface
type ReportService struct {
	repository repository.ReportRepository
}

// NewReportService creates a new report domain service
func NewReportService(
	repository repository.ReportRepository,
) ReportServiceInterface {
	return &ReportService{
		repository: repository,
	}
}
