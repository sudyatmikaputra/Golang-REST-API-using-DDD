package endpoint

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/medicplus-inc/medicplus-feedback/cmd/container"
	"github.com/medicplus-inc/medicplus-feedback/config"
	"github.com/medicplus-inc/medicplus-feedback/internal"
	feedbackReportApp "github.com/medicplus-inc/medicplus-feedback/internal/application"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	"github.com/medicplus-inc/medicplus-kit/database"
	libError "github.com/medicplus-inc/medicplus-kit/error"
	httpResponse "github.com/medicplus-inc/medicplus-kit/net/http"
)

func DeleteReportParameterForAdmin(application feedbackReportApp.Application) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (res interface{}, err error) {
		payload := req.(*public.DeleteReportParameterRequest)

		err = database.RunInTransaction(ctx, config.DB(), func(ctx context.Context) (err error) {
			if payload == nil {
				return libError.New(err, http.StatusBadRequest, internal.ErrInvalidRequest.Error())
			}
			if err = container.Injector().Validation.Validator.Struct(payload); err != nil {
				return libError.New(err, http.StatusBadRequest, internal.ErrInvalidRequest.Error())
			}
			err = application.Commands.DeleteReportParameterForAdmin.Execute(ctx, *payload)

			return err
		})

		return httpResponse.ResponseWithRequestTime(ctx, nil, nil), err
	}
}
