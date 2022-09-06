package initial

import (
	"gorm.io/gorm"
	appControllers "rest/app/controllers"
	userControllers "rest/domain/user/controllers"
)

var controllers *Controllers

type Controllers struct {
	AuthController *userControllers.AuthController
	MainController *appControllers.Controller
	UserController *userControllers.Controller
}

func LoadControllers(s *Services, _ *gorm.DB) {
	controllers = &Controllers{
		AuthController: userControllers.NewAuthController(s.Repositories.AuthRepo, s.Repositories.UserRepo),
		UserController: userControllers.NewController(s.UserService),
	}
}

func GetControllers() *Controllers {
	return controllers
}
