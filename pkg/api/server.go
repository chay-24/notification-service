package api

import (
	"log"
	"notification-service/pkg/database"
	"notification-service/pkg/handlers"

	"github.com/gin-gonic/gin"
)

// Start initializes and runs the Gin server
func Start() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Could not connect to the database: ", err)
	}

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	handlers.SetupRoutes(r, db)

	r.Run(":8089")
}
