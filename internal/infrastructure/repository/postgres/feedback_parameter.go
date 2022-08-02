package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/config"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/infrastructure/repository"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	"github.com/medicplus-inc/medicplus-kit/database"
	"gorm.io/gorm"
)

// feedbackParameterPostgres implements the feedback parameter repository service interface
type feedbackParameterPostgres struct {
	db *gorm.DB
}

// FindAllParameter queries all feedback parameters
func (s *feedbackParameterPostgres) FindAllFeedbackParameters(ctx context.Context, params *public.ListFeedbackParameterRequest) ([]repository.FeedbackParameter, error) {
	db := s.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	feedbackParameters := []repository.FeedbackParameter{}
	args := []interface{}{}
	where := `"deleted_at" IS NULL`
	if params.Search != "" {
		where += ` AND ("name" ILIKE ?)`
		args = append(args, "%"+params.Search+"%")
	}
	if params.ParameterType != "" {
		where += ` AND "parameter_type" = ? `
		args = append(args, params.ParameterType)
	}

	if params.IsDefault != nil {
		where += ` AND "is_default" = ?`
		args = append(args, params.IsDefault)
	}

	order := `"created_at" DESC`
	if err := db.Where(
		where,
		args...,
	).
		Order(order).
		Offset(((params.Page - 1) * params.Limit)).
		Limit(params.Limit).
		Find(&feedbackParameters).Error; err != nil {
		return nil, err
	}

	return feedbackParameters, nil
}

// FindByParameterID finds feedback by its id
func (s *feedbackParameterPostgres) FindFeedbackParameterByID(ctx context.Context, feedbackParameterID uuid.UUID) (*repository.FeedbackParameter, error) {
	db := s.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	feedbackParameter := repository.FeedbackParameter{}
	err := db.First(&feedbackParameter, `"deleted_at" IS NULL AND id" = ? `, feedbackParameterID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &feedbackParameter, nil
}

func (s *feedbackParameterPostgres) FindFeedbackParameterByParameterType(ctx context.Context, parameterType internal.ParameterType, languageCode string) (*repository.FeedbackParameter, error) {
	db := s.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	if languageCode == "" {
		languageCode = string(internal.BahasaIndonesia)
	}

	feedbackParameter := repository.FeedbackParameter{}
	err := db.First(&feedbackParameter, `"deleted_at" IS NULL AND "parameter_type" = ? AND "language_code" = ? `, parameterType, languageCode).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &feedbackParameter, nil
}

// Insert inserts feedback parameter
func (d *feedbackParameterPostgres) InsertFeedbackParameter(ctx context.Context, feedbackParameter *repository.FeedbackParameter) (*repository.FeedbackParameter, error) {
	db := d.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	feedbackParameter.ID, _ = uuid.NewRandom()

	// now := time.Now().UTC()
	// feedbackParameter.CreatedAt = now
	// feedbackParameter.UpdatedAt = now

	err := db.Create(feedbackParameter).Error
	if err != nil {
		return nil, err
	}

	return feedbackParameter, nil
}

// UpdateParameter updates feedback parameter
func (d *feedbackParameterPostgres) UpdateFeedbackParameter(ctx context.Context, feedbackParameter *repository.FeedbackParameter) (*repository.FeedbackParameter, error) {
	db := d.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}
	// now := time.Now().UTC()
	// feedbackParameter.UpdatedAt = now
	err := db.Save(feedbackParameter).Error
	if err != nil {
		return nil, err
	}

	return feedbackParameter, nil
}

// DeleteParameter deletes a feedback parameter based on its id
func (s *feedbackParameterPostgres) DeleteFeedbackParameter(ctx context.Context, feedbackParameter *repository.FeedbackParameter) error {
	db := s.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	now := time.Now().UTC()
	feedbackParameter.DeletedAt = &now
	err := db.Delete(feedbackParameter).Error
	if err != nil {
		return err
	}

	return nil
}

// NewParameterPostgres creates new feedback parameter repository
func NewFeedbackParameterPostgres() repository.FeedbackParameterRepository {
	return &feedbackParameterPostgres{
		db: config.DB(),
	}
}
