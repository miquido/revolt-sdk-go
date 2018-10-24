package revolt

import (
	"time"

	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

var ErrEventType = errors.New("Event type cannot be empty")
var ErrEventData = errors.New("Event data cannot be nil")

// Event is a representation of a message sent to Revolt API
type Event struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

// Meta is a representation of Meta tags
type Meta struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Timestamp int64  `json:"timestamp"`
}

// NewEvent creates a new instance of a Event with given type and embedded data.
func NewEvent(eventType string, data interface{}) (Event, error) {
	eventId, err := uuid.NewV1()
	if err != nil {
		return Event{}, err
	}

	ev := Event{
		Meta: Meta{
			ID:        eventId.String(),
			Type:      eventType,
			Timestamp: unixMilliseconds(),
		},
		Data: data,
	}

	err = ev.validate()
	if err != nil {
		return Event{}, err
	}

	return ev, nil
}

func unixMilliseconds() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func (e Event) validate() error {
	if e.Meta.Type == "" {
		return ErrEventType
	}

	if e.Data == nil || e.Data == struct{}{} {
		return ErrEventData
	}

	return nil
}
