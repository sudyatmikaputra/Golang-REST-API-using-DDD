package container

import "sync"

var ioc IoC
var iocSingleton sync.Once

type IoC struct {
	Application   ApplicationServiceIoC
	DomainService DomainServiceIoC
	Repository    RepositoryIoC
	Validation    ValidationIoC
}

func (ioc IoC) IsEmpty() bool {
	return (IoC{}) == ioc
}

func NewIOC() IoC {
	iocSingleton.Do(func() {
		repository := NewRepositoryIoC()
		domainService := NewDomainServiceIoC(repository)
		application := NewApplicationServiceIoC(domainService, repository)
		validation := NewValidationService()

		ioc = IoC{
			Application:   application,
			DomainService: domainService,
			Repository:    repository,
			Validation:    validation,
		}
	})

	return ioc
}

func Injector() IoC {
	if ioc.IsEmpty() {
		ioc = NewIOC()
	}

	return ioc
}
