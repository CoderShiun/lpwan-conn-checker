package networkserver

import (
	"context"
	"github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/conn_checker/internal/config"
	"gitlab.com/MXCFoundation/cloud/conn_checker/internal/email"
	"gitlab.com/MXCFoundation/cloud/conn_checker/pkg/api/conn_checker"
	"google.golang.org/grpc"
)

func Setup() error {
	logrus.Info("Setup NetworkServer API")
	return nil
}

func connNsServer(address, port string) (*grpc.ClientConn, error) {
	logrus.WithFields(logrus.Fields{
		"address": address,
		"port": port,
	}).Debug("nsServer/connNetworkServer")

	conn, err := grpc.Dial(address + port, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NsServerConnTest(conf *config.Config) error {
	conn, err := connNsServer(conf.NsServer.NsServiceAddress, conf.NsServer.NsServicePort)
	if err != nil {
		logrus.WithError(err).Error("Cannot conn to NetworkServer.")
		email.SendError("NetworkServer", err.Error())
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			logrus.WithError(err).Error("Cannot close NetworkServer connection.")
		}
	}()

	client := conn_checker.NewConnCheckerServerServiceClient(conn)

	_, err = client.CheckConnection(context.Background(), nil)
	if err != nil {
		logrus.WithError(err).Error("Cannot get reply from NetworkServer.")
		email.SendError("NetworkServer", err.Error())
	}

	return nil
}