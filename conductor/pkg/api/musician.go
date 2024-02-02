package api

import (
	"conductor/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterMusician(historianConnection *db.Database) gin.HandlerFunc {
	// Create a node in database and pass back the credentials for that node in the response.
	return func(context *gin.Context) {
		musician := db.Node{}

		if err := context.BindJSON(&musician); err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}

		context.JSON(http.StatusAccepted, &musician)

	}
}
