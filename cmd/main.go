package main

import (
	"github.com/sirupsen/logrus"
	"krp_project/internal/db"
	"krp_project/internal/handlers/user"
	"krp_project/internal/repositories"
	"krp_project/internal/server"
	"krp_project/internal/services"
	"log"
)

func main() {

	//ctx := context.Background()
	pgDB, err := db.NewPostgresPool("postgres://postgres:12345@localhost:5432/krp_auth")
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
