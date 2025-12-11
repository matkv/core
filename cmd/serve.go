package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/matkv/core/internal/config"
	"github.com/matkv/core/internal/server"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the web server",
	RunE: func(cmd *cobra.Command, args []string) error {
		port := 8080

		// start server in a goroutine
		go func() {
			if err := server.Start(port); err != nil {
				fmt.Fprintf(os.Stderr, "server error: %v\n", err)
			}
		}()

		// open browser to the server URL
		if config.C.Device != config.WSL {
			url := fmt.Sprintf("http://localhost:%d", port)
			if err := OpenURL(url); err != nil {
				fmt.Fprintf(os.Stderr, "failed to open browser: %v\n", err)
			}
		}

		// wait for user interrupt to shut down (Ctrl+C)
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		<-sigCh

		return nil
	},
}
