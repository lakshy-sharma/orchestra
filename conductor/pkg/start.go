package pkg

import (
	"conductor/pkg/api"
	"conductor/pkg/db"
	"conductor/pkg/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func initRouter(database *db.Database) *gin.Engine {
	router := gin.Default()
	return router
}

func StartConductor(conductorConfig db.ConductorConfig, logger *zap.Logger) {
	historianConnection, err := db.NewDatabase("127.0.0.1:6379", "", logger)
	if err != nil {
		logger.Fatal("Failed to connect to with the database. Shutting Down Conductor.")
	}

	router := initRouter(historianConnection)

	// Version 1 APIs
	v1 := router.Group("/v1")
	{
		// NON-Auth APIs.

		// Authenticated APIs
		v1.POST("/deployment/register", middleware.AuthMiddleware(), api.RegisterDeployment(historianConnection))
		v1.POST("/musician/register", middleware.AuthMiddleware(), api.RegisterMusician(historianConnection))
	}

	// Start the router
	router.Run("localhost:8080")
}
