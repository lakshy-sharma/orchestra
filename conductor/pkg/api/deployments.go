package api

import (
	"conductor/pkg/db"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterDeployment(historianConnection *db.Database) gin.HandlerFunc {
	// Create a simple entry inside the historian.
	return func(context *gin.Context) {
		deployment := db.Deployment{}
		// Perform a simple parsing check to ensure that data is valid JSON.
		if err := context.BindJSON(&deployment); err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}
		// Convert data into string and Register the deployment into redis queue.
		deploymentJsonString, err := json.Marshal(deployment)
		if err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}
		historianConnection.Client.LPush("deployments", string(deploymentJsonString))

		// Return Success
		context.JSON(http.StatusAccepted, &deployment)
	}
}
