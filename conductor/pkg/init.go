package pkg

import (
	"conductor/pkg/db"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func InitializeCondutor() {
	// Install any pre-requisites.
	// Setup Required Services.

	// Setup a connection with the database.
	dbConn, err := db.NewDatabase(fmt.Sprintf("%s:%s", viper.GetString("historian.redis_cluster_endpoint"), viper.GetString("historian.redis_cluster_port")))
	if err != nil {
		log.Fatal(err)
	}

}
