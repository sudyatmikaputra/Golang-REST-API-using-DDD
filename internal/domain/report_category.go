package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/infrastructure/repository"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	"github.com/medicplus-inc/medicplus-kit/encoding"
)

type ReportCategory struct {
	ID           uuid.UUID              `json:"id" `
	ReportType   internal.ParameterType `json:"report_type"`
	Name         string                 `json:"name"`
	LanguageCode internal.LanguageCode  `json:"language_code"`
	IsDefault    bool                   `json:"is_default"`
	CreatedBy    uuid.UUID              `json:"created_by"`
	CreatedAt    time.Time              `json:"created_at"`
	UpdatedAt    time.Time              `json:"updated_at"`
	UpdatedBy    uuid.UUID              `json:"updated_by"`
	DeletedAt    *time.Time             `json:"deleted_at"`
	DeletedBy    *uuid.UUID             `json:"deleted_by"`
}

func (a *ReportCategory) FromPublicModel(categoryPublic interface{}) {
	_ = encoding.TransformObject(categoryPublic, a)
}

func (a *ReportCategory) ToPublicModel() *public.ReportCategoryResponse {
	categoryPublic := &public.ReportCategoryResponse{}
	_ = encoding.TransformObject(a, categoryPublic)
	return categoryPublic
}

func (a *ReportCategory) FromRepositoryModel(categoryRepo interface{}) {
	_ = encoding.TransformObject(categoryRepo, a)
}

func (a *ReportCategory) ToRepositoryModel() *repository.ReportCategory {
	categoryRepo := &repository.ReportCategory{}
	_ = encoding.TransformObject(a, categoryRepo)
	return categoryRepo
}
