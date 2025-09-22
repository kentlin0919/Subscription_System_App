package main

import (
	adapterhttp "subscription_system_app/internal/adapter/http"
	dbinfra "subscription_system_app/internal/infra/db"
	loginfra "subscription_system_app/internal/infra/log"
	"subscription_system_app/internal/usecase"
)

func main() {
	logger := loginfra.NewStandardLogger()
	pg := dbinfra.NewPostgres(nil)
	repo := dbinfra.NewUserRepository(pg)
	creator := usecase.NewUserCreator(repo)
	handler := adapterhttp.NewUserHandler(creator)
	router := adapterhttp.NewRouter(handler)

	logger.Printf("subscription service listening on :8080")

	if err := router.Run(":8080"); err != nil {
		logger.Printf("http server stopped: %v", err)
	}
}
