package go_clickhouse

import "time"

type Event struct {
	EventID   int       `json:"-" db:"eventID"`
	EventType string    `json:"eventType" binding:"required" db:"eventType"`
	UserID    int       `json:"userID" binding:"required" db:"userID"`
	EventTime time.Time `json:"eventTime" binding:"required" db:"eventTime"`
	Payload   string    `json:"payload" binding:"required" db:"payload"`
}
