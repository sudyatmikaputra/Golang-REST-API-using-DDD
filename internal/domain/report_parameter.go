package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/infrastructure/repository"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	"github.com/medicplus-inc/medicplus-kit/encoding"
)

type ReportParameter struct {
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

func (a *ReportParameter) FromPublicModel(reportParameterPublic interface{}) {
	_ = encoding.TransformObject(reportParameterPublic, a)
}

func (a *ReportParameter) ToPublicModel() *public.ReportParameterResponse {
	reportParameterPublic := &public.ReportParameterResponse{}
	_ = encoding.TransformObject(a, reportParameterPublic)
	return reportParameterPublic
}

func (a *ReportParameter) FromRepositoryModel(reportParameterRepo interface{}) {
	_ = encoding.TransformObject(reportParameterRepo, a)
}

func (a *ReportParameter) ToRepositoryModel() *repository.ReportParameter {
	reportParameterRepo := &repository.ReportParameter{}
	_ = encoding.TransformObject(a, reportParameterRepo)
	return reportParameterRepo
}
