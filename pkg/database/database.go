package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// InitDB initializes the SQLite database and returns the connection.
func InitDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", "./notifications.db")
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

// Get retrieves all notifications from the database
func Get(db *sqlx.DB) ([]Notification, error) {
	var notifications []Notification

	if err := db.Select(&notifications, "SELECT * FROM notifications ORDER BY time DESC LIMIT 100"); err != nil {
		log.Println("Error fetching notifications:", err)

		return nil, err
	}

	return notifications, nil
}

// Save stores the notification in the database.
func Save(db *sqlx.DB, source, title, message string) error {
	_, err := db.Exec(
		"INSERT INTO notifications (source, title, message, time) VALUES (?, ?, ?, datetime('now'))",
		source, title, message)

	if err != nil {
		log.Printf("Failed to save notification: %v\n", err)
	}

	return err
}
