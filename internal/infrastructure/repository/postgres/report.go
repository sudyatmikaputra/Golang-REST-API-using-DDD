package postgres

import (
	"context"
	"errors"

	"github.com/medicplus-inc/medicplus-feedback/config"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal/infrastructure/repository"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	"github.com/medicplus-inc/medicplus-kit/database"
	"gorm.io/gorm"
)

// reportPostgres implements the report repository service interface
type reportPostgres struct {
	db *gorm.DB
}

// test comment
// FindAllReports queries all reports
func (s *reportPostgres) FindAllReports(ctx context.Context, params *public.ListReportRequest) ([]repository.Report, error) {
	db := s.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	var reports []repository.Report
	args := []interface{}{}
	where := ``
	if params.Search != "" {
		where += ` AND "notes" ILIKE ?`
		args = append(args, "%"+params.Search+"%")
	}
	if params.ReportType != "" {
		if where != "" {
			where += ` AND `
		}
		where += ` "report_type" = ?`
		args = append(args, params.ReportType)
	}
	if params.ReportFromID != uuid.Nil {
		if where != "" {
			where += ` AND `
		}
		where += ` "report_from_id" = ?`
		args = append(args, params.ReportFromID)
	}
	if params.ReportToID != uuid.Nil {
		if where != "" {
			where += ` AND `
		}
		where += ` "report_to_id" = ?`
		args = append(args, params.ReportToID)
	}

	order := `"created_at" DESC`
	if err := db.Where(
		where,
		args...,
	).
		Order(order).
		Offset(((params.Page - 1) * params.Limit)).
		Limit(params.Limit).
		Find(&reports).Error; err != nil {
		return nil, err
	}

	return reports, nil
}

// FindByReportID finds report by its id
func (s *reportPostgres) FindReportByID(ctx context.Context, reportID uuid.UUID) (*repository.Report, error) {
	db := s.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	report := repository.Report{}
	err := db.First(&report, ` "deleted_at" IS NULL AND "id" = ?`, reportID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &report, nil
}

// InsertReport inserts report
func (d *reportPostgres) InsertReport(ctx context.Context, report *repository.Report) (*repository.Report, error) {
	db := d.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	report.ID, _ = uuid.NewRandom()

	err := db.Create(report).Error
	if err != nil {
		return nil, err
	}

	return report, nil
}

// UpdateReport updates report
func (d *reportPostgres) UpdateReport(ctx context.Context, report *repository.Report) (*repository.Report, error) {
	db := d.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	err := db.Save(report).Error
	if err != nil {
		return nil, err
	}

	return report, nil
}

// DeleteReport deletes an report based on its id
func (s *reportPostgres) DeleteReport(ctx context.Context, report *repository.Report) error {
	db := s.db
	tx, ok := database.QueryFromContext(ctx)
	if ok {
		db = tx
	}

	err := db.Delete(report).Error
	if err != nil {
		return err
	}

	return nil
}

// NewReportPostgres creates new report repository
func NewReportPostgres() repository.ReportRepository {
	return &reportPostgres{
		db: config.DB(),
	}
}
