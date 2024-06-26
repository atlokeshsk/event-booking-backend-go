package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/atlokeshsk/event-booking/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      int64     `json:"user_id"`
}

// Save inserts a new event into the database and updates the event's ID with the last inserted ID.
// It returns an error if any database operation fails.
func (e *Event) Save() error {
	query := `INSERT INTO events(name,description,location,datetime,user_id)
              VALUES(?,?,?,?,?)
              `
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id
	return nil
}

func (e *Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?
	WHERE id = ?	
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.ID)
	return err
}

func (e *Event) Delete() error {
	query := `
	DELETE FROM events WHERE id = ?
`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID)
	return err
}

// GetAllEvents retrieves all events from the database.
// It returns a slice of Event structs and an error if any occurs during the query execution or row scanning.
func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events = []Event{}
	for rows.Next() {
		var event Event
		err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

// GetEventById retrieves an event from the database by its ID.
// It returns a pointer to the Event struct and an error if any occurs.
// If no event is found for the given ID, it returns a custom error message.
func GetEventById(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id=?`
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("no event present for the given id")
		}
		return nil, err
	}
	return &event, nil
}

func (e *Event) Register(userID int64) error {
	query := `
		INSERT INTO registrations(event_id,user_id)
		VALUES(?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID, userID)
	return err
}

func (e *Event) CancelRegistration(userID int64) error {
	query := `
		DELETE FROM registrarions WHERE event_id = ? AND user_id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID, userID)
	return err

}
