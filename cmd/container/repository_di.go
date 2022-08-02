package container

import (
	feedbackReportRepo "github.com/medicplus-inc/medicplus-feedback/internal/infrastructure/repository"
	"github.com/medicplus-inc/medicplus-feedback/internal/infrastructure/repository/postgres"
)

type RepositoryIoC struct {
	feedback          feedbackReportRepo.FeedbackRepository
	feedbackParameter feedbackReportRepo.FeedbackParameterRepository
	report            feedbackReportRepo.ReportRepository
	reportCategory    feedbackReportRepo.ReportCategoryRepository
}

func NewRepositoryIoC() RepositoryIoC {
	feedback := postgres.NewFeedbackPostgres()
	feedbackParameter := postgres.NewFeedbackParameterPostgres()
	report := postgres.NewReportPostgres()
	reportCategory := postgres.NewReportCategoryPostgres()
	return RepositoryIoC{
		feedback:          feedback,
		feedbackParameter: feedbackParameter,
		report:            report,
		reportCategory:    reportCategory,
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

func (ioc RepositoryIoC) ReportCategory() feedbackReportRepo.ReportCategoryRepository {
	return ioc.reportCategory
}
