package initial

import (
	"gorm.io/gorm"
	uRepos "rest/domain/user/repositories"
)

var repositories *Repositories

type Repositories struct {
	UserRepo uRepos.UserRepository
	AuthRepo uRepos.AuthRepository
}

func LoadRepos(db *gorm.DB) {
	repositories = &Repositories{
		UserRepo: uRepos.NewUserRepo(db),
		AuthRepo: uRepos.NewAuthRepo(db),
	}
}

func GetRepositories() *Repositories {
	return repositories
}
