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

func GetFeedbackParameterForAdmin(application feedbackReportApp.Application) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (res interface{}, err error) {
		payload := req.(*public.GetFeedbackParameterRequest)

		err = database.RunInTransaction(ctx, config.DB(), func(ctx context.Context) (err error) {
			if payload == nil {
				return libError.New(err, http.StatusBadRequest, internal.ErrInvalidRequest.Error())
			}
			if err = container.Injector().Validation.Validator.Struct(payload); err != nil {
				return libError.New(err, http.StatusBadRequest, internal.ErrInvalidRequest.Error())
			}
			res, err = application.Queries.GetFeedbackParameterForAdmin.Execute(ctx, *payload)

			return err
		})

		return httpResponse.ResponseWithRequestTime(ctx, res, nil), err
	}
}
