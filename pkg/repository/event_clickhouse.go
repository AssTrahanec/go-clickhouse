package repository

import (
	"github.com/asstrahanec/go-clickhouse"
	"github.com/jmoiron/sqlx"
)

type EventClickHouse struct {
	db *sqlx.DB
}

func NewEventClickHouse(db *sqlx.DB) *EventClickHouse {
	return &EventClickHouse{db: db}
}

func (r *EventClickHouse) CreateEvent(event go_clickhouse.Event) error {
	insertEventQuery := `INSERT INTO events (eventID, eventType, userID, eventTime, payload)
    VALUES (?, ?, ?, ?, ?)`

	_, err := r.db.Exec(insertEventQuery, event.EventID, event.EventType, event.UserID, event.EventTime, event.Payload)
	if err != nil {
		return err
	}

	return nil
}
func (r *EventClickHouse) GetEvents(eventType, startTime, endTime string) ([]go_clickhouse.Event, error) {
	var events []go_clickhouse.Event
	query := `SELECT eventID, eventType, userID, eventTime, payload
	          FROM events
	          WHERE eventType = ? AND eventTime BETWEEN ? AND ?`
	err := r.db.Select(&events, query, eventType, startTime, endTime)
	if err != nil {
		return nil, err
	}

	return events, nil
}
