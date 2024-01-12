package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "saglyk-backend/docs"
	"saglyk-backend/internal/config"
	handlermanager "saglyk-backend/internal/handlers/manager"
	"saglyk-backend/pkg/client/postgresql"
	"saglyk-backend/pkg/logging"
)

// @title Admin Course APIs
// @version 1.0
// @description Testing Swagger APIs.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support

// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8081
// @BasePath /v1

// @Success     200  {object} appresult.AppSuccess  "Success Code"
// @Failure     400  {object} appresult.AppError  "Error Code"
// @schemes http
func main() {

	cfg := config.GetConfig()
	logger := logging.GetLogger()

	postgresSQLClient := startPostgresql(cfg, logger)
	start(handlermanager.Manager(postgresSQLClient, logger), cfg)
}

func startPostgresql(cfg *config.Config, logger *logging.Logger) *pgxpool.Pool {
	postgresSQLClient, err := postgresql.NewClient(context.TODO(), 3, cfg.Storage)
	if err != nil {
		return nil
	}
	return postgresSQLClient
}

func start(router *gin.Engine, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("start application")

	router.StaticFS("/public", gin.Dir(cfg.PublicFilePath, false))
	ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(`:` + cfg.Listen.Port)

}
