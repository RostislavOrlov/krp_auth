package main

import (
	"github.com/IlyaZayats/auth/internal/db"
	"github.com/IlyaZayats/auth/internal/handlers/user"
	"github.com/IlyaZayats/auth/internal/repositories"
	"github.com/IlyaZayats/auth/internal/server"
	"github.com/IlyaZayats/auth/internal/services"
	"github.com/sirupsen/logrus"
	"log"
)

func main() {

	//ctx := context.Background()

	/*не умеешь юзать контексты?*/

	/* зачем ты юзаешь джин и стандартный http*/

	pgDB, err := db.NewPostgresPool("postgres://krp_auth:krp_auth@postgres_auth:5432/krp_auth")
	if err != nil {
		logrus.Panicf("unable get postgres pool: %v", err)
	}

	userRepo, err := repositories.NewUserRepository(pgDB)
	if err != nil {
		logrus.Panicf("unable build user repo: %v", err)
	}

	userService, err := services.NewUserService(userRepo)
	if err != nil {
		logrus.Panicf("unable build user service: %v", err)
	}

	userHandler, err := user.NewUserHandler(userService)
	if err != nil {
		logrus.Panicf("unable build user handler: %v", err)
	}

	srv := new(server.Server)
	if err := srv.Run("8080", userHandler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}
