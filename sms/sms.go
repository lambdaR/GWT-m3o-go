package sms

import (
	"go.m3o.com/client"
)

type Sms interface {
	Send(*SendRequest) (*SendResponse, error)
}

func NewSmsService(token string) *SmsService {
	return &SmsService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type SmsService struct {
	client *client.Client
}

// Send an SMS.
func (t *SmsService) Send(request *SendRequest) (*SendResponse, error) {

	rsp := &SendResponse{}
	return rsp, t.client.Call("sms", "Send", request, rsp)

}

type SendRequest struct {
	// who is the message from? The message will be suffixed with "Sent from <from>"
	From string `json:"from,omitempty"`
	// the main body of the message to send
	Message string `json:"message,omitempty"`
	// the destination phone number including the international dialling code (e.g. +44)
	To string `json:"to,omitempty"`
}

type SendResponse struct {
	// any additional info
	Info string `json:"info,omitempty"`
	// will return "ok" if successful
	Status string `json:"status,omitempty"`
}
