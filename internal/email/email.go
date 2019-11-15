package email

import (
	"github.com/sirupsen/logrus"
	"github.com/zbindenren/logrus_mail"
	"gitlab.com/MXCFoundation/cloud/conn_checker/internal/config"
	"time"
)

var (
	sender     string
	password   string
	host       string
	smtpPort   int
)

var sethook *logrus_mail.MailAuthHook

func Setup(c config.Config) error {
	sender = c.SmtpServer.SmtpSender
	password = c.SmtpServer.SmtpPassword
	smtpPort = c.SmtpServer.SmtpPort
	host = c.SmtpServer.SmtpHost

	hook, err := logrus_mail.NewMailAuthHook("Conn-Checker", host, smtpPort,
		sender, "shiun.chen@mxc.org", sender, password)
	if err != nil {
		logrus.WithError(err).Error("emailHook set up error")
		return err
	}

	sethook = hook

	logrus.AddHook(hook)

	return nil
}

func SendError(service, err string) error {

	contextLogger := logrus.WithFields(logrus.Fields{
		"Service": service,
		"Error": err,
		"Time": time.Now().String(),
	})

	contextLogger.Time = time.Now()
	contextLogger.Message = "E-Mail from Conn-Checker"
	contextLogger.Level = logrus.FatalLevel

	if err := sethook.Fire(contextLogger); err != nil {
		logrus.WithError(err).Error("SendError/Cannot send the mail.")
		return err
	}

	return nil
}
