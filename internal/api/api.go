package api

import (
	"github.com/pkg/errors"
	"gitlab.com/MXCFoundation/cloud/conn_checker/internal/config"
	"gitlab.com/MXCFoundation/cloud/conn_checker/internal/service/appserver"
	"gitlab.com/MXCFoundation/cloud/conn_checker/internal/service/m2m"
	"gitlab.com/MXCFoundation/cloud/conn_checker/internal/service/networkserver"
	"gitlab.com/MXCFoundation/cloud/conn_checker/internal/service/payment"
	"time"
)

func Setup(conf config.Config) error {
	if err := payment.Setup(); err != nil {
		return errors.Wrap(err, "setup payment service api error")
	}

	if err := networkserver.Setup(); err != nil {
		return errors.Wrap(err, "setup network server api error")
	}

	if err := m2m.Setup(); err != nil {
		return errors.Wrap(err, "setup M2M api error")
	}

	if err := appserver.Setup(); err != nil {
		return errors.Wrap(err, "setup AppServer api error")
	}

	go paymentConnTest(conf)
	go m2mConnTest(conf)
	go nsConnTest(conf)
	go asConnTest(conf)

	return nil
}

func paymentConnTest(conf config.Config) {
	for {
		time.Sleep(conf.PaymentServer.PaymentSendTime * time.Second)
		if err := payment.PaymentConnTest(&conf); err != nil{
			//email.SendError()
		}
	}
}

func m2mConnTest(conf config.Config) {
	for {
		time.Sleep(conf.M2MServer.M2MSendTime * time.Second)
		if err := m2m.M2MServerConnTest(&conf); err != nil{
			//email.SendError()
		}
	}
}

func nsConnTest(conf config.Config) {
	for {
		time.Sleep(conf.NsServer.NsSendTime * time.Second)
		if err := networkserver.NsServerConnTest(&conf); err != nil{
			//email.SendError()
		}
	}
}

func asConnTest(conf config.Config) {
	for {
		time.Sleep(conf.AppServer.AppSendTime * time.Second)
		if err := appserver.AppServerConnTest(&conf); err != nil{
			//email.SendError()
		}
	}
}