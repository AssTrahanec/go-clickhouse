package service

import (
	"github.com/asstrahanec/go-clickhouse"
	"github.com/asstrahanec/go-clickhouse/pkg/repository"
)

type Event interface {
	CreateEvent(event go_clickhouse.Event) error
	GetEvents(eventType, startTime, endTime string) ([]go_clickhouse.Event, error)
}

type Service struct {
	Event
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Event: NewEventService(repos),
	}
}
