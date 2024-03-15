package pkg

import (
	"conductor/pkg/api"
	"conductor/pkg/db"
	"conductor/pkg/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func StartConductor(conductorConfig db.ConductorConfig, logger *zap.Logger) {
	historianConnection, err := db.NewDatabase(
		fmt.Sprintf("%s:%d", viper.GetString("historian.redis_cluster_endpoint"), viper.GetInt("historian.redis_cluster_port")),
		viper.GetString("historian.redis_password"),
		logger,
	)
	if err != nil {
		logger.Fatal("Failed to connect to with the database. Shutting Down Conductor.")
	}

	router := gin.Default()

	// Version 1 APIs
	v1Secure := router.Group("/v1").Group("/secure")
	v1Insecure := router.Group("/v1").Group("/insecure")
	{
		// NON-Authenticated APIs.
		v1Insecure.GET("/status", api.ClusterStatus)

		// Authenticated APIs
		v1Secure.POST("/deployment/register", middleware.AuthMiddleware(), api.RegisterDeployment(historianConnection))
		v1Secure.POST("/musician/register", middleware.AuthMiddleware(), api.RegisterMusician(historianConnection))
	}

	// Start the router
	router.Run("localhost:8080")
}
