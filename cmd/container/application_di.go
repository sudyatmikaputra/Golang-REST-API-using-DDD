package container

type ApplicationServiceIoC struct {
}

func NewApplicationServiceIoC(dsIoc DomainServiceIoC, rIoc RepositoryIoC) ApplicationServiceIoC {
	return ApplicationServiceIoC{}
}
