package cmd

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gitlab.com/MXCFoundation/cloud/conn_checker/internal/api"
	"gitlab.com/MXCFoundation/cloud/conn_checker/internal/config"
	"gitlab.com/MXCFoundation/cloud/conn_checker/internal/email"
	"gitlab.com/MXCFoundation/cloud/conn_checker/internal/log"
	"os"
	"os/signal"
	"syscall"
)

func run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tasks := []func() error{
		setLog,
		printStartMessage,
		setUpSMTP,
		setUpAPI,
	}

	for _, t := range tasks {
		if err := t(); err != nil {
			logrus.Fatal(err)
		}
	}

	sigChan := make(chan os.Signal)
	exitChan := make(chan struct{})
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	logrus.WithField("signal", <-sigChan).Info("signal received.")

	go func() {
		logrus.Warning("stopping conn-checker")
		exitChan <- struct{}{}
	}()

	select {
	case <-exitChan:
	case s := <-sigChan:
		logrus.WithField("signal", s).Info("signal received, stopping immediately.")
	}

	return nil
}

func setLog() error {
	if err := log.SetUp(); err != nil {
		return errors.Wrap(err, "setup Log error")
	}
	return nil
}

func printStartMessage() error {
	logrus.WithFields(logrus.Fields{
		"version": version,
	}).Info("Starting conn-checker ")
	return nil
}

func setUpSMTP() error {
	if err := email.Setup(config.Conf); err != nil {
		fmt.Println("setUpSMTP Error")
		return errors.Wrap(err, "setup SMTP error")
	}
	return nil
}

func setUpAPI() error {
	if err := api.Setup(config.Conf); err != nil {
		return errors.Wrap(err, "setup api error")
	}
	return nil
}