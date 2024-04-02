package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/shosato0306/golang-ddd-boilerplate/internal/app/api"
)

func main() {
	var rootCmd = &cobra.Command{Use: "myapp"}

	var apiCmd = &cobra.Command{
		Use:   "api-server",
		Short: "Start the API server",
		Run: func(_ *cobra.Command, _ []string) {
			api.Run()
		},
	}

	var batchCmd = &cobra.Command{
		Use:   "batch",
		Short: "Run the batch job",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println("Running batch job...")
		},
	}

	rootCmd.AddCommand(apiCmd)
	rootCmd.AddCommand(batchCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
