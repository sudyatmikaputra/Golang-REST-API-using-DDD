package application

import (
	"github.com/medicplus-inc/medicplus-feedback/internal/application/command"
	"github.com/medicplus-inc/medicplus-feedback/internal/application/query"
	feedbackDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/feedback"
	feedbackParameterDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/feedback_parameter"
	reportDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/report"
	reportParameterDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/report_parameter"
)

type Commands struct {
	CreateFeedbackParameterForAdmin           command.CreateFeedbackParameterForAdminCommand
	CreateReportParameterForAdmin             command.CreateReportParameterForAdminCommand
	DeleteFeedbackParameterForAdmin           command.DeleteFeedbackParameterForAdminCommand
	DeleteReportParameterForAdmin             command.DeleteReportParameterForAdminCommand
	UpdateFeedbackParameterForAdmin           command.UpdateFeedbackParameterForAdminCommand
	UpdateReportParameterForAdmin             command.UpdateReportParameterForAdminCommand
	CreateFeedbackForPatientToDoctor          command.CreateFeedbackForPatientToDoctorCommand
	CreateReportForPatientToDoctor            command.CreateReportForPatientToDoctorCommand
	CreateFeedbackForPatientDoctorToMedicplus command.CreateFeedbackForPatientDoctorToMedicplusCommand
	CreateFeedbackForPatientDoctorToMerchant  command.CreateFeedbackForPatientDoctorToMerchantCommand
	CreateReportForPatientDoctorToMedicplus   command.CreateReportForPatientDoctorToMedicplusCommand
	CreateReportForPatientDoctorToMerchant    command.CreateReportForPatientDoctorToMerchantCommand
	UpdateFeedbackForPatientDoctor            command.UpdateFeedbackForPatientDoctorCommand
	UpdateReportForPatientDoctor              command.UpdateReportForPatientDoctorCommand
}

type Queries struct {
	GetFeedbackParameterForAdmin      query.GetFeedbackParameterForAdminQuery
	GetReportParameterForAdmin        query.GetReportParameterForAdminQuery
	ListFeedbacksForAdmin             query.ListFeedbacksForAdminQuery
	ListReportsForAdmin               query.ListReportsForAdminQuery
	GetReport                         query.GetReportQuery
	ListFeedbackParameters            query.ListFeedbackParametersQuery
	ListReportParameters              query.ListReportParametersQuery
	ListFeedbacksForDoctorAnonymously query.ListFeedbacksForDoctorAnonymouslyQuery
	ListReportsForDoctorAnonymously   query.ListReportsForDoctorAnonymouslyQuery
	GetFeedbackForPatientDoctor       query.GetFeedbackForPatientDoctorQuery
	ListFeedbacksForPatient           query.ListFeedbacksForPatientQuery
	ListReportsForPatient             query.ListReportsForPatientQuery
}

type Application struct {
	Commands Commands
	Queries  Queries
}

func New(
	feedbackService feedbackDomainService.FeedbackServiceInterface,
	feedbackParameterService feedbackParameterDomainService.FeedbackParameterServiceInterface,
	reportService reportDomainService.ReportServiceInterface,
	reportParameterService reportParameterDomainService.ReportParameterServiceInterface,
) Application {
	return Application{
		Commands: Commands{
			CreateFeedbackParameterForAdmin:           command.NewCreateFeedbackParameterForAdminCommand(feedbackParameterService),
			CreateReportParameterForAdmin:             command.NewCreateReportParameterForAdminCommand(reportParameterService),
			DeleteFeedbackParameterForAdmin:           command.NewDeleteFeedbackParameterForAdminCommand(feedbackParameterService),
			DeleteReportParameterForAdmin:             command.NewDeleteReportParameterForAdminCommand(reportParameterService),
			UpdateFeedbackParameterForAdmin:           command.NewUpdateFeedbackParameterForAdminCommand(feedbackParameterService),
			UpdateReportParameterForAdmin:             command.NewUpdateReportParameterForAdminCommand(reportParameterService),
			CreateFeedbackForPatientToDoctor:          command.NewCreateFeedbackForPatientToDoctorCommand(feedbackService, feedbackParameterService),
			CreateReportForPatientToDoctor:            command.NewCreateReportForPatientToDoctorCommand(reportService, reportParameterService),
			CreateFeedbackForPatientDoctorToMedicplus: command.NewCreateFeedbackForPatientDoctorToMedicplusCommand(feedbackService, feedbackParameterService),
			CreateFeedbackForPatientDoctorToMerchant:  command.NewCreateFeedbackForPatientDoctorToMerchantCommand(feedbackService, feedbackParameterService),
			CreateReportForPatientDoctorToMedicplus:   command.NewCreateReportForPatientDoctorToMedicplusCommand(reportService, reportParameterService),
			CreateReportForPatientDoctorToMerchant:    command.NewCreateReportForPatientDoctorToMerchantCommand(reportService, reportParameterService),
			UpdateFeedbackForPatientDoctor:            command.NewUpdateFeedbackForPatientDoctorCommand(feedbackService),
			UpdateReportForPatientDoctor:              command.NewUpdateReportForPatientDoctorCommand(reportService),
		},
		Queries: Queries{
			GetFeedbackParameterForAdmin:      query.NewGetFeedbackParameterForAdminQuery(feedbackParameterService),
			GetReportParameterForAdmin:        query.NewGetReportParameterForAdminQuery(reportParameterService),
			ListFeedbacksForAdmin:             query.NewListFeedbacksForAdminQuery(feedbackService, feedbackParameterService),
			ListReportsForAdmin:               query.NewListReportsForAdminQuery(reportService, reportParameterService),
			GetReport:                         query.NewGetReportQuery(reportService, reportParameterService),
			ListFeedbackParameters:            query.NewListFeedbackParametersQuery(feedbackParameterService),
			ListReportParameters:              query.NewListReportParametersQuery(reportParameterService),
			ListFeedbacksForDoctorAnonymously: query.NewListFeedbacksForDoctorAnonymouslyQuery(feedbackService, feedbackParameterService),
			ListReportsForDoctorAnonymously:   query.NewListReportsForDoctorAnonymouslyQuery(reportService, reportParameterService),
			GetFeedbackForPatientDoctor:       query.NewGetFeedbackForPatientDoctorQuery(feedbackService, feedbackParameterService),
			ListFeedbacksForPatient:           query.NewListFeedbacksForPatientQuery(feedbackService, feedbackParameterService),
			ListReportsForPatient:             query.NewListReportsForPatientQuery(reportService, reportParameterService),
		},
	}
}
