package main

import (
	"clipnote/desktop/cmd/services"
	"clipnote/desktop/cmd/user"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"syscall"

	"github.com/spf13/cobra"
)

var background bool

func main() {

	var rootCmd = &cobra.Command{
		Use:   "clipnote",
		Short: "Clipnote CLI",
		Long:  `Clipnote CLI tool to copy`,
	}

	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(stopCmd)

	cobra.OnInitialize(func() {
		if os.Args[1] == "start" {
			initLogger()
		}
	})

	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
	}
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Clipnote",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("login cmd initiated")
		err := user.Login()
		if err != nil {
			log.Println(err)
		}

	},
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Clipnote service",
	Run: func(cmd *cobra.Command, args []string) {

		if background {
			services.RunTrayApp()
			return

		}

		exe, _ := os.Executable()
		proc := exec.Command(exe, "start", "--background")
		proc.SysProcAttr = &syscall.SysProcAttr{
			HideWindow:    true,
			CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP | 0x00000008,
		}
		err := proc.Start()
		if err != nil {
			log.Println("Failed to start background process", err)
		}

	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop Clipnote service",
	Run: func(cmd *cobra.Command, args []string) {
		services.StopClipnoteReading()
		log.Println("Clipnote sevice stopped")
	},
}

func initLogger() {

	var logPath = filepath.Join(`C:\ProgramData\clipnote`, "clipnote.log")

	os.MkdirAll(filepath.Dir(logPath), 0700) // create dir if not exists

	// Open log file in append mode
	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}

	// Redirect default logger to file
	log.SetOutput(f)

	// Add timestamp + short file name to log lines
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("Logger initialized")
}

func init() {
	// Called automatically when package services is imported

	// Register the --background flag for start command
	startCmd.Flags().BoolVar(&background, "background", false, "Run in background mode")
}
