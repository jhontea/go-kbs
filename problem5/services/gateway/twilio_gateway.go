package gateway

import (
	"fmt"
	"sync"
)

type TwilioGateway struct {
}

var (
	twilioOnce    sync.Once
	twilioGateway *TwilioGateway
)

func NewTwilioGateway() SMSGateway {
	twilioOnce.Do(func() {
		twilioGateway = &TwilioGateway{}
	})
	return twilioGateway
}

func (g *TwilioGateway) SendSMS(to, message string) error {
	fmt.Println("send sms from twilio gateway")
	return nil
}
