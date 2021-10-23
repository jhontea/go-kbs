package gateway

type SMSGateway interface {
	SendSMS(to, message string) error
}
