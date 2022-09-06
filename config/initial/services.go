package initial

var services *Services

type Services struct {
	Repositories *Repositories
}

func LoadServices(r *Repositories) {
	services = &Services{
		Repositories: r,
	}
}

func GetServices() *Services {
	return services
}
