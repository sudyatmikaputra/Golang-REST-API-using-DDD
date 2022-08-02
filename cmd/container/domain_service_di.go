package container

import (
	feedback "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/feedback"
	feedbackParameter "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/feedback_parameter"
	report "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/report"
	reportParameter "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/report_parameter"
)

type DomainServiceIoC struct {
	feedback          feedback.FeedbackServiceInterface
	feedbackParameter feedbackParameter.FeedbackParameterServiceInterface
	report            report.ReportServiceInterface
	reportParameter   reportParameter.ReportParameterServiceInterface
}

func NewDomainServiceIoC(ioc RepositoryIoC) DomainServiceIoC {
	return DomainServiceIoC{
		feedback:          feedback.NewFeedbackService(ioc.Feedback()),
		feedbackParameter: feedbackParameter.NewFeedbackParameterService(ioc.FeedbackParameter()),
		report:            report.NewReportService(ioc.Report()),
		reportParameter:   reportParameter.NewReportParameterService(ioc.ReportParameter()),
	}
}

func (ioc DomainServiceIoC) Feedback() feedback.FeedbackServiceInterface {
	return ioc.feedback
}

func (ioc DomainServiceIoC) FeedbackParameter() feedbackParameter.FeedbackParameterServiceInterface {
	return ioc.feedbackParameter
}

func (ioc DomainServiceIoC) Report() report.ReportServiceInterface {
	return ioc.report
}

func (ioc DomainServiceIoC) ReportParameter() reportParameter.ReportParameterServiceInterface {
	return ioc.reportParameter
}
