package public

import (
	"github.com/google/uuid"
)

type ReportParameterResponse struct {
	ID           uuid.UUID `json:"id" `
	ReportType   string    `json:"report_type"`
	Name         string    `json:"name"`
	LanguageCode string    `json:"language_code"`
	IsDefault    bool      `json:"is_default"`
}

type CreateReportParameterRequest struct {
	ReportType   string `json:"report_type" validate:"required,oneof=all doctor merchant medicplus"`
	Name         string `json:"name" validate:"required, alpha"`
	LanguageCode string `json:"language_code" validate:"required"`
	IsDefault    bool   `json:"is_default" validate:"required"`
}

type UpdateReportParameterRequest struct {
	ID           uuid.UUID `json:"id" validate:"required"`
	ReportType   string    `json:"report_type" validate:"required,oneof=all doctor merchant medicplus"`
	Name         string    `json:"name" validate:"required"`
	LanguageCode string    `json:"language_code" validate:"required"`
	IsDefault    bool      `json:"is_default" validate:"required"`
}

type ListReportParameterRequest struct {
	Search       string `qs:"search"` //name
	Page         int    `qs:"page"`
	Limit        int    `qs:"limit"`
	ReportType   string `qs:"report_type"`
	LanguageCode string `qs:"language_code" validate:"required,oneof=id en"`
	IsDefault    *bool  `qs:"is_default"`
}

type GetReportParameterRequest struct {
	ID uuid.UUID `url_param:"id" validate:"required"`
}

type DeleteReportParameterRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}
