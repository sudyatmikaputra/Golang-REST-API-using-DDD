package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
)

type ReportParameter struct {
	ID           uuid.UUID              `json:"id" gorm:"primaryKey,not null"`
	ReportType   internal.ParameterType `json:"report_type" gorm:"index:idx_report_parameter_report_type,not null"`
	Name         string                 `json:"name" gorm:"index:idx_report_parameter_name,not null"`
	LanguageCode internal.LanguageCode  `json:"language_code" gorm:"index:idx_report_parameter_language_code,not null"`
	IsDefault    bool                   `json:"is_default" gorm:"not null"`
	CreatedBy    uuid.UUID              `json:"created_by" gorm:"not null"`
	CreatedAt    time.Time              `json:"created_at" gorm:"not null,autoCreateTime"`
	UpdatedAt    time.Time              `json:"updated_at" gorm:"not null,autoUpdateTime"`
	UpdatedBy    uuid.UUID              `json:"updated_by" gorm:"not null"`
	DeletedAt    *time.Time             `json:"deleted_at"`
	DeletedBy    *uuid.UUID             `json:"deleted_by"`
}

type ReportParameterRepository interface {
	FindAllReportParameters(ctx context.Context, params *public.ListReportParameterRequest) ([]ReportParameter, error)
	FindReportParameterByID(ctx context.Context, reportParameterID uuid.UUID) (*ReportParameter, error)
	FindReportParameterByReportType(ctx context.Context, reportType internal.ParameterType, languageCode string) (*ReportParameter, error)
	InsertReportParameter(ctx context.Context, reportParameter *ReportParameter) (*ReportParameter, error)
	UpdateReportParameter(ctx context.Context, reportParameter *ReportParameter) (*ReportParameter, error)
	DeleteReportParameter(ctx context.Context, reportParameter *ReportParameter) error
}
