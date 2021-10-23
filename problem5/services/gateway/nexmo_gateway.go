package gateway

import (
	"fmt"
	"sync"
)

type NexmoGateway struct {
}

var (
	nexMoonce    sync.Once
	nexmoGateway *NexmoGateway
)

func NewNexmoGateway() SMSGateway {
	nexMoonce.Do(func() {
		nexmoGateway = &NexmoGateway{}
	})

	return nexmoGateway
}

func (g *NexmoGateway) SendSMS(to, message string) error {
	fmt.Println("send sms from Nexmo gateway")
	return nil
}
