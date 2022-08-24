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
	jwtAuth := jwtauth.New("HS256", []byte(config.GetValue(config.JWT_SECRET)), nil)

	// For general use
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(jwtAuth))
		r.Use(jwtauth.Authenticator)

		//Feedback Parameter
		r.Get("/feedback-parameters", listFeedbackParameters(opts))

		//Report
		r.Get("/report", getReport(opts))

		//Report Parameter
		r.Get("/report-parameters", listReportParameters(opts))

	})

	// Accessed by admin only
	r.Group(func(r chi.Router) {
		r.Use(middleware.AdminVerifier(jwtAuth))
		r.Use(jwtauth.Authenticator)

		//Feedback Parameter
		r.Post("/feedback-parameter", createFeedbackParameterForAdmin(opts))
		r.Get("/feedback-parameter", getFeedbackParameterForAdmin(opts))
		r.Put("/feedback-parameter", updateFeedbackParameterForAdmin(opts))
		r.Delete("/feedback-parameter", deleteFeedbackParameterForAdmin(opts))

		//Feedback
		r.Get("/feedbacks/admin", listFeedbacksForAdmin(opts))

		//Report Parameter
		r.Post("/report-parameter", createReportParameterForAdmin(opts))
		r.Get("/report-parameter", getReportParameterForAdmin(opts))
		r.Put("/report-parameter", updateReportParameterForAdmin(opts))
		r.Delete("/report-parameter", deleteReportParameterForAdmin(opts))

		//Report
		r.Get("/reports/admin", listReportsForAdmin(opts))

	})

	// Accessed by doctor only
	r.Group(func(r chi.Router) {
		r.Use(middleware.DoctorVerifier(jwtAuth))
		r.Use(jwtauth.Authenticator)

		//Feedback
		r.Get("/feedbacks/doctor", listFeedbacksForDoctorAnonymously(opts))

		//Report
		r.Get("/reports/doctor", listReportsForDoctorAnonymously(opts))

	})

	// Accessed by patient only
	r.Group(func(r chi.Router) {
		r.Use(middleware.PatientVerifier(jwtAuth))
		r.Use(jwtauth.Authenticator)

		//Feedback
		r.Post("/feedback-doctor", createFeedbackForPatientToDoctor(opts))
		r.Get("/feedbacks/patient", listFeedbacksForPatient(opts))

		//Report
		r.Post("/report-doctor", createReportForPatientToDoctor(opts))
		r.Get("/reports/patient", listReportsForPatient(opts))

	})

	// Accessed by patient and doctor only
	r.Group(func(r chi.Router) {
		r.Use(middleware.PatientOrDoctorVerifier(jwtAuth))
		r.Use(jwtauth.Authenticator)

		//Feedback
		r.Post("/feedback-medicplus", createFeedbackForPatientDoctorToMedicplus(opts))
		r.Post("/feedback-merchant", createFeedbackForPatientDoctorToMerchant(opts))
		r.Get("/feedback", getFeedbackForPatientDoctor(opts))
		r.Put("/feedback", updateFeedbackForPatientDoctor(opts))

		//Report
		r.Post("/report-medicplus", createReportForPatientDoctorToMedicplus(opts))
		r.Post("/report-merchant", createReportForPatientDoctorToMerchant(opts))
		r.Put("/report", updateReportForPatientDoctor(opts))

	})

	return r
}

//Feedback =========================================================================================================================================================================================
func createFeedbackForPatientToDoctor(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.CreateFeedbackForPatientToDoctor(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.CreateFeedbackRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func createFeedbackForPatientDoctorToMedicplus(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.CreateFeedbackForPatientDoctorToMedicplus(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.CreateFeedbackRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func createFeedbackForPatientDoctorToMerchant(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.CreateFeedbackForPatientDoctorToMerchant(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.CreateFeedbackRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func getFeedbackForPatientDoctor(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.GetFeedbackForPatientDoctor(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.GetFeedbackRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func listFeedbacksForDoctorAnonymously(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.ListFeedbacksForDoctorAnonymously(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.ListFeedbackRequest{},
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

func listFeedbacksForAdmin(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.ListFeedbacksForAdmin(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.ListFeedbackRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func updateFeedbackForPatientDoctor(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.UpdateFeedbackForPatientDoctor(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.UpdateFeedbackRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

//Feedback Parameter =========================================================================================================================================================================================
func createFeedbackParameterForAdmin(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.CreateFeedbackParameterForAdmin(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.CreateFeedbackParameterRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func getFeedbackParameterForAdmin(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.GetFeedbackParameterForAdmin(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.GetFeedbackParameterRequest{},
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

func updateFeedbackParameterForAdmin(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.UpdateFeedbackParameterForAdmin(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.UpdateFeedbackParameterRequest{},
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

//Report =========================================================================================================================================================================================
func createReportForPatientToDoctor(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.CreateReportForPatientToDoctor(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.CreateReportRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func createReportForPatientDoctorToMedicplus(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.CreateReportForPatientDoctorToMedicplus(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.CreateReportRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func createReportForPatientDoctorToMerchant(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.CreateReportForPatientDoctorToMerchant(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.CreateReportRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func getReport(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.GetReport(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.GetReportRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func listReportsForAdmin(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.ListReportsForAdmin(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.ListReportRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func listReportsForPatient(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.ListReportsForPatient(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.ListReportRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func listReportsForDoctorAnonymously(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.ListReportsForDoctorAnonymously(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.ListReportRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func updateReportForPatientDoctor(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.UpdateReportForPatientDoctor(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.UpdateReportRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

//Report Parameter =========================================================================================================================================================================================
func createReportParameterForAdmin(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.CreateReportParameterForAdmin(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.CreateReportParameterRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func getReportParameterForAdmin(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.GetReportParameterForAdmin(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.GetReportParameterRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func listReportParameters(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.ListReportParameters(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.ListReportParameterRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func updateReportParameterForAdmin(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.UpdateReportParameterForAdmin(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.UpdateReportParameterRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func deleteReportParameterForAdmin(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.DeleteReportParameterForAdmin(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.DeleteReportParameterRequest{},
		}, opts).ServeHTTP(w, r)
	}
}
