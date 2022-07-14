package application

type Commands struct {
}

type Queries struct {
}

type Application struct {
	Commands Commands
	Queries  Queries
}

func New() Application {
	return Application{
		Commands: Commands{},
		Queries:  Queries{},
	}
}
