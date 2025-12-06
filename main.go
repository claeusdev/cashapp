package main

import (
	"cashapp/core"
	"cashapp/core/database"

	"cashapp/repository"
	"cashapp/routes"
	"cashapp/services"

	"cashapp/models"

	"go.uber.org/zap"
)

// @title CashApp API
// @version 1.0
// @description A payment processing API built with Go and Gin
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@cashapp.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:5454
// @BasePath /
// @BasePath /
func main() {
	config := core.NewConfig()
	core.InitLogger(config.ENVIRONMENT)

	pg, err := database.NewPostgres(config)
	if err != nil {
		core.Log.Fatal("failed to initialize postgres database", zap.Error(err))
	}

	err = database.RunMigrations(pg, &models.Transaction{}, &models.User{}, &models.Wallet{})
	if err != nil {
		core.Log.Fatal("failed to run migrations", zap.Error(err))
	}

	if config.RUN_SEEDS {
		models.RunSeeds(pg)
	}

	cache := database.NewRedis(config)
	repo := repository.NewRepository(pg)
	service := services.NewService(repo, cache, config)

	server := core.NewHTTPServer(config)
	router := routes.NewRouter(server.Engine, config, service)

	router.RegisterRoutes()
	server.Start()

}
