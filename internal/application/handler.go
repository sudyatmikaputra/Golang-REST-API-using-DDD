package application

import (
	"github.com/medicplus-inc/medicplus-feedback/internal/application/command"
	"github.com/medicplus-inc/medicplus-feedback/internal/application/query"
	feedbackDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/feedback"
	feedbackParameterDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/feedback_parameter"
	reportDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/report"
	reportCategoryDomainService "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/report_category"
)

type Commands struct {
	CreateFeedbackParameter command.CreateFeedbackParameterCommand
	CreateFeedback          command.CreateFeedbackCommand
	CreateReportCategory    command.CreateReportCategoryCommand
	CreateReport            command.CreateReportCommand
	DeleteFeedbackParameter command.DeleteFeedbackParameterCommand
	DeleteReportCategory    command.DeleteReportCategoryCommand
	UpdateFeedbackParameter command.UpdateFeedbackParameterCommand
	UpdateReportCategory    command.UpdateReportCategoryCommand
}

type Queries struct {
	GetFeedback            query.GetFeedbackQuery
	GetReport              query.GetReportQuery
	ListFeedbackParameters query.ListFeedbackParametersQuery
	ListFeedbacks          query.ListFeedbacksQuery
	ListReportCategories   query.ListReportCategoriesQuery
	ListReports            query.ListReportsQuery
}

type Application struct {
	Commands Commands
	Queries  Queries
}

func New(
	feedbackService feedbackDomainService.FeedbackServiceInterface,
	parameterService feedbackParameterDomainService.FeedbackParameterServiceInterface,
	reportService reportDomainService.ReportServiceInterface,
	categoryService reportCategoryDomainService.ReportCategoryServiceInterface,
) Application {
	return Application{
		Commands: Commands{
			CreateFeedbackParameter: command.NewCreateFeedbackParameterCommand(parameterService),
			CreateFeedback:          command.NewCreateFeedbackCommand(feedbackService, parameterService),
			CreateReportCategory:    command.NewCreateReportCategoryCommand(categoryService),
			CreateReport:            command.NewCreateReportCommand(reportService, categoryService),
			DeleteFeedbackParameter: command.NewDeleteFeedbackParameterCommand(parameterService),
			DeleteReportCategory:    command.NewDeleteReportCategoryCommand(categoryService),
			UpdateFeedbackParameter: command.NewUpdateFeedbackParameterCommand(parameterService),
			UpdateReportCategory:    command.NewUpdateReportCategoryCommand(categoryService),
		},
		Queries: Queries{
			GetFeedback:            query.NewGetFeedbackQuery(feedbackService),
			GetReport:              query.NewGetReportQuery(reportService, categoryService),
			ListFeedbackParameters: query.NewListFeedbackParametersQuery(parameterService),
			ListFeedbacks:          query.NewListFeedbacksQuery(feedbackService, parameterService),
			ListReportCategories:   query.NewListReportCategoriesQuery(categoryService),
			ListReports:            query.NewListReportsQuery(reportService, categoryService),
		},
	}
}
