#!/bin/bash

DB_NAME="notifications.db"
SQLITE_CMD="sqlite3"

create_table() {
  local table_name=$1
  local create_sql=$2

  echo "Creating table '$table_name' in database '$DB_NAME'..."
  echo "$create_sql" | $SQLITE_CMD $DB_NAME
  echo "Table '$table_name' created successfully."
}

# Table schema for Notification model.
create_notification_table_sql="
CREATE TABLE IF NOT EXISTS notifications (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    source TEXT NOT NULL,
    title TEXT NOT NULL,
    message TEXT NOT NULL,
    timestamp TEXT NOT NULL
);
"

if [ ! -f "$DB_NAME" ]; then
  echo "Database '$DB_NAME' does not exist. Creating it..."
  touch $DB_NAME
fi

create_table "notifications" "$create_notification_table_sql"
