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
		r.Get("/feedback-parameters", listFeedbackParameters(opts)) //A

		//Report
		r.Get("/report/{id}", getReport(opts)) //A

		//Report Parameter
		r.Get("/report-parameters", listReportParameters(opts)) //A

	})

	// Accessed by admin only
	r.Group(func(r chi.Router) {
		r.Use(middleware.AdminVerifier(jwtAuth))
		r.Use(jwtauth.Authenticator)

		//Feedback Parameter
		r.Post("/feedback-parameter", createFeedbackParameterForAdmin(opts))        //A
		r.Get("/feedback-parameter/{id}", getFeedbackParameterForAdmin(opts))       //A
		r.Put("/feedback-parameter", updateFeedbackParameterForAdmin(opts))         //A
		r.Delete("/feedback-parameter/{id}", deleteFeedbackParameterForAdmin(opts)) //A

		//Feedback
		r.Get("/feedbacks/admin", listFeedbacksForAdmin(opts)) //A

		//Report Parameter
		r.Post("/report-parameter", createReportParameterForAdmin(opts))        //A
		r.Get("/report-parameter/{id}", getReportParameterForAdmin(opts))       //A
		r.Put("/report-parameter", updateReportParameterForAdmin(opts))         //A
		r.Delete("/report-parameter/{id}", deleteReportParameterForAdmin(opts)) //A

		//Report
		r.Get("/reports/admin", listReportsForAdmin(opts)) //A

	})

	// Accessed by doctor only
	r.Group(func(r chi.Router) {
		r.Use(middleware.DoctorVerifier(jwtAuth))
		r.Use(jwtauth.Authenticator)

		//Feedback
		r.Get("/feedbacks/doctor", listFeedbacksForDoctorAnonymously(opts)) //A

		//Report
		r.Get("/reports/doctor", listReportsForDoctorAnonymously(opts)) //A

	})

	// Accessed by patient only
	r.Group(func(r chi.Router) {
		r.Use(middleware.PatientVerifier(jwtAuth))
		r.Use(jwtauth.Authenticator)

		//Feedback
		r.Post("/feedback-doctor", createFeedbackForPatientToDoctor(opts)) //A
		r.Get("/feedbacks/patient", listFeedbacksForPatient(opts))         //A

		//Report
		r.Post("/report-doctor", createReportForPatientToDoctor(opts)) //A
		r.Get("/reports/patient", listReportsForPatient(opts))         //A

	})

	// Accessed by patient and doctor only
	r.Group(func(r chi.Router) {
		r.Use(middleware.PatientOrDoctorVerifier(jwtAuth))
		r.Use(jwtauth.Authenticator)

		//Feedback
		r.Post("/feedback-medicplus", createFeedbackForPatientDoctorToMedicplus(opts)) //A
		r.Post("/feedback-merchant", createFeedbackForPatientDoctorToMerchant(opts))   //A
		r.Get("/feedback/{id}", getFeedbackForPatientDoctor(opts))                     //A
		r.Put("/feedback", updateFeedbackForPatientDoctor(opts))                       //A

		//Report
		r.Post("/report-medicplus", createReportForPatientDoctorToMedicplus(opts)) //A
		r.Post("/report-merchant", createReportForPatientDoctorToMerchant(opts))   //A
		r.Put("/report", updateReportForPatientDoctor(opts))                       //A

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
} //D

func createFeedbackForPatientDoctorToMedicplus(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.CreateFeedbackForPatientDoctorToMedicplus(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.CreateFeedbackRequest{},
		}, opts).ServeHTTP(w, r)
	}
} //D

func createFeedbackForPatientDoctorToMerchant(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.CreateFeedbackForPatientDoctorToMerchant(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.CreateFeedbackRequest{},
		}, opts).ServeHTTP(w, r)
	}
} //D

func getFeedbackForPatientDoctor(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.GetFeedbackForPatientDoctor(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.GetFeedbackRequest{},
		}, opts).ServeHTTP(w, r)
	}
} //D

func listFeedbacksForDoctorAnonymously(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.ListFeedbacksForDoctorAnonymously(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.ListFeedbackRequest{},
		}, opts).ServeHTTP(w, r)
	}
} //D

func listFeedbacksForPatient(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.ListFeedbacksForPatient(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.ListFeedbackRequest{},
		}, opts).ServeHTTP(w, r)
	}
} //D

func listFeedbacksForAdmin(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.ListFeedbacksForAdmin(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.ListFeedbackRequest{},
		}, opts).ServeHTTP(w, r)
	}
} //D

func updateFeedbackForPatientDoctor(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.UpdateFeedbackForPatientDoctor(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.UpdateFeedbackRequest{},
		}, opts).ServeHTTP(w, r)
	}
} //D

//Feedback Parameter =========================================================================================================================================================================================
func createFeedbackParameterForAdmin(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.CreateFeedbackParameterForAdmin(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.CreateFeedbackParameterRequest{},
		}, opts).ServeHTTP(w, r)
	}
} //D

func getFeedbackParameterForAdmin(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.GetFeedbackParameterForAdmin(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.GetFeedbackParameterRequest{},
		}, opts).ServeHTTP(w, r)
	}
} //D

func listFeedbackParameters(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.ListFeedbackParameters(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.ListFeedbackParameterRequest{},
		}, opts).ServeHTTP(w, r)
	}
} //D

func updateFeedbackParameterForAdmin(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.UpdateFeedbackParameterForAdmin(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.UpdateFeedbackParameterRequest{},
		}, opts).ServeHTTP(w, r)
	}
} //D

func deleteFeedbackParameterForAdmin(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.DeleteFeedbackParameterForAdmin(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.DeleteFeedbackParameterRequest{},
		}, opts).ServeHTTP(w, r)
	}
} //D

//Report =========================================================================================================================================================================================
func createReportForPatientToDoctor(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.CreateReportForPatientToDoctor(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.CreateReportRequest{},
		}, opts).ServeHTTP(w, r)
	}
} //D

func createReportForPatientDoctorToMedicplus(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.CreateReportForPatientDoctorToMedicplus(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.CreateReportRequest{},
		}, opts).ServeHTTP(w, r)
	}
} //D

func createReportForPatientDoctorToMerchant(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.CreateReportForPatientDoctorToMerchant(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.CreateReportRequest{},
		}, opts).ServeHTTP(w, r)
	}
} //D

func getReport(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.GetReport(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.GetReportRequest{},
		}, opts).ServeHTTP(w, r)
	}
} //D

func listReportsForAdmin(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.ListReportsForAdmin(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.ListReportRequest{},
		}, opts).ServeHTTP(w, r)
	}
} //D

func listReportsForPatient(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.ListReportsForPatient(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.ListReportRequest{},
		}, opts).ServeHTTP(w, r)
	}
} //D

func listReportsForDoctorAnonymously(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.ListReportsForDoctorAnonymously(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.ListReportRequest{},
		}, opts).ServeHTTP(w, r)
	}
} //D

func updateReportForPatientDoctor(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.UpdateReportForPatientDoctor(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.UpdateReportRequest{},
		}, opts).ServeHTTP(w, r)
	}
} //D

//Report Parameter =========================================================================================================================================================================================
func createReportParameterForAdmin(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.CreateReportParameterForAdmin(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.CreateReportParameterRequest{},
		}, opts).ServeHTTP(w, r)
	}
} //D

func getReportParameterForAdmin(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.GetReportParameterForAdmin(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.GetReportParameterRequest{},
		}, opts).ServeHTTP(w, r)
	}
} //D

func listReportParameters(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.ListReportParameters(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.ListReportParameterRequest{},
		}, opts).ServeHTTP(w, r)
	}
} //D

func updateReportParameterForAdmin(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.UpdateReportParameterForAdmin(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.UpdateReportParameterRequest{},
		}, opts).ServeHTTP(w, r)
	}
} //D

func deleteReportParameterForAdmin(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medicplusKitHttp.NewHTTPServer(endpoint.DeleteReportParameterForAdmin(container.Injector().Application.FeedbackReport()), medicplusKitHttp.Option{
			DecodeModel: &public.DeleteReportParameterRequest{},
		}, opts).ServeHTTP(w, r)
	}
} //D
