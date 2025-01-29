package models

import (
	"time"

	"example.com/event-booking-api/db"
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
	dbQuery, stErr := db.Db.Prepare(query)
	if stErr != nil {
		return stErr
	}
	defer dbQuery.Close()

	res, exErr := dbQuery.Exec(e.Name, e.Description, e.Location, e.Datetime, e.UserId)
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
