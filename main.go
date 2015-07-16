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

func updateGui(ct closableTicker, cs, ps, ps2, ns, ns2 *song, nt *notificator.Notificator, ntc ui.Checkbox) {
	var err error
	var previous fipAPIType
	var title string
	var artist string
	var album string
	for {
		select {
		case <-ct.ticker.C:
			previous = current
			current, err = fetchLatest()
			if err != nil {
				log.Fatal(err)
			}
			if previous != current {
				if ntc.Checked() {
					title = current.Current.songAPIType.Titre
					artist = current.Current.songAPIType.Interpretemorceau
					album = current.Current.songAPIType.Titrealbum
					nt.Push(title, artist+" ("+album+")", "")
				}
				cs.updateTab(current.Current.songAPIType)
				ps.updateTab(current.Previous1.songAPIType)
				ps2.updateTab(current.Previous2.songAPIType)
				ns.updateTab(current.Next1.songAPIType)
				ns2.updateTab(current.Next2.songAPIType)
			}
		case <-ct.halt:
			return
		}
	}
}

func initPlayer() (player *gst.Element) {
	player = gst.ElementFactoryMake("playbin", "player")
	player.SetProperty("uri", fipStreamURL)
	player.SetState(gst.STATE_PLAYING)
	return
}

func initGui() {
	player := initPlayer()
	playing := true

	var cs, ps, ps2, ns, ns2 song

	nt := notificator.New(notificator.Options{
		DefaultIcon: "icon/default.png",
		AppName:     "GoFip",
	})

	ntc := ui.NewCheckbox("Notifications")
	ntc.SetChecked(true)
	prc := ui.NewCheckbox("Periodic Check")
	prc.SetChecked(true)
	ct := closableTicker{
		ticker: time.NewTicker(1 * time.Minute),
		halt:   make(chan bool, 1),
	}
	go updateGui(ct, &cs, &ps, &ps2, &ns, &ps, nt, ntc)
	prc.OnToggled(func() {
		if !prc.Checked() {
			ct.stop()
		} else {
			ct = closableTicker{
				ticker: time.NewTicker(1 * time.Minute),
				halt:   make(chan bool, 1),
			}
			go updateGui(ct, &cs, &ps, &ps2, &ns, &ns2, nt, ntc)
		}
	})

	cs.createTab(current.Current.songAPIType)
	ps.createTab(current.Previous1.songAPIType)
	ps2.createTab(current.Previous2.songAPIType)
	ns.createTab(current.Next1.songAPIType)
	ns2.createTab(current.Next2.songAPIType)
	ts := ui.NewTab()
	ts.Append("Current", cs.vstack)
	ts.Append("Previous", ps.vstack)
	ts.Append("Next", ns.vstack)
	ts.Append("Settings", ui.NewVerticalStack(ntc, prc))
	ts.Append("Credits", ui.NewLabel("Depado 2015"))

	psl := ui.NewLabel("Currently Playing")
	ppbtn := ui.NewButton("Pause")
	ppbtn.OnClicked(func() {
		if playing {
			ppbtn.SetText("Play")
			player.SetState(gst.STATE_PAUSED)
			psl.SetText("Currently Paused")
			playing = false
		} else {
			ppbtn.SetText("Pause")
			psl.SetText("Currently Playing")
			player.SetState(gst.STATE_PLAYING)
			playing = true
		}
	})
	mvs := ui.NewVerticalStack(ts, ppbtn, psl)
	mvs.SetStretchy(0)

	window = ui.NewWindow("GoFIP", 400, 200, mvs)
	window.OnClosing(func() bool {
		ui.Stop()
		return true
	})
	window.Show()
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
