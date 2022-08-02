package command

import (
	"context"
	"net/http"

	"github.com/medicplus-inc/medicplus-feedback/internal"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	libError "github.com/medicplus-inc/medicplus-kit/error"
)

func (r CreateReportCommand) ExecuteToDoctor(ctx context.Context, params public.CreateReportRequest) (*public.ReportResponse, error) {
	// category, err := r.categoryService.GetReportCategory(ctx, params.ReportCategory.ID)
	// if err != nil {
	// 	return nil, err
	// }
	// if category == nil {
	// 	return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	// }

	report, err := r.reportService.CreateReport(ctx, &public.CreateReportRequest{
		ReportTo:         string(internal.ToDoctor),
		ReportCategoryID: params.ReportCategoryID,
		ReportToID:       params.ReportToID,
		ReportFromID:     params.ReportFromID,
		Context:          string(internal.Consultation),
		ContextID:        params.ContextID,
		Notes:            params.Notes,
	})
	if err != nil {
		return nil, err
	}
	if report == nil {
		return nil, libError.New(internal.ErrInvalidResponse, http.StatusBadRequest, internal.ErrInvalidResponse.Error())
	}

	// report.ReportCategory = *category

	return report, nil
}
