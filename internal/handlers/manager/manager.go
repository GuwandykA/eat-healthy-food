package handlermanager

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag"
	"net/http"
	"saglyk-backend/internal/admin/category"
	catdb "saglyk-backend/internal/admin/category/db"
	"saglyk-backend/pkg/logging"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	categoryURL = "/v1/admin/categories"
)

func Manager(client *pgxpool.Pool, logger *logging.Logger) *gin.Engine {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", http.MethodGet, http.MethodPatch, http.MethodPost},
		AllowHeaders: []string{
			"Content-Type", "Content-Length", "Accept-Encoding",
			"Authorization", "Cache-Control",
			"access-control-allow-origin", "access-control-allow-headers",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	categoryRouterManager := router.Group(categoryURL)
	categoryRepository := catdb.NewRepository(client, logger)
	categoryRouterHandler := category.NewHandler(categoryRepository, logger)
	categoryRouterHandler.Register(categoryRouterManager)

	return router
}
