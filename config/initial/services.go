package initial

import (
	us "rest/domain/user/services"
)

var services *Services

type Services struct {
	Repositories *Repositories
	UserService  us.UserService
}

func LoadServices(r *Repositories) {
	services = &Services{
		Repositories: r,
		UserService: *us.NewUserService(us.ServiceConfig{
			UserRepo: r.UserRepo,
		}),
	}
}

func GetServices() *Services {
	return services
}
