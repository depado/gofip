package main

import (
	"log"
	"time"

	"github.com/0xAX/notificator"
	"github.com/andlabs/ui"
	"github.com/ziutek/gst"
)

const (
	fipURL       = "http://www.fipradio.fr/sites/default/files/import_si/si_titre_antenne/FIP_player_current.json"
	fipStreamURL = "http://audio.scdn.arkena.com/11016/fip-midfi128.mp3"
)

var window ui.Window
var current fipAPIType

// Updates the GUI according to the closableTicker ct, and the songs passed as arguments.
func updateGui(ct closableTicker, nt *notificator.Notificator, ntc ui.Checkbox, songs ...*song) {
	var err error
	var previous fipAPIType
	for {
		select {
		case <-ct.ticker.C:
			previous = current
			current, err = fetchLatest()
			if err != nil {
				log.Fatal(err)
			}
			// In case the current song changed, push a notification if the ntc
			// checkbox is checked and update the tabs.
			if previous != current {
				if ntc.Checked() {
					current.Current.songAPIType.push(nt)
				}
				updateTabs(songs...)
			}
		case <-ct.halt:
			// When the closableTicker stops, end the goroutine.
			return
		}
	}
}

// Initiliazes the GST Player to read the stream.
func initPlayer() (player *gst.Element) {
	player = gst.ElementFactoryMake("playbin", "player")
	player.SetProperty("uri", fipStreamURL)
	return
}

func initGui() {
	// Creates the initial songs that will be used as long as the program runs.
	cs := song{api: &current.Current.songAPIType}
	ps := song{api: &current.Previous1.songAPIType}
	// ps2 := song{api: &current.Previous2.songAPIType}
	ns := song{api: &current.Next1.songAPIType}
	// ns2 := song{api: &current.Next2.songAPIType}

	// Creates the player and the controls for the player (as well as label).
	player := initPlayer()
	playing := true
	psl := ui.NewLabel("Currently Playing")
	ppbtn := ui.NewButton("Pause")
	ppbtn.OnClicked(func() {
		if playing {
			ppbtn.SetText("Play")
			player.SetState(gst.STATE_PAUSED)
			psl.SetText("Currently Paused")
		} else {
			ppbtn.SetText("Pause")
			psl.SetText("Currently Playing")
			player.SetState(gst.STATE_PLAYING)
		}
		playing = !playing
	})

	// Creates the notification system
	nt := notificator.New(notificator.Options{
		DefaultIcon: "icon/default.png",
		AppName:     "GoFip",
	})

	// Notification settings. Will be passed to the updateGui goroutine to Check
	// whether or not to send a system notification when the music changes.
	ntc := ui.NewCheckbox("Notifications")
	ntc.SetChecked(true)
	// Defines a closableTicker that is used in the updateGui goroutine
	// This allows to close the ticker when the setting button is unchecked.
	ct := closableTicker{
		ticker: time.NewTicker(1 * time.Minute),
		halt:   make(chan bool, 1),
	}
	prc := ui.NewCheckbox("Periodic Check")
	prc.SetChecked(true)
	prc.OnToggled(func() {
		if !prc.Checked() {
			ct.stop()
		} else {
			ct = closableTicker{
				ticker: time.NewTicker(1 * time.Minute),
				halt:   make(chan bool, 1),
			}
			go updateGui(ct, nt, ntc, &cs, &ps, &ns)
		}
	})

	// Start the goroutine to update the GUI every minute (default behaviour)
	// Uses the closableTicker defined earlier so the goroutine can be stopped.
	go updateGui(ct, nt, ntc, &cs, &ps, &ns)

	// Creating the tabs with the songs as well as settings and credits.
	createTabs(&cs, &ps, &ns)
	ts := ui.NewTab()
	ts.Append("Current", cs.vstack)
	ts.Append("Previous", ps.vstack)
	ts.Append("Next", ns.vstack)
	ts.Append("Settings", ui.NewVerticalStack(ntc, prc))
	ts.Append("Credits", ui.NewLabel("Depado 2015"))

	// Creates the main vertical stack that is passed to the main window.
	mvs := ui.NewVerticalStack(ts, ppbtn, psl)
	// The tab control must be set to stretchy otherwise it won't display the content
	mvs.SetStretchy(0)

	// Creates the main window and the behaviour on close event.
	window = ui.NewWindow("GoFIP", 400, 200, mvs)
	window.OnClosing(func() bool {
		ui.Stop()
		return true
	})
	window.Show()
	player.SetState(gst.STATE_PLAYING)
}

func main() {
	var err error
	current, err = fetchLatest()
	if err != nil {
		log.Fatal(err)
	}
	go ui.Do(initGui)
	err = ui.Go()
	if err != nil {
		panic(err)
	}
}
