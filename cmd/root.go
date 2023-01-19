package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/AliRasoulinejad/cryptos-backend/internal/app"
	"github.com/AliRasoulinejad/cryptos-backend/internal/config"
	"github.com/AliRasoulinejad/cryptos-backend/internal/log"
)

var (
	// Flags vars
	configPath string

	rootCMD = &cobra.Command{
		Use:              "cryptos-backend",
		Short:            "a simple accountant for personal daily purchases",
		PersistentPreRun: preRun,
	}
)

func init() {
	cobra.OnInitialize(initialize)

	rootCMD.PersistentFlags().StringVarP(&configPath, "config", "c", "config.yml", "Path of config file")
	// Registering commands
	rootCMD.AddCommand(serveCMD)
}

func initialize() {
	fmt.Println(app.Banner())
}

func preRun(_ *cobra.Command, _ []string) {
	config.Init(configPath)
	log.InitLogger()
}

// Execute executes the root command.
func Execute() {
	err := rootCMD.Execute()
	if err != nil {
		log.Logger.Error(err)
	}
}
