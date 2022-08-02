package container

import (
	feedback "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/feedback"
	feedbackParameter "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/feedback_parameter"
	report "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/report"
	reportCategory "github.com/medicplus-inc/medicplus-feedback/internal/domain/service/report_category"
)

type DomainServiceIoC struct {
	feedback          feedback.FeedbackServiceInterface
	feedbackParameter feedbackParameter.FeedbackParameterServiceInterface
	report            report.ReportServiceInterface
	reportCategory    reportCategory.ReportCategoryServiceInterface
}

func NewDomainServiceIoC(ioc RepositoryIoC) DomainServiceIoC {
	return DomainServiceIoC{
		feedback:          feedback.NewFeedbackService(ioc.Feedback()),
		feedbackParameter: feedbackParameter.NewFeedbackParameterService(ioc.FeedbackParameter()),
		report:            report.NewReportService(ioc.Report()),
		reportCategory:    reportCategory.NewReportCategoryService(ioc.ReportCategory()),
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

func (ioc DomainServiceIoC) ReportCategory() reportCategory.ReportCategoryServiceInterface {
	return ioc.reportCategory
}
