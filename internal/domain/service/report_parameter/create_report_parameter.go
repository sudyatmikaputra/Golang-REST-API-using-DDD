package report_parameter

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/domain"
	"github.com/medicplus-inc/medicplus-feedback/internal/global"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

// CreateReportCategory creates a new report catgeory
func (s *ReportParameterService) CreateReportParameter(ctx context.Context, params *public.CreateReportParameterRequest) (*public.ReportParameterResponse, error) {
	userLoggedIn, _ := global.GetClaimsFromContext(ctx)
	userLoggedInID := uuid.MustParse(userLoggedIn["uuid"].(string))

	fmt.Println(internal.ParameterType(params.ReportType))
	fmt.Println(params.Name)
	fmt.Println(internal.LanguageCode(params.LanguageCode))
	fmt.Println(params.IsDefault)
	fmt.Println("before domain")
	fmt.Println(userLoggedInID)

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

	fmt.Println("after domain")

	reportParameterRepo := reportParameterDomain.ToRepositoryModel()

	insertedCategory, err := s.repository.InsertReportParameter(ctx, reportParameterRepo)
	if err != nil {
		return nil, err
	}
	if insertedCategory == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	reportParameterDomain.FromRepositoryModel(insertedCategory)

	return reportParameterDomain.ToPublicModel(), nil
}
