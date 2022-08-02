package public

import (
	"github.com/google/uuid"
)

type FeedbackParameterResponse struct {
	ID            uuid.UUID `json:"id"`
	ParameterType string    `json:"parameter_type"`
	Name          string    `json:"name"`
	LanguageCode  string    `json:"language_code"`
	IsDefault     bool      `json:"is_default"`
}

type CreateFeedbackParameterRequest struct {
	ParameterType string `json:"parameter_type" validate:"required"`
	Name          string `json:"name" validate:"required"`
	LanguageCode  string `json:"language_code" validate:"required,oneof=id en"`
	IsDefault     bool   `json:"is_default" validate:"required"`
}

type UpdateFeedbackParameterRequest struct {
	ID            uuid.UUID `json:"id" validate:"required"`
	ParameterType string    `json:"parameter_type" validate:"required,oneof=all doctor merchant medicplus"`
	Name          string    `json:"name" validate:"required"`
	LanguageCode  string    `json:"language_code" validate:"required"`
	IsDefault     bool      `json:"is_default" validate:"required"`
}

type ListFeedbackParameterRequest struct {
	Search        string `qs:"search"` //name
	Page          int    `qs:"page" validate:"min=1"`
	Limit         int    `qs:"limit"`
	ParameterType string `qs:"parameter_type" validate:"required,oneof=all doctor merchant medicplus"`
	LanguageCode  string `qs:"language_code" validate:"required,oneof=id en"`
	IsDefault     *bool  `qs:"is_default"`
}

type GetFeedbackParameterRequest struct {
	FeedbackParameterID uuid.UUID `url_param:"feedback_parameter_id" validate:"required"`
}

type DeleteFeedbackParameterRequest struct {
	FeedbackParameterID uuid.UUID `json:"feedback_parameter_id" validate:"required"`
}
