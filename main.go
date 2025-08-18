package main

import (
	"clipnote/desktop/cmd/user"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func main() {

	var rootCmd = &cobra.Command{
		Use:   "clipnote",
		Short: "Clipnote CLI",
		Long:  `Clipnote CLI tool to copy`,
	}

	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(stopCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
	}
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Clipnote",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login cmd initiated")
		err := user.Login()
		if err != nil {
			fmt.Println(err)
		}

	},
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Clipnote service",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("clipnote service started")

		fmt.Println("clipnote service started in background")
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop Clipnote service",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Clipnote sevice stopped")
	},
}
