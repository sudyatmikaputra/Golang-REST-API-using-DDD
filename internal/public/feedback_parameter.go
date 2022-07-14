package public

import (
	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
)

type FeedbackParameterResponse struct {
	ID            uuid.UUID              `json:"feedback_parameter"`
	ParameterType internal.ParameterType `json:"parameter_type"`
	Name          string                 `json:"name"`
	Language      string                 `json:"language"`
	IsDefault     bool                   `json:"is_default"`
}

type CreateFeedbackParameterRequest struct {
	ParameterType internal.ParameterType `json:"parameter_type" validate:"required"`
	Name          string                 `json:"name" validate:"required"`
	Language      string                 `json:"language" validate:"required"`
	IsDefault     bool                   `json:"is_default" validate:"required"`
}

type UpdateFeedbackParameterRequest struct {
	ID            uuid.UUID              `json:"feedback_parameter" validate:"required"`
	ParameterType internal.ParameterType `json:"parameter_type"`
	Name          string                 `json:"name"`
	Language      string                 `json:"language"`
	IsDefault     bool                   `json:"is_default"`
}

// ListFeedbackParameterRequest represents params to get List feedbacks parameter
type ListFeedbackParameterRequest struct {
	Search string      `json:"search"`
	ID     uuid.UUID   `json:"id"`
	IDs    []uuid.UUID `json:"ids"`
	Page   int         `json:"page" validate:"min=1"`
	Limit  int         `json:"limit"`
	Type   string      `json:"parameter_type"`
}

type GetFeedbackParameterRequest struct {
	ParameterID uuid.UUID `url_param:"parameter_id" validate:"required"`
}
