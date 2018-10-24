package revolt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEvent(t *testing.T) {
	tests := []struct {
		name      string
		eventType string
		data      interface{}
		expError  error
	}{
		{
			name:      "Event test 1",
			eventType: "Event.test.type",
			data: struct {
				Key   string
				Value string
			}{
				Key:   "key",
				Value: "value",
			},
			expError: nil,
		},
		{
			name:      "Event with empty struct data",
			eventType: "Event.test.type",
			data:      struct{}{},
			expError:  ErrEventData,
		},
		{
			name:      "Event with nil data",
			eventType: "Event.test.type",
			data:      nil,
			expError:  ErrEventData,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			newEvent, err := NewEvent(test.eventType, test.data)

			if test.expError != nil {
				assert.Equal(t, err, test.expError)
			} else {
				assert.Equal(t, newEvent.Meta.Type, test.eventType)
				assert.Equal(t, newEvent.Data, test.data)
				assert.NotEmpty(t, newEvent.Meta.ID)
				assert.NotEmpty(t, newEvent.Meta.Timestamp)
			}
		})
	}
}
