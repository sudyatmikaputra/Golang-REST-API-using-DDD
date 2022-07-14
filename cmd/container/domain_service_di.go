package container

type DomainServiceIoC struct {
}

func NewDomainServiceIoC(ioc RepositoryIoC) DomainServiceIoC {
	return DomainServiceIoC{}
}
