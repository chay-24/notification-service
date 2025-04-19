package database

// Notification struct represents a single notification.
type Notification struct {
	ID        int    `json:"id" db:"id"`
	Source    string `json:"source" db:"source"`
	Title     string `json:"title" db:"title"`
	Message   string `json:"message" db:"message"`
	Timestamp string `json:"timestamp" db:"timestamp"`
}
