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

// categoryPostgres implements the report category parameter repository service interface
type reportCategoryPostgres struct {
	db *gorm.DB
}

// FindAllCategories queries all report category parameters
func (s *reportCategoryPostgres) FindAllReportCategories(ctx context.Context, params *public.ListReportCategoryRequest) ([]repository.ReportCategory, error) {
	db := s.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	var reportCategories []repository.ReportCategory
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
		Find(&reportCategories).Error; err != nil {
		return nil, err
	}

	return reportCategories, nil
}

// FindByCategoryID finds report category by its id
func (s *reportCategoryPostgres) FindReportCategoryByID(ctx context.Context, reportCategoryID uuid.UUID) (*repository.ReportCategory, error) {
	db := s.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	reportCategory := repository.ReportCategory{}
	err := db.First(&reportCategory, ` "deleted_at" IS NULL AND "id" = ? `, reportCategoryID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &reportCategory, nil
}

func (s *reportCategoryPostgres) FindReportCategoryByReportType(ctx context.Context, reportType internal.ParameterType, languageCode string) (*repository.ReportCategory, error) {
	db := s.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	if languageCode == "" {
		languageCode = string(internal.BahasaIndonesia)
	}

	reportCategory := repository.ReportCategory{}
	err := db.First(&reportCategory, `"deleted_at" IS NULL AND "report_type" = ? AND "language_code" = ? `, reportType, languageCode).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &reportCategory, nil
}

// Insert inserts report category
func (d *reportCategoryPostgres) InsertReportCategory(ctx context.Context, reportCategory *repository.ReportCategory) (*repository.ReportCategory, error) {
	db := d.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	reportCategory.ID, _ = uuid.NewRandom()

	// now := time.Now().UTC()
	// reportCategory.CreatedAt = now
	// reportCategory.UpdatedAt = now

	err := db.Create(reportCategory).Error
	if err != nil {
		return nil, err
	}

	return reportCategory, nil
}

// UpdateCategory updates report category
func (d *reportCategoryPostgres) UpdateReportCategory(ctx context.Context, reportCategory *repository.ReportCategory) (*repository.ReportCategory, error) {
	db := d.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	// now := time.Now().UTC()
	// reportCategory.UpdatedAt = now
	err := db.Save(reportCategory).Error
	if err != nil {
		return nil, err
	}

	return reportCategory, nil
}

// DeleteCategory deletes a report category based on its id
func (s *reportCategoryPostgres) DeleteReportCategory(ctx context.Context, reportCategory *repository.ReportCategory) error {
	db := s.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	now := time.Now().UTC()
	reportCategory.DeletedAt = &now
	err := db.Delete(reportCategory).Error
	if err != nil {
		return err
	}

	return nil
}

// NewCategoryPostgres creates new report category repository
func NewReportCategoryPostgres() repository.ReportCategoryRepository {
	return &reportCategoryPostgres{
		db: config.DB(),
	}
}
