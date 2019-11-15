package payment

import (
	"context"
	"crypto/rand"
	"github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/conn_checker/internal/config"
	"gitlab.com/MXCFoundation/cloud/conn_checker/internal/email"
	"gitlab.com/MXCFoundation/cloud/conn_checker/pkg/api/payment"
	"google.golang.org/grpc"
	"math/big"
)

func Setup() error {
	logrus.Info("Setup Payment API")
	return nil
}

func connPayment(address, port string) (*grpc.ClientConn, error) {
	logrus.WithFields(logrus.Fields{
		"address": address,
		"port": port,
	}).Debug("payment/connPayment")

	conn, err := grpc.Dial(address + port, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func PaymentConnTest(conf *config.Config) error {
	conn, err := connPayment(conf.PaymentServer.PaymentServiceAddress, conf.PaymentServer.PaymentServicePort)
	if err != nil {
		logrus.WithError(err).Error("Cannot conn to payment service.")
		email.SendError("Payment", err.Error())
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			logrus.WithError(err).Error("Cannot close payment connection.")
		}
	}()

	client := payment.NewPaymentClient(conn)

	reqIdclient, _ := rand.Int(rand.Reader, big.NewInt(99999999))

	_, err = client.TokenTxReq(context.Background(), &payment.TxReqType{
		PaymentClientEnum: 1,
		ReqIdClient:       reqIdclient.Int64(),
		ReceiverAdr:       conf.PaymentServer.PaymentSendAccount,
		Amount:            "0.000001",
		TokenNameEnum:     0})

	if err != nil {
		logrus.WithError(err).Error("Cannot get reply from payment service.")
		email.SendError("Payment", err.Error())
	}

	return nil
}