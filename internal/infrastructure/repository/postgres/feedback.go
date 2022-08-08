package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/config"
	"github.com/medicplus-inc/medicplus-feedback/internal/infrastructure/repository"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	"github.com/medicplus-inc/medicplus-kit/database"
	"gorm.io/gorm"
)

// feedbackPostgres implements the feedback repository service interface
type feedbackPostgres struct {
	db *gorm.DB
}

// FindAll queries all feedbacks
func (s *feedbackPostgres) FindAllFeedbacks(ctx context.Context, params *public.ListFeedbackRequest) ([]repository.Feedback, error) {
	db := s.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	var feedbacks []repository.Feedback
	args := []interface{}{}
	where := ``
	if params.Search != "" {
		where += `("feedback_value" ILIKE ? OR "notes" ILIKE ?)`
		args = append(args, "%"+params.Search+"%")
		args = append(args, "%"+params.Search+"%")
	}
	if params.FeedbackType != "" {
		if where != "" {
			where += ` AND `
		}
		where += ` "feedback_type" = ?`
		args = append(args, params.FeedbackType)
	}
	if params.FeedbackToID != uuid.Nil {
		if where != "" {
			where += ` AND `
		}
		where += ` "feedback_to_id" = ?`
		args = append(args, params.FeedbackToID)
	}
	if params.FeedbackFromID != uuid.Nil {
		if where != "" {
			where += ` AND `
		}
		where += ` "feedback_from_id" = ?`
		args = append(args, params.FeedbackFromID)
	}

	order := `"created_at" DESC`
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

	return feedbacks, nil
}

// FindByID finds feedback by its id
func (s *feedbackPostgres) FindFeedbackByID(ctx context.Context, feedbackID uuid.UUID) (*repository.Feedback, error) {
	db := s.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	feedback := repository.Feedback{}
	err := db.First(&feedback, `"id" = ? `, feedbackID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &feedback, nil
}

// Insert inserts feedback
func (d *feedbackPostgres) InsertFeedback(ctx context.Context, feedback *repository.Feedback) (*repository.Feedback, error) {
	db := d.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	feedback.ID, _ = uuid.NewRandom()

	_ = db.Create(feedback)

	return feedback, nil
}

// Update updates feedback
func (d *feedbackPostgres) UpdateFeedback(ctx context.Context, feedback *repository.Feedback) (*repository.Feedback, error) {
	db := d.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	err := db.Save(feedback).Error
	if err != nil {
		return nil, err
	}

	return feedback, nil
}

// Delete deletes an feedback based on its id
func (s *feedbackPostgres) DeleteFeedback(ctx context.Context, feedback *repository.Feedback) error {
	db := s.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	err := db.Delete(feedback).Error
	if err != nil {
		return err
	}

	return nil
}

// NewFeedbackPostgres creates new feedback repository
func NewFeedbackPostgres() repository.FeedbackRepository {
	return &feedbackPostgres{
		db: config.DB(),
	}
}
