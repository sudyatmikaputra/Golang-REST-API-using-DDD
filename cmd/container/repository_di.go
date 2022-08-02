package container

import (
	feedbackReportRepo "github.com/medicplus-inc/medicplus-feedback/internal/infrastructure/repository"
	"github.com/medicplus-inc/medicplus-feedback/internal/infrastructure/repository/postgres"
)

type RepositoryIoC struct {
	feedback          feedbackReportRepo.FeedbackRepository
	feedbackParameter feedbackReportRepo.FeedbackParameterRepository
	report            feedbackReportRepo.ReportRepository
	reportParameter   feedbackReportRepo.ReportParameterRepository
}

func NewRepositoryIoC() RepositoryIoC {
	feedback := postgres.NewFeedbackPostgres()
	feedbackParameter := postgres.NewFeedbackParameterPostgres()
	report := postgres.NewReportPostgres()
	reportParameter := postgres.NewReportParameterPostgres()
	return RepositoryIoC{
		feedback:          feedback,
		feedbackParameter: feedbackParameter,
		report:            report,
		reportParameter:   reportParameter,
	}
}

func (ioc RepositoryIoC) Feedback() feedbackReportRepo.FeedbackRepository {
	return ioc.feedback
}

func (ioc RepositoryIoC) FeedbackParameter() feedbackReportRepo.FeedbackParameterRepository {
	return ioc.feedbackParameter
}

func (ioc RepositoryIoC) Report() feedbackReportRepo.ReportRepository {
	return ioc.report
}

func (ioc RepositoryIoC) ReportParameter() feedbackReportRepo.ReportParameterRepository {
	return ioc.reportParameter
}
