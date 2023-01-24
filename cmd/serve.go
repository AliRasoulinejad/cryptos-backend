package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/AliRasoulinejad/cryptos-backend/internal/app"
	internalHttp "github.com/AliRasoulinejad/cryptos-backend/internal/http"
)

var serveCMD = &cobra.Command{
	Use:   "serve",
	Short: "Serve the application",
	Long:  `Serve the HTTP server of the application`,
	Run:   serve,
}

func serve(_ *cobra.Command, _ []string) {
	shutdownRequest := make(chan struct{})
	application := &app.Application{}
	application.WithTracer()
	application.WithDB()
	application.WithRepositories()

	shutdownReady := internalHttp.
		NewServer().
		Serve(application).
		WaitForSignals(shutdownRequest)

	quit := make(chan os.Signal, 2)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	shutdownRequest <- struct{}{}
	<-shutdownReady
}
