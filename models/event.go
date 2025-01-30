package models

import (
	"example.com/event-booking-api/db"
	"time"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	Location    string    `binding:"required" json:"location"`
	Datetime    time.Time `binding:"required" json:"datetime"`
	UserId      int
}

var events = []Event{}

func (e *Event) Save() error {
	query := `
  INSERT INTO events(
    name,
    description,
    location,
    datetime,
    userId
  ) VALUES (?, ?, ?, ?, ?)
  `
	stmt, qErr := db.Db.Prepare(query)
	if qErr != nil {
		return qErr
	}
	defer stmt.Close()

	res, exErr := stmt.Exec(e.Name, e.Description, e.Location, e.Datetime, e.UserId)
	if exErr != nil {
		return exErr
	}
	createdId, idErr := res.LastInsertId()
	if idErr != nil {
		return idErr
	}

	e.Id = createdId

	return nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.Db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		rsErr := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.Datetime, &event.UserId)
		if rsErr != nil {
			return nil, rsErr
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE Id = ?"
	qRow := db.Db.QueryRow(query, id)

	var event Event
	sErr := qRow.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.Datetime, &event.UserId)
	if sErr != nil {
		return nil, sErr
	}

	return &event, nil
}

func (e *Event) Update() error {
	query := `
  UPDATE events
  SET 
  name = ?, 
  description = ?,
  location = ?,
  datetime = ?
  WHERE id = ?
  `
	stmt, err := db.Db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.Datetime, e.Id)

	return err
}

func (e *Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"

	stmt, err := db.Db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, eErr := stmt.Exec(e.Id)

	return eErr
}
