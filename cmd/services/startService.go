package services

import (
	"log"
	"sync/atomic"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/go-vgo/robotgo/clipboard"
	hook "github.com/robotn/gohook"
)

var (
	StopChan   chan struct{}
	Monitoring atomic.Bool
)

// StartClipnoteReading starts clipboard monitoring
func StartClipnoteReading() {
	if Monitoring.Load() {
		log.Println("Already monitoring")
		return
	}

	log.Println("Clipnote monitoring started")
	log.Println("Select text anywhere, then press Ctrl + Q")

	StopChan = make(chan struct{})
	Monitoring.Store(true)
	UpdateTooltip()

	evChan := hook.Start()

	go func() {
		defer func() {
			hook.End()
			Monitoring.Store(false)
			UpdateTooltip()
			log.Println("Clipboard monitor stopped")
		}()

		ctrlPress := false

		for {
			select {
			case ev := <-evChan:
				switch ev.Kind {
				case hook.KeyDown:
					// 162 = Ctrl (Windows), consider using ev.Keychar for portability
					if ev.Rawcode == 162 {
						ctrlPress = true
						log.Println("ctrlpress true")
					}

					// 81 = Q
					if ev.Rawcode == 81 && ctrlPress {
						log.Println("ctrl + q detected")

						// Copy selected text
						robotgo.KeyTap("c", "ctrl")
						log.Println("text copied")

						go func() {
							time.Sleep(1500 * time.Millisecond)

							text, err := clipboard.ReadAll()
							if err != nil {
								log.Println("Clipboard read error:", err)
								return
							}
							if text == "" {
								log.Println("No text found, nothing was selected")
								return
							}
							log.Println("Captured:", text)
							SendServer(text)
						}()
					}

				case hook.KeyUp:
					if ev.Rawcode == 162 {
						ctrlPress = false
						log.Println("ctrlpress false")
					}
				}

			case <-StopChan:
				Monitoring.Store(false)
				return
			}
		}
	}()
}

// StopClipnoteReading stops monitoring safely
func StopClipnoteReading() {
	if Monitoring.Load() {
		StopChan <- struct{}{}
	}
}
