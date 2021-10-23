package gateway

import (
	"fmt"
	"sync"
)

type Vendor3Gateway struct {
}

var (
	vendor3Once    sync.Once
	vendor3Gateway *Vendor3Gateway
)

func NewVendor3Gateway() SMSGateway {
	vendor3Once.Do(func() {
		vendor3Gateway = &Vendor3Gateway{}
	})
	return vendor3Gateway
}

func (g *Vendor3Gateway) SendSMS(to, message string) error {
	fmt.Println("send sms from vendor3 gateway")
	return nil
}
