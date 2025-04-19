# Notification Service

## Overview

A simple notification service built using Go, Gin framework, and SQLite.

## Features

- Save notifications with `source`, `title`, and `message`.
- Retrieve and display notifications in JSON format or on a webpage.
- SQLite database for persistent storage.
- Dockerized for easy deployment.

## Project Structure

- **`main.go`**: Entry point of the application.
- **`pkg/api/server.go`**: Initializes the Gin server and sets up routes.
- **`pkg/database/`**: Contains database models and functions for interacting with SQLite.
- **`pkg/handlers/`**: Defines API handlers for creating and displaying notifications.
- **`templates/`**: HTML templates for rendering notifications on the webpage.
- **`migration.sh`**: Script to initialize the SQLite database and create the required tables.

## API Endpoints

### 1. Create Notification

- **URL**: `POST /api/notifications`
- **Request Body**

  ```json
  {
    "source": "WhatsApp",
    "title": "New Message",
    "message": "I'm IronMan"
  }

- **Response**

    _200 OK_: Notification saved successfully. \
    _400 Bad Request_: Invalid notification data. \
    _500 Internal Server Error_: Failed to save notification.

### 2. View Notifications

- **URL**: `GET /`
- **Response**

    _200 OK_: Renders a webpage displaying the list of notifications.

## How to Run

### 1. Clone the repository

```bash
git clone <repository-url>
cd notification-service
```

### 2. Run the migration script

./migration.sh

### 3. run the application

You can run the application using debugger or
using,

```bash
go run main.go
```

## Using Docker

### 1. Build the Docker image

```bash
docker build -t notification-service .
```

### 2. Run the Docker container

```bash
docker run -p 8089:8089 notification-service
```
