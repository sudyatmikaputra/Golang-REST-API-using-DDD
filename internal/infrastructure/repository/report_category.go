package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
)

type ReportCategory struct {
	ID           uuid.UUID              `json:"id" gorm:"primaryKey,not null"`
	ReportType   internal.ParameterType `json:"report_type" gorm:"not null"`
	Name         string                 `json:"name" gorm:"index:idx_report_name,not null"`
	LanguageCode internal.LanguageCode  `json:"language_code" gorm:"index:idx_report_language_code,not null"`
	IsDefault    bool                   `json:"is_default" gorm:"not null"`
	CreatedBy    uuid.UUID              `json:"created_by" gorm:"not null"`
	CreatedAt    time.Time              `json:"created_at" gorm:"not null,autoCreateTime"`
	UpdatedAt    time.Time              `json:"updated_at" gorm:"not null,autoUpdateTime"`
	UpdatedBy    uuid.UUID              `json:"updated_by" gorm:"not null"`
	DeletedAt    *time.Time             `json:"deleted_at"`
	DeletedBy    *uuid.UUID             `json:"deleted_by"`
}

type ReportCategoryRepository interface {
	FindAllReportCategories(ctx context.Context, params *public.ListReportCategoryRequest) ([]ReportCategory, error)
	FindReportCategoryByID(ctx context.Context, reportCategoryID uuid.UUID) (*ReportCategory, error)
	FindReportCategoryByReportType(ctx context.Context, reportType internal.ParameterType, languageCode string) (*ReportCategory, error)
	InsertReportCategory(ctx context.Context, reportCategory *ReportCategory) (*ReportCategory, error)
	UpdateReportCategory(ctx context.Context, reportCategory *ReportCategory) (*ReportCategory, error)
	DeleteReportCategory(ctx context.Context, reportCategory *ReportCategory) error
}
