/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"github.com/pkg/errors"
	"gitlab.com/MXCFoundation/cloud/conn_checker/internal/config"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

const configTemplate = `[general]
# Log level
#
# debug=5, info=4, warning=3, error=2, fatal=1, panic=0
log_level={{ .General.LogLevel }}

# The number of times passwords must be hashed. A higher number is safer as
# an attack takes more time to perform.
password_hash_iterations={{ .General.PasswordHashIterations }}
`

// configCmd represents the configuration command
var configCmd = &cobra.Command{
	Use:   "configfile",
	Short: "Print the Conn Checker configuration file",
	Long: ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		t := template.Must(template.New("config").Parse(configTemplate))
		err := t.Execute(os.Stdout, config.Conf)
		if err != nil {
			return errors.Wrap(err, "execute config template error")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
