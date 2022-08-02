package transport

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/medicplus-inc/medicplus-feedback/cmd/container"
	"github.com/medicplus-inc/medicplus-feedback/config"
	"github.com/medicplus-inc/medicplus-feedback/internal/endpoint"
	"github.com/medicplus-inc/medicplus-feedback/internal/public"
	"github.com/medicplus-inc/medicplus-feedback/lib/net/http/middleware"
	medicplusKitHttp "github.com/medicplus-inc/medicplus-kit/net/http"
)

func CompileRoute(
	r *chi.Mux,
	opts []kitHttp.ServerOption,
) http.Handler {
	jwtAuth := jwtauth.New("HS256", []byte(config.GetEnv(config.JWT_SECRET)), nil)

	// For general use
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(jwtAuth))
		r.Use(jwtauth.Authenticator)

		//Feedback Parameter
		r.Get("/feedback-parameter", listFeedbackParameters(opts))

		//Report
		r.Get("/report/{id}", getReport(opts))
		r.Get("/report", listReports(opts))

		//Report Category
		r.Get("/report-category", listReportCategories(opts))

	})

	// Accessed by admin only
	r.Group(func(r chi.Router) {
		r.Use(middleware.AdminVerifier(jwtAuth))
		r.Use(jwtauth.Authenticator)

		//Feedback Parameter
		r.Post("/feedback-parameter", createFeedbackParameterForAdmin(opts))
		r.Delete("/feedback-parameter/{id}", deleteFeedbackParameterForAdmin(opts))
		r.Put("/feedback-parameter", updateFeedbackParameterForAdmin(opts))

		//Report Category
		r.Post("/report-category", createReportCategoryForAdmin(opts))
		r.Delete("/report-category/{id}", deleteReportCategoryForAdmin(opts))
		r.Put("/report-category", updateReportCategoryForAdmin(opts))

	})

	// Accessed by doctor only
	r.Group(func(r chi.Router) {
		r.Use(middleware.DoctorVerifier(jwtAuth))
		r.Use(jwtauth.Authenticator)

		//Feedback
		r.Get("/feedback/doctor", listFeedbacksAnonymouslyForDoctor(opts))

	})

	// Accessed by patient only
	r.Group(func(r chi.Router) {
		r.Use(middleware.PatientVerifier(jwtAuth))
		r.Use(jwtauth.Authenticator)

		//Feedback
		r.Post("/feedback", createFeedbackForPatient(opts))
		r.Get("/feedback/{id}", getFeedbackForPatient(opts))
		r.Get("/feedback", listFeedbacksForPatient(opts))

	})

	// Accessed by patient and doctor only
	r.Group(func(r chi.Router) {
		r.Use(middleware.PatientOrDoctorVerifier(jwtAuth))
		r.Use(jwtauth.Authenticator)

		//Report
		r.Post("/report", createReportForPatientAndDoctor(opts))

	})

	return r
}

//Feedback
func listFeedbacksAnonymouslyForDoctor(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.ListFeedbacksAnonymouslyForDoctor(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.ListFeedbackRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func createFeedbackForPatient(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.CreateFeedbackForPatient(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.CreateFeedbackRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func getFeedbackForPatient(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.GetFeedbackForPatient(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.GetFeedbackRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func listFeedbacksForPatient(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.ListFeedbacksForPatient(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.ListFeedbackRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

//Feedback Parameter
func createFeedbackParameterForAdmin(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.CreateFeedbackParameterForAdmin(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.CreateFeedbackParameterRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func listFeedbackParameters(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.ListFeedbackParameters(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.ListFeedbackParameterRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func deleteFeedbackParameterForAdmin(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.DeleteFeedbackParameterForAdmin(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.DeleteFeedbackParameterRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func updateFeedbackParameterForAdmin(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.UpdateFeedbackParameterForAdmin(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.UpdateFeedbackParameterRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

//Report
func getReport(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.GetReport(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.GetReportRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func listReports(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.ListReports(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.ListReportRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func createReportForPatientAndDoctor(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.CreateReportForPatientAndDoctor(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.CreateReportRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

//Report Category
func listReportCategories(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.ListReportCategories(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.ListReportCategoryRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func createReportCategoryForAdmin(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.CreateReportCategoryForAdmin(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.CreateReportCategoryRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func deleteReportCategoryForAdmin(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.DeleteReportCategoryForAdmin(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.DeleteReportCategoryRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func updateReportCategoryForAdmin(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.UpdateReportCategoryForAdmin(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.UpdateReportCategoryRequest{},
		}, opts).ServeHTTP(w, r)
	}
}
