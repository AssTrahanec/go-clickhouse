package repository

import (
	"github.com/asstrahanec/go-clickhouse"
	"github.com/jmoiron/sqlx"
)

type Event interface {
	CreateEvent(event go_clickhouse.Event) error
	GetEvents(eventType, startTime, endTime string) ([]go_clickhouse.Event, error)
}

type Repository struct {
	Event
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Event: NewEventClickHouse(db),
	}
}
