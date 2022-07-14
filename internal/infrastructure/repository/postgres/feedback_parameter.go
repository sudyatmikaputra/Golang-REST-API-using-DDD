package postgres

import (
	"context"
	"time"

	"github.com/medicplus-inc/medicplus-feedback/config"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal/infrastructure/repository"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	"github.com/medicplus-inc/medicplus-kit/database"
	"gorm.io/gorm"
)

// parameterPostgres implements the feedback parameter repository service interface
type parameterPostgres struct {
	db *gorm.DB
}

// FindAllParameter queries all feedback parameters
func (s *parameterPostgres) FindAllParameter(ctx context.Context, params *public.ListFeedbackParameterRequest) ([]*repository.FeedbackParameter, error) {
	db := s.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	var feedbacks []repository.FeedbackParameter
	args := []interface{}{}
	where := `"deleted_at" IS NULL`
	if params.Search != "" {
		where += ` AND ("name" ILIKE ? OR "parameter_type" ILIKE ?)`
		args = append(args, "%"+params.Search+"%")
		args = append(args, "%"+params.Search+"%")
	}
	if params.ID != uuid.Nil {
		where += ` AND "id" = ?`
		args = append(args, params.ID)
	}
	if len(params.IDs) != 0 {
		where += ` AND "ids" IN ?`
		args = append(args, params.IDs)
	}
	if params.Type != "" {
		where += ` AND "parameter_type" ILIKE ?`
		args = append(args, params.Type)
	}

	order := `"id" DESC`
	if err := db.Where(
		where,
		args...,
	).
		Order(order).
		Offset(((params.Page - 1) * params.Limit)).
		Limit(params.Limit).
		Find(&feedbacks).Error; err != nil {
		return nil, err
	}

	var result []*repository.FeedbackParameter
	for _, _feedback := range feedbacks {
		result = append(result, &repository.FeedbackParameter{
			ID:            _feedback.ID,
			ParameterType: _feedback.ParameterType,
			Name:          _feedback.Name,
			Language:      _feedback.Language,
			IsDefault:     _feedback.IsDefault,
			CreatedBy:     _feedback.CreatedBy,
			CreatedAt:     _feedback.CreatedAt,
			UpdatedAt:     _feedback.UpdatedAt,
			UpdatedBy:     _feedback.UpdatedBy,
			DeletedAt:     _feedback.DeletedAt,
			DeletedBy:     _feedback.DeletedBy,
		})
	}

	return result, nil
}

// FindByParameterID finds feedback by its id
func (s *parameterPostgres) FindByParameterID(ctx context.Context, parameterID uuid.UUID) (*repository.FeedbackParameter, error) {
	db := s.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	var feedback repository.FeedbackParameter
	if err := db.First(&feedback, `"id" = ? AND "deleted_at" IS NULL`, parameterID).Error; err != nil {
		return nil, err
	}

	return &feedback, nil
}

// DeleteParameter deletes a feedback parameter based on its id
func (s *parameterPostgres) DeleteParameter(ctx context.Context, parameter *repository.FeedbackParameter) error {
	db := s.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	if err := db.Delete(parameter).Error; err != nil {
		return err
	}

	return nil
}

// Insert inserts feedback parameter
func (d *parameterPostgres) InsertParameter(ctx context.Context, parameter *repository.FeedbackParameter) (*repository.FeedbackParameter, error) {
	db := d.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	parameter.ID, _ = uuid.NewRandom()
	if parameter.CreatedBy == uuid.Nil {
		parameter.CreatedBy = parameter.ID
		parameter.UpdatedBy = parameter.ID
	}
	parameter.CreatedAt = time.Now().UTC()
	parameter.UpdatedAt = time.Now().UTC()

	var err error

	if err = db.Create(parameter).Error; err != nil {
		return nil, err
	}

	return parameter, nil
}

// UpdateParameter updates feedback parameter
func (d *parameterPostgres) UpdateParameter(ctx context.Context, parameter *repository.FeedbackParameter) (*repository.FeedbackParameter, error) {
	db := d.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	var err error

	if err = db.Save(parameter).Error; err != nil {
		return nil, err
	}

	return parameter, nil
}

// NewParameterPostgres creates new feedback parameter repository
func NewParameterPostgres() repository.FeedbackParameterRepository {
	return &parameterPostgres{
		db: config.DB(),
	}
}

// NewParameterPostgresMock creates new feedback parameter repository mock for testing purpose only
func NewParameterPostgresMock(db *gorm.DB) repository.FeedbackParameterRepository {
	return &parameterPostgres{
		db: db,
	}
}
