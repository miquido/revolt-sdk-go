# Revolt Golang SDK
Go SDK for event tracking with Revolt.

## Installation
Use the go command:
```
go run -u bitbucket.org/miquido/revolt-sdk-go
```

## Creating events
Event can be created with `revolt.NewEvent()` method. 

### Using struct
Example:
```
revolt.NewEvent("test.event.type", struct {
    UserId       int    `json:"userId"`
    CreationType string `json:"creationType"`
    Description  string `json:"description"`
}{
    UserId:       1,
    CreationType: "webbapp.test",
    Description:  "short description",
})
```
When using struct{} don't forget to set explicit json tags for format purposes. 

### Using map
Example:
```
revolt.NewEvent("test.event.type",
    map[string]interface{}{
        "userId": 5,
        "description": "short description",
    },
)
```

## Example
Example usage

	client, err := revolt.NewClient("trackingId", "app.code", "secret")
	if err != nil {
	    panic(err)
	}

	event, err := revolt.NewEvent("test.event.type", struct {
		UserId       int    `json:"userId"`
	}{
		UserId:       1,
	})

	event2, err := revolt.NewEvent("test.event.type",
		map[string]interface{}{
			"userId": 5,
		},
	)


	resp, err := client.SendEvents([]revolt.Event{event, event2})

## Async event sending
For async event sending use channels and goroutines.

# Custom Parameters
There are few parameters which will be supported soon:

- [x] endpoint - Specifies endpoint on which communication with Revolt service should take place.
- [ ] maxBatchSize - Specifies maximum batch size of events that can be sent to Revolt API.
- [ ] eventDelayMillis - Specifies maximum number of seconds for Event to be stored in queue. After delay is up, all events in queue will be sent automatically.
- [ ] offlineMaxSize - Specifies maximum size of events that can be stored in queue when service does not respond.
- [ ] retryIntervalSecs - Specifies first time interval in seconds to retry sending batch of events when any error occurs.
- [ ] maxRetryIntervalSecs - Specifies maximum time interval in seconds to retry sending batch of events when any error occurs.



