package config

import "time"

var ConnCheckerVersion string

// C holds the global configuration.
var Conf Config

type Config struct {
	General struct {
		LogLevel int `mapstructure:"log_level"`
	} `mapstructure:"general"`

	PaymentServer struct {
		PaymentServiceAddress string        `mapstructure:"payment_service_address"`
		PaymentServicePort    string        `mapstructure:"payment_service_port"`
		PaymentSendAccount    string        `mapstructure:"payment_account"`
		PaymentSendTime       time.Duration `mapstructure:"payment_send_minutes"`
	} `mapstructure:"paymentserver"`

	M2MServer struct {
		M2MServiceAddress string        `mapstructure:"m2m_service_address"`
		M2MServicePort    string        `mapstructure:"m2m_service_port"`
		M2MSendTime       time.Duration `mapstructure:"m2m_send_minutes"`
	} `mapstructure:"m2mserver"`

	AppServer struct {
		AppServiceAddress string        `mapstructure:"app_service_address"`
		AppServicePort    string        `mapstructure:"app_service_port"`
		AppSendTime       time.Duration `mapstructure:"app_send_minutes"`
	} `mapstructure:"appserver"`

	NsServer struct {
		NsServiceAddress string        `mapstructure:"ns_service_address"`
		NsServicePort    string        `mapstructure:"ns_service_port"`
		NsSendTime       time.Duration `mapstructure:"ns_send_minutes"`
	} `mapstructure:"networkserver"`

	SmtpServer struct {
		SmtpHost     string `mapstructure:"smtp_host"`
		SmtpSender   string `mapstructure:"smtp_sender"`
		SmtpPassword string `mapstructure:"smtp_password"`
		SmtpPort     int    `mapstructure:"smtp_port"`
	} `mapstructure:"smtpserver"`
}
