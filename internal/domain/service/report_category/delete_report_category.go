package reportcategory

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/global"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

//DeleteReportCategory deleting report category
func (s *ReportCategoryService) DeleteReportCategory(ctx context.Context, params *public.DeleteReportCategoryRequest) error {
	userLoggedIn, _ := global.GetClaimsFromContext(ctx)
	categoryRepo, err := s.repository.FindReportCategoryByID(ctx, params.CategoryID)
	if err != nil {
		return err
	}
	if categoryRepo == nil {
		return libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	userLoggedInID := userLoggedIn["uuid"].(uuid.UUID)
	categoryRepo.DeletedBy = &userLoggedInID

	err = s.repository.DeleteReportCategory(ctx, categoryRepo)
	if err != nil {
		return err
	}

	return nil
}
