package services

import (
	"example-service/internal/config"
	"example-service/internal/storage/repositories"
	//"github.com/redis/go-redis/v9"
	//"gopkg.in/gomail.v2"
)

type Services struct {
	// ...
}

func Init(repos *repositories.Repos, cfg *config.Config) *Services {
	return &Services{
		// ...
	}
}
