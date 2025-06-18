package main

import (
	"fmt"
	"os"

	"github.com/ocrosby/go-load-tester/internal/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "loadtester",
	Short: "A CLI tool for load testing",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to loadtester!")
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
