package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB initializes the database connection and sets up the database.
// It opens a connection to a SQLite database located at "./api.db".
// The function also configures the maximum number of open and idle connections
// and calls createTables to set up the necessary database tables.
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./api.db")
	if err != nil {
		panic(err.Error())
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

// createTables creates the necessary tables for the application if they do not already exist.
// It creates three tables: users, events, and registrations.
// The users table stores user information.
// The events table stores event information and references the users table.
// The registrations table stores event registrations and references both the users and events tables.
func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic(err.Error())
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		datetime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)	
	`
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic(err.Error())
	}

	createRegistrationTable := `
		CREATE TABLE IF NOT EXISTS registrations(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id)
		FOREIGN KEY(user_id ) REFERENCES users(id)
		)
	`
	_, err = DB.Exec(createRegistrationTable)
	if err != nil {
		panic(err.Error())
	}

}
