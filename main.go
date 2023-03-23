package main

import (
	"testAPI/common/env"
	"testAPI/controller"
	"testAPI/repository"
	"testAPI/router"
	"testAPI/service"
)

func main() {
	env.Load()
	router := setupAllDependencies()

	router.Run(env.URL)
}

func setupAllDependencies() router.Router {
	repository := repository.NewRepository()
	service := service.NewService(repository)
	controller := controller.NewController(service)
	router := router.NewRouter(controller)

	return router
}
