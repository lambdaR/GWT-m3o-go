package event

import (
	"go.m3o.com/client"
)

type Event interface {
	Consume(*ConsumeRequest) (*ConsumeResponseStream, error)
	Publish(*PublishRequest) (*PublishResponse, error)
	Read(*ReadRequest) (*ReadResponse, error)
}

func NewEventService(token string) *EventService {
	return &EventService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type EventService struct {
	client *client.Client
}

// Consume events from a given topic.
func (t *EventService) Consume(request *ConsumeRequest) (*ConsumeResponseStream, error) {
	stream, err := t.client.Stream("event", "Consume", request)
	if err != nil {
		return nil, err
	}
	return &ConsumeResponseStream{
		stream: stream,
	}, nil

}

type ConsumeResponseStream struct {
	stream *client.Stream
}

func (t *ConsumeResponseStream) Recv() (*ConsumeResponse, error) {
	var rsp ConsumeResponse
	if err := t.stream.Recv(&rsp); err != nil {
		return nil, err
	}
	return &rsp, nil
}

// Publish a event to the event stream.
func (t *EventService) Publish(request *PublishRequest) (*PublishResponse, error) {

	rsp := &PublishResponse{}
	return rsp, t.client.Call("event", "Publish", request, rsp)

}

// Read stored events
func (t *EventService) Read(request *ReadRequest) (*ReadResponse, error) {

	rsp := &ReadResponse{}
	return rsp, t.client.Call("event", "Read", request, rsp)

}

type ConsumeRequest struct {
	// The topic to subscribe to
	Topic string `json:"topic,omitempty"`
	// Optional group for the subscription
	Group string `json:"group,omitempty"`
	// Optional offset to read from e.g "2006-01-02T15:04:05.999Z07:00"
	Offset string `json:"offset,omitempty"`
}

type ConsumeResponse struct {
	// Unique message id
	Id string `json:"id,omitempty"`
	// The next json message on the topic
	Message map[string]interface{} `json:"message,omitempty"`
	// Timestamp of publishing
	Timestamp string `json:"timestamp,omitempty"`
	// The topic subscribed to
	Topic string `json:"topic,omitempty"`
}

type Ev struct {
	// event id
	Id string `json:"id,omitempty"`
	// event message
	Message map[string]interface{} `json:"message,omitempty"`
	// event timestamp
	Timestamp string `json:"timestamp,omitempty"`
}

type PublishRequest struct {
	// The json message to publish
	Message map[string]interface{} `json:"message,omitempty"`
	// The topic to publish to
	Topic string `json:"topic,omitempty"`
}

type PublishResponse struct {
}

type ReadRequest struct {
	// number of events to read; default 25
	Limit int32 `json:"limit,omitempty"`
	// offset for the events; default 0
	Offset int32 `json:"offset,omitempty"`
	// topic to read from
	Topic string `json:"topic,omitempty"`
}

type ReadResponse struct {
	// the events
	Events []Ev `json:"events,omitempty"`
}
