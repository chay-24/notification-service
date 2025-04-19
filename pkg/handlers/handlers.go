package handlers

import (
	"log"
	"net/http"
	"notification-service/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// SetupRoutes sets up the routes and handlers for the application.
func SetupRoutes(r *gin.Engine, db *sqlx.DB) {
	r.POST("/api/notifications", createNotification(db))

	r.GET("/", showNotifications(db))
}

// createNotification handles saving notifications to the database.
func createNotification(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var notification database.Notification
		if err := c.ShouldBindJSON(&notification); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid notification data"})

			return
		}

		if err := database.Save(db, notification.Source, notification.Title, notification.Message); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save notification"})

			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "Notification saved"})
	}
}

// showNotifications fetches notifications and displays them on a webpage.
func showNotifications(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		notifications, err := database.Get(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching notifications"})
			log.Println("Error fetching notifications: ", err)
			return
		}

		if notifications == nil {
			notifications = []database.Notification{}
		}

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"notifications": notifications,
		})
	}
}
