package services

import (
	"log"
	"os"

	"github.com/getlantern/systray"
)

func RunTrayApp() {
	systray.Run(onReady, onExit)
}

func UpdateTooltip() {
	if Monitoring.Load() {
		systray.SetTooltip("Running")
	} else {
		systray.SetTooltip("Not Running")
	}
}

func onReady() {

	systray.SetTitle("Clipnote")

	iconPath := `D:\Learning\go_project\Clipnote\clipnote_desktop\c.ico`
	if iconData, err := os.ReadFile(iconPath); err == nil {
		systray.SetIcon(iconData)
	} else {
		log.Println("failed to load tray icon:", err)
	}

	mStart := systray.AddMenuItem("Start Clinpnote monitor", "Start monitoring clipboard")
	mStop := systray.AddMenuItem("Stop Clinpnote monitor", "Stop monitoring clipboard")
	mQuit := systray.AddMenuItem("Quit", "Exit the app")

	if !Monitoring.Load() {
		go StartClipnoteReading()
	}

	UpdateTooltip()

	go func() {

		for {
			select {
			case <-mStart.ClickedCh:
				if !Monitoring.Load() {
					go StartClipnoteReading()

				}

			case <-mStop.ClickedCh:
				if Monitoring.Load() {
					StopClipnoteReading()
				}

			case <-mQuit.ClickedCh:
				systray.Quit()
				StopClipnoteReading()
				return

			}
		}
	}()

}

func onExit() {
	log.Println("Clipnote tray exiting")
}
