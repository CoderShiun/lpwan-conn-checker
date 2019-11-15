package appserver

import (
	"context"
	"github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/conn_checker/internal/config"
	"gitlab.com/MXCFoundation/cloud/conn_checker/internal/email"
	"gitlab.com/MXCFoundation/cloud/conn_checker/pkg/api/conn_checker"
	"google.golang.org/grpc"
)

func Setup() error {
	logrus.Info("Setup AppServer API")
	return nil
}

func connAppServer(address, port string) (*grpc.ClientConn, error) {
	logrus.WithFields(logrus.Fields{
		"address": address,
		"port": port,
	}).Debug("appServer/connAppServer")

	conn, err := grpc.Dial(address + port, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func AppServerConnTest(conf *config.Config) error {
	conn, err := connAppServer(conf.AppServer.AppServiceAddress, conf.AppServer.AppServicePort)
	if err != nil {
		logrus.WithError(err).Error("Cannot conn to Lora App Server.")
		email.SendError("AppServer", err.Error())
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			logrus.WithError(err).Error("Cannot close App connection.")
		}
	}()

	client := conn_checker.NewConnCheckerServerServiceClient(conn)

	_, err = client.CheckConnection(context.Background(), nil)
	if err != nil {
		logrus.WithError(err).Error("Cannot get reply from App Server.")
		email.SendError("AppServer", err.Error())
	}

	return nil
}