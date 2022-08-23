package public

import (
	"github.com/google/uuid"
)

type ReportResponse struct {
	ID              uuid.UUID               `json:"id"`
	ReportType      string                  `json:"report_type"`
	ReportToID      uuid.UUID               `json:"report_to_id"`
	ReportFromID    uuid.UUID               `json:"report_from_id"`
	ReportParameter ReportParameterResponse `json:"report_parameter"`
	Context         string                  `json:"context"`
	ContextID       uuid.UUID               `json:"context_id"`
	Notes           string                  `json:"notes"`
}

type AnonymousReportResponse struct {
	ID              uuid.UUID               `json:"id"`
	ReportType      string                  `json:"report_type"`
	ReportToID      uuid.UUID               `json:"report_to_id"`
	ReportParameter ReportParameterResponse `json:"report_parameter"`
	Context         string                  `json:"context"`
	ContextID       uuid.UUID               `json:"context_id"`
	Notes           string                  `json:"notes"`
}

type CreateReportRequest struct {
	ReportType string    `json:"report_type" validate:"required,oneof=all doctor merchant medicplus"`
	ReportToID uuid.UUID `json:"report_to_id" validate:"required"`
	Context    string    `json:"context" validate:"required"`
	ContextID  uuid.UUID `json:"context_id" validate:"required"`
	Notes      string    `json:"notes" validate:"required"`
}

type UpdateReportRequest struct {
	ID    uuid.UUID `json:"id" validate:"required"`
	Notes string    `json:"notes" validate:"required"`
}

type ListReportRequest struct {
	Search       string    `qs:"search"` //Notes
	Page         int       `qs:"page"`
	Limit        int       `qs:"limit"`
	ReportFromID uuid.UUID `qs:"report_from_id"`
	ReportToID   uuid.UUID `qs:"report_to_id"`
	ReportType   string    `qs:"report_type" validate:"required,oneof=all doctor merchant medicplus"`
	LanguageCode string    `qs:"language_code" validate:"required,oneof=id en"`
}

type GetReportRequest struct {
	ID           uuid.UUID `qs:"id" validate:"required"`
	LanguageCode string    `qs:"language_code" validate:"required,oneof=id en"`
}

type DeleteReportRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}
