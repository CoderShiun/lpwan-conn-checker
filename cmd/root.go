/*
Copyright © 2019 MXC_Foundation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
  "bytes"
  "github.com/spf13/cobra"
  "gitlab.com/MXCFoundation/cloud/conn_checker/internal/config"
  "io/ioutil"
  "reflect"
  "strings"

  "github.com/spf13/viper"

  "github.com/sirupsen/logrus"
)

var cfgFile string
var version string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "connchecker",
	Short: "A brief description of your application",
	Long:  ``,
	RunE:  run,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(v string) {
	version = v
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
		//os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "configuration", "", "configuration file (default is $configuration.toml)")

	rootCmd.PersistentFlags().Int("log-level", 4, "debug=5, info=4, error=2, fatal=1, panic=0")

	// bind flag to configuration vars
	err := viper.BindPFlag("general.log_level", rootCmd.PersistentFlags().Lookup("log-level"))
	if err != nil {
		logrus.WithError(err).Error("cannot bind log level")
	}

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in configuration file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		b, err := ioutil.ReadFile(cfgFile)
		if err != nil {
			logrus.WithError(err).WithField("configuration", cfgFile).Fatal("error loading configuration file")
		}
		//viper.SetConfigFile(cfgFile)
		viper.SetConfigType("toml")
		if err := viper.ReadConfig(bytes.NewBuffer(b)); err != nil {
			logrus.WithError(err).WithField("configuration", cfgFile).Fatal("error loading configuration file")
		}
	} else {
		viper.SetConfigName("conn-checker")
		// Search configuration in fallowing paths
		viper.AddConfigPath(".")
		viper.AddConfigPath("../configuration/conn-checker")
		viper.AddConfigPath("$HOME/.configuration/conn-checker")
		viper.AddConfigPath("/etc/conn-checker")

		if err := viper.ReadInConfig(); err != nil {
			switch err.(type) {
			case viper.ConfigFileNotFoundError:
				logrus.Warning("No configuration file found, using defaults.")
			default:
				logrus.WithError(err).Fatal("read configuration file error")
			}
		}
	}

	viperBindEnvs(config.Conf)

	if err := viper.Unmarshal(&config.Conf); err != nil {
		logrus.WithError(err).Fatal("unmarshal config error")
	}

	/*if config.Conf.ConnCheckerServer.Integration.Backend != "" {
	  config.Conf.ConnCheckerServer.Integration.Enabled = []string{config.C.ApplicationServer.Integration.Backend}
	}*/

	config.ConnCheckerVersion = version
}

func viperBindEnvs(iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)

	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			tv = strings.ToLower(t.Name)
		}
		if tv == "-" {
			continue
		}

		switch v.Kind() {
		case reflect.Struct:
			viperBindEnvs(v.Interface(), append(parts, tv)...)
		default:
			key := strings.Join(append(parts, tv), ".")
			err := viper.BindEnv(key)
			if err != nil {
			  logrus.WithError(err).Error("config format error")
            }
		}
	}
}
