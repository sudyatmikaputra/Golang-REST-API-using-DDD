package report_parameter

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

// UpdateReportParameter updates report parameter data
func (s *ReportParameterService) UpdateReportParameter(ctx context.Context, params *public.UpdateReportParameterRequest) (*public.ReportParameterResponse, error) {
	userLoggedIn, _ := global.GetClaimsFromContext(ctx)

	updatedReportDomain := &domain.ReportParameter{}
	updatedReportRepo, err := s.repository.FindReportParameterByID(ctx, params.ID)
	if err != nil {
		return nil, err
	}
	if updatedReportRepo == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	updatedReportDomain.FromRepositoryModel(updatedReportRepo)
	if params.ReportType != "" {
		updatedReportDomain.ReportType = internal.ParameterType(params.ReportType)
	}
	if params.Name != "" {
		updatedReportDomain.Name = params.Name
	}
	if params.LanguageCode != "" {
		updatedReportDomain.LanguageCode = internal.LanguageCode(params.LanguageCode)
	}
	updatedReportDomain.IsDefault = params.IsDefault

	userLoggedInID := userLoggedIn["uuid"].(uuid.UUID)
	updatedReportRepo.UpdatedBy = userLoggedInID

	updatedReportRepo, err = s.repository.UpdateReportParameter(ctx, updatedReportRepo)
	if err != nil {
		return nil, err
	}
	updatedReportDomain.FromRepositoryModel(updatedReportRepo)

	return updatedReportDomain.ToPublicModel(), nil
}
