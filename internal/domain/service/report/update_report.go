package report

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/domain"
	"github.com/medicplus-inc/medicplus-feedback/internal/global"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

// UpdateReport updates report data
func (s *ReportService) UpdateReport(ctx context.Context, params *public.UpdateReportRequest) (*public.ReportResponse, error) {
	userLoggedIn, _ := global.GetClaimsFromContext(ctx)
	updatedReportDomain := &domain.Report{}
	updatedReportRepo, err := s.repository.FindReportByID(ctx, params.ID)
	if err != nil {
		return nil, err
	}
	if updatedReportRepo == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	if userLoggedIn["uuid"].(uuid.UUID) != updatedReportRepo.ReportFromID {
		return nil, libError.New(internal.ErrNotAuthorized, http.StatusUnauthorized, internal.ErrNotAuthorized.Error())
	}
	updatedReportDomain.FromRepositoryModel(updatedReportRepo)
	if params.Notes != "" {
		updatedReportDomain.Notes = params.Notes
	}
	updatedReportRepo, err = s.repository.UpdateReport(ctx, updatedReportDomain.ToRepositoryModel())
	if err != nil {
		return nil, err
	}
	updatedReportDomain.FromRepositoryModel(updatedReportRepo)

	return updatedReportDomain.ToPublicModel(), nil
}
