package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
)

type Report struct {
	ID               uuid.UUID              `json:"id" gorm:"primaryKey,not null"`
	ReportTo         internal.ReceiverType  `json:"report_to" gorm:"not null"`
	ReportToID       uuid.UUID              `json:"report_to_id" gorm:"not null"`
	ReportFromID     uuid.UUID              `json:"report_from_id" gorm:"not null"`
	ReportCategoryID uuid.UUID              `json:"report_category_id" gorm:"not null"`
	Context          internal.ReportContext `json:"context"`
	ContextID        uuid.UUID              `json:"context_id"`
	CreatedAt        time.Time              `json:"created_at" gorm:"not null,autoCreateTime"`
	UpdatedAt        time.Time              `json:"updated_at" gorm:"not null,autoUpdateTime"`
}

type ReportRepository interface {
	FindAllReports(ctx context.Context, params *public.ListReportRequest) ([]Report, error)
	FindReportByID(ctx context.Context, reportID uuid.UUID) (*Report, error)
	InsertReport(ctx context.Context, report *Report) (*Report, error)
	UpdateReport(ctx context.Context, report *Report) (*Report, error)
	DeleteReport(ctx context.Context, report *Report) error
}
