package services

import (
	"fmt"
	"log"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/go-vgo/robotgo/clipboard"
	hook "github.com/robotn/gohook"
)

func StartClipnoteReading() {

	log.Println("Clipboard monitor running")
	log.Println("Select text anywhere, then press Ctrl + Q")

	evChan := hook.Start()
	defer hook.End()

	ctrlPress := false

	for ev := range evChan {

		switch ev.Kind {

		case hook.KeyDown:
			if ev.Rawcode == 162 {
				ctrlPress = true
				log.Println("ctrlpress true")

			}

			if ev.Rawcode == 81 && ctrlPress {
				log.Println("ctrl + q detected")

				robotgo.KeyTap("c", "ctrl")
				fmt.Println("text copied")

				go func() {
					time.Sleep(1500 * time.Millisecond)

					text, err := clipboard.ReadAll()
					if err != nil {
						fmt.Println("Clipboard read error", err)
					}
					if text == "" {
						fmt.Println("No text found nothing was selected")
						return
					}
					log.Println(text)
					SendServer(text)
				}()

			}

		case hook.KeyUp:
			if ev.Rawcode == 162 {
				ctrlPress = false
				fmt.Println("ctrlpress false")
			}

		}

	}

}
