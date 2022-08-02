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

// reportParameterPostgres implements the report parameter parameter repository service interface
type reportParameterPostgres struct {
	db *gorm.DB
}

// FindAllParameters queries all report parameter parameters
func (s *reportParameterPostgres) FindAllReportParameters(ctx context.Context, params *public.ListReportParameterRequest) ([]repository.ReportParameter, error) {
	db := s.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	var reportParameters []repository.ReportParameter
	args := []interface{}{}
	where := `"deleted_at" IS NULL`
	if params.Search != "" {
		where += ` AND "name" ILIKE ?`
		args = append(args, "%"+params.Search+"%")
	}
	if params.ReportType != "" {
		where += ` AND "report_type" = ? `
		args = append(args, params.ReportType)
	}
	if params.LanguageCode != "" {
		where += ` AND "language_code" = ? `
		args = append(args, params.LanguageCode)
	}
	if params.IsDefault != nil {
		where += ` AND "is_default" = ? `
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
		Find(&reportParameters).Error; err != nil {
		return nil, err
	}

	return reportParameters, nil
}

// FindByParameterID finds report parameter by its id
func (s *reportParameterPostgres) FindReportParameterByID(ctx context.Context, reportParameterID uuid.UUID) (*repository.ReportParameter, error) {
	db := s.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	reportParameter := repository.ReportParameter{}
	err := db.First(&reportParameter, ` "deleted_at" IS NULL AND "id" = ? `, reportParameterID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &reportParameter, nil
}

func (s *reportParameterPostgres) FindReportParameterByReportType(ctx context.Context, reportType internal.ParameterType, languageCode string) (*repository.ReportParameter, error) {
	db := s.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	if languageCode == "" {
		languageCode = string(internal.BahasaIndonesia)
	}

	reportParameter := repository.ReportParameter{}
	err := db.First(&reportParameter, `"deleted_at" IS NULL AND "report_type" = ? AND "language_code" = ? `, reportType, languageCode).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &reportParameter, nil
}

// Insert inserts report parameter
func (d *reportParameterPostgres) InsertReportParameter(ctx context.Context, reportParameter *repository.ReportParameter) (*repository.ReportParameter, error) {
	db := d.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	reportParameter.ID, _ = uuid.NewRandom()

	err := db.Create(reportParameter).Error
	if err != nil {
		return nil, err
	}

	return reportParameter, nil
}

// UpdateParameter updates report parameter
func (d *reportParameterPostgres) UpdateReportParameter(ctx context.Context, reportParameter *repository.ReportParameter) (*repository.ReportParameter, error) {
	db := d.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	err := db.Save(reportParameter).Error
	if err != nil {
		return nil, err
	}

	return reportParameter, nil
}

// DeleteParameter deletes a report parameter based on its id
func (s *reportParameterPostgres) DeleteReportParameter(ctx context.Context, reportParameter *repository.ReportParameter) error {
	db := s.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	now := time.Now().UTC()
	reportParameter.DeletedAt = &now
	err := db.Delete(reportParameter).Error
	if err != nil {
		return err
	}

	return nil
}

// NewParameterPostgres creates new report parameter repository
func NewReportParameterPostgres() repository.ReportParameterRepository {
	return &reportParameterPostgres{
		db: config.DB(),
	}
}
