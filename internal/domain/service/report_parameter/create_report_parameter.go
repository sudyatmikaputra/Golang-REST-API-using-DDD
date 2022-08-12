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

func (s *ReportParameterService) CreateReportParameter(ctx context.Context, params *public.CreateReportParameterRequest) (*public.ReportParameterResponse, error) {
	userLoggedIn, _ := global.GetClaimsFromContext(ctx)
	userLoggedInID := uuid.MustParse(userLoggedIn["uuid"].(string))

	existingReportParameter, err := s.repository.FindReportParameterByReportType(ctx, internal.ParameterType(params.ReportType), params.LanguageCode)
	if err != nil {
		return nil, err
	}
	if existingReportParameter != nil {
		return nil, libError.New(internal.ErrLanguageCodeAlreadyExists, http.StatusBadRequest, internal.ErrLanguageCodeAlreadyExists.Error())
	}

	reportParameterDomain := &domain.ReportParameter{
		ReportType:   internal.ParameterType(params.ReportType),
		Name:         params.Name,
		LanguageCode: internal.LanguageCode(params.LanguageCode),
		IsDefault:    params.IsDefault,
		CreatedBy:    userLoggedInID,
		UpdatedBy:    userLoggedInID,
	}

	reportParameterRepo := reportParameterDomain.ToRepositoryModel()

	insertedReportParameter, err := s.repository.InsertReportParameter(ctx, reportParameterRepo)
	if err != nil {
		return nil, err
	}
	if insertedReportParameter == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	reportParameterDomain.FromRepositoryModel(insertedReportParameter)

	return reportParameterDomain.ToPublicModel(), nil
}
