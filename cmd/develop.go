package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/axarus/vectrag/internal/application"
	infrahttp "github.com/axarus/vectrag/internal/infrastructure/http"
	"github.com/spf13/cobra"
)

var developCmd = &cobra.Command{
	Use:   "develop",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		basePort := 51987
		host := "localhost"

		svc := application.NewDevelopService(
			infrahttp.ListenerProvider{},
			infrahttp.ServerStarter{},
			infrahttp.AdminHandlerProvider{},
		)

		url, shutdown, err := svc.Start(basePort, host)
		if err != nil {
			fmt.Println("Error starting server:", err)
			return
		}
		fmt.Println("Server started at", url)

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		fmt.Println("Shutting down server...")
		_ = shutdown(context.Background())
	},
}

func init() {
	rootCmd.AddCommand(developCmd)
}
