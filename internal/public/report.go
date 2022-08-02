package public

import (
	"github.com/google/uuid"
)

type ReportResponse struct {
	ID               uuid.UUID              `json:"id"`
	ReportTo         string                 `json:"report_to"`
	ReportCategoryID uuid.UUID              `json:"report_category_id"`
	ReportToID       uuid.UUID              `json:"report_to_id"`
	ReportFromID     uuid.UUID              `json:"report_from_id"`
	ReportCategory   ReportCategoryResponse `json:"report_category"`
	Context          string                 `json:"context"`
	ContextID        uuid.UUID              `json:"context_id"`
	Notes            string                 `json:"notes"`
}

type CreateReportRequest struct {
	ReportTo         string                 `json:"report_to" validate:"required"`
	ReportCategoryID uuid.UUID              `json:"report_category_id" validate:"required"`
	ReportToID       uuid.UUID              `json:"report_to_id" validate:"required"`
	ReportFromID     uuid.UUID              `json:"report_from_id" validate:"required"`
	ReportCategory   ReportCategoryResponse `json:"report_category" validate:"required"`
	Context          string                 `json:"context" validate:"required"`
	ContextID        uuid.UUID              `json:"context_id" validate:"required"`
	Notes            string                 `json:"notes" validate:"required"`
}

type UpdateReportRequest struct {
	ID    uuid.UUID `json:"id" validate:"required"`
	Notes string    `json:"notes" validate:"required"`
}

type ListReportRequest struct {
	Search           string    `qs:"search"` //Notes
	Page             int       `qs:"page"`
	Limit            int       `qs:"limit"`
	ReportCategoryID uuid.UUID `qs:"report_category_id" validate:"required"`
	ReportFromID     uuid.UUID `qs:"report_from_id"`
	ReportTo         string    `qs:"report_to" validate:"required"`
	ReportToID       uuid.UUID `qs:"report_to_id"`
}

type GetReportRequest struct {
	ReportID uuid.UUID `url_param:"report_id"`
}

type DeleteReportRequest struct {
	ReportID uuid.UUID `json:"report_id" validate:"required"`
}
