package models

import (
	"database/sql"
	"fmt"
	"max-tuts/event-booking-rest-api/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	query := `
		INSERT INTO events(name, description, location, dateTime, user_id) 
		VALUES(?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println("Error closing statement:", err)
		}
	}(stmt)
	if err != nil {
		return err // exit if the query failed
	}
	exec, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err // exit if the execution failed
	}
	id, err := exec.LastInsertId()
	if err != nil {
		return err // exit if the id retrieval failed
	}

	e.ID = id
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		println("Error querying events table:", err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows:", err)
		}
	}(rows)
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id=?"
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (e Event) Update() error {
	query := "UPDATE events SET name=?, description=?, location=?, dateTime=? WHERE id=?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println("Error preparing update statement:", err)
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println("Error closing statement:", err)
		}
	}(stmt)
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	return err
}

func (e Event) Delete() error {
	query := "DELETE FROM events WHERE id=?"
	smtp, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println("Error preparing delete statement:", err)
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println("Error closing statement:", err)
		}
	}(smtp)
	_, err = smtp.Exec(e.ID)
	return err
}
