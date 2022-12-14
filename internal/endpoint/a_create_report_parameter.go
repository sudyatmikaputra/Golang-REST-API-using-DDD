package endpoint

import (
	"context"
	"fmt"
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

func CreateReportParameterForAdmin(application feedbackReportApp.Application) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (res interface{}, err error) {
		payload := req.(*public.CreateReportParameterRequest)

		err = database.RunInTransaction(ctx, config.DB(), func(ctx context.Context) (err error) {
			if payload == nil {
				return libError.New(err, http.StatusBadRequest, internal.ErrInvalidRequest.Error())
			}
			fmt.Println("TEST")
			if err = container.Injector().Validation.Validator.Struct(payload); err != nil {
				fmt.Println("after TEST")
				return libError.New(err, http.StatusBadRequest, internal.ErrInvalidRequest.Error())
			}

			res, err = application.Commands.CreateReportParameterForAdmin.Execute(ctx, *payload)

			return err
		})

		return httpResponse.ResponseWithRequestTime(ctx, res, nil), err
	}
}
