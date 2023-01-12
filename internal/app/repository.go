package app

type Repositories struct {
}

func (application *Application) WithRepositories() {
	application.Repositories = new(Repositories)
}
