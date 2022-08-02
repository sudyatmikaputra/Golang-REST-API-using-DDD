package container

import (
	feedbackReportApp "github.com/medicplus-inc/medicplus-feedback/internal/application"
)

type ApplicationServiceIoC struct {
	feedbackReport feedbackReportApp.Application
}

func NewApplicationServiceIoC(dsIoc DomainServiceIoC, rIoc RepositoryIoC) ApplicationServiceIoC {
	return ApplicationServiceIoC{
		feedbackReport: feedbackReportApp.New(
			dsIoc.Feedback(),
			dsIoc.FeedbackParameter(),
			dsIoc.Report(),
			dsIoc.ReportParameter(),
		),
	}
}

func (ioc ApplicationServiceIoC) FeedbackReport() feedbackReportApp.Application {
	return ioc.feedbackReport
}
