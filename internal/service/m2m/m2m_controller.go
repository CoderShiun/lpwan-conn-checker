package m2m

import (
	"context"
	"github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/conn_checker/internal/config"
	"gitlab.com/MXCFoundation/cloud/conn_checker/internal/email"
	"gitlab.com/MXCFoundation/cloud/conn_checker/pkg/api/conn_checker"
	"google.golang.org/grpc"
)

func Setup() error {
	logrus.Info("Setup M2M API")
	return nil
}

func connM2MServer(address, port string) (*grpc.ClientConn, error) {
	logrus.WithFields(logrus.Fields{
		"address": address,
		"port": port,
	}).Debug("m2mServer/connM2MServer")

	conn, err := grpc.Dial(address + port, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func M2MServerConnTest(conf *config.Config) error {
	conn, err := connM2MServer(conf.M2MServer.M2MServiceAddress, conf.M2MServer.M2MServicePort)
	if err != nil {
		logrus.WithError(err).Error("Cannot conn to M2M Server.")
		email.SendError("M2MServer", err.Error())
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			logrus.WithError(err).Error("Cannot close M2M connection.")
		}
	}()

	client := conn_checker.NewConnCheckerServerServiceClient(conn)

	_, err = client.CheckConnection(context.Background(), nil)
	if err != nil {
		logrus.WithError(err).Error("Cannot get reply from M2M Server.")
		email.SendError("M2MServer", err.Error())
	}

	return nil
}