package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/infrastructure/repository"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	"github.com/medicplus-inc/medicplus-kit/encoding"
)

type Report struct {
	ID               uuid.UUID              `json:"id"`
	ReportTo         internal.ReceiverType  `json:"report_to"`
	ReportToID       uuid.UUID              `json:"report_to_id"`
	ReportFromID     uuid.UUID              `json:"report_from_id"`
	ReportCategoryID uuid.UUID              `json:"report_category_id"`
	Context          internal.ReportContext `json:"context"`
	ContextID        uuid.UUID              `json:"context_id"`
	Notes            string                 `json:"notes"`
	CreatedAt        time.Time              `json:"created_at"`
	UpdatedAt        time.Time              `json:"updated_at"`
}

func (a *Report) FromPublicModel(reportPublic interface{}) {
	_ = encoding.TransformObject(reportPublic, a)
}

func (a *Report) ToPublicModel() *public.ReportResponse {
	reportPublic := &public.ReportResponse{}
	_ = encoding.TransformObject(a, reportPublic)
	return reportPublic
}

func (a *Report) FromRepositoryModel(reportRepo interface{}) {
	_ = encoding.TransformObject(reportRepo, a)
}

func (a *Report) ToRepositoryModel() *repository.Report {
	reportRepo := &repository.Report{}
	_ = encoding.TransformObject(a, reportRepo)
	return reportRepo
}
