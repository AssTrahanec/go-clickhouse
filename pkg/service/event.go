package service

import (
	"github.com/asstrahanec/go-clickhouse"
	"github.com/asstrahanec/go-clickhouse/pkg/repository"
)

type EventService struct {
	repo *repository.Repository
}

func NewEventService(repo *repository.Repository) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) CreateEvent(event go_clickhouse.Event) error {
	return s.repo.CreateEvent(event)
}

func (s *EventService) GetEvents(eventType, startTime, endTime string) ([]go_clickhouse.Event, error) {
	return s.repo.GetEvents(eventType, startTime, endTime)
}
