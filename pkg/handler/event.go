package handler

import (
	"encoding/json"
	"github.com/asstrahanec/go-clickhouse"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary createEvent
// @Description Create a new event
// @ID create-event
// @Accept  json
// @Produce  json
// @Param input body go_clickhouse.Event true "event info. eventTime (in ISO format, e.g., 2024-09-01T00:00:00Z"
// }"
// @Success 200 {object} go_clickhouse.Event
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/event [post]
func (h *Handler) createEvent(c *gin.Context) {
	var input go_clickhouse.Event

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if !isValidJSON(input.Payload) {
		NewErrorResponse(c, http.StatusBadRequest, "Invalid JSON in payload")
		return
	}

	err := h.services.CreateEvent(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": input})
}

func isValidJSON(payload string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(payload), &js) == nil
}

//type GetAllEventsResponse struct {
//	Events []go_clickhouse.Event `json:"events"`
//}

// @Summary getEvents
// @Description Get events filtered by eventType and time range
// @ID get-events
// @Accept  json
// @Produce  json
// @Param eventType query string true "Event Type"
// @Param startTime query string true "Start Time (in ISO format, e.g., 2024-09-01T00:00:00)"
// @Param endTime query string true "End Time (in ISO format, e.g., 2024-09-07T23:59:59)"
// @Success 200 {array} go_clickhouse.Event
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/event [get]
func (h *Handler) getEvents(c *gin.Context) {
	eventType := c.Query("eventType")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")

	if eventType == "" || startTime == "" || endTime == "" {
		NewErrorResponse(c, http.StatusBadRequest, "Missing query parameters")
		return
	}

	events, err := h.services.GetEvents(eventType, startTime, endTime)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"events": events})
}
