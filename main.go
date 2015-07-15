package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/0xAX/notificator"
	"github.com/andlabs/ui"
)

const (
	fipURL = "http://www.fipradio.fr/sites/default/files/import_si/si_titre_antenne/FIP_player_current.json"
)

var window ui.Window
var current fipAPIType

func fetchLatest() {
	res, err := http.Get(fipURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&current)
	if err != nil {
		log.Fatal(err)
	}
}

func updateGui(ct closableTicker, c, p, p2, n, n2 *song, nt *notificator.Notificator, ntc ui.Checkbox) {
	var previous fipAPIType
	var title string
	var artist string
	var album string
	for {
		select {
		case <-ct.ticker.C:
			previous = current
			fetchLatest()
			if previous != current {
				title = current.Current.songAPIType.Titre
				artist = current.Current.songAPIType.Interpretemorceau
				album = current.Current.songAPIType.Titrealbum
				if ntc.Checked() {
					nt.Push("GoFIP", title+" par "+artist+" ("+album+")", "")
				}
				updateTab(c, current.Current.songAPIType)
				updateTab(p, current.Previous1.songAPIType)
				updateTab(p2, current.Previous2.songAPIType)
				updateTab(n, current.Next1.songAPIType)
				updateTab(n2, current.Next2.songAPIType)
			}
		case <-ct.halt:
			return
		}
	}
}

func createTab(s *song, c songAPIType) {
	s.title = ui.NewLabel("Titre : " + c.Titre)
	s.album = ui.NewLabel("Album : " + c.Titrealbum)
	s.artist = ui.NewLabel("Artiste : " + c.Interpretemorceau)
	s.year = ui.NewLabel("Année : " + c.Anneeeditionmusique)
	s.gstack = ui.NewGrid()
	s.gstack.Add(s.title, nil, ui.South, false, ui.LeftTop, false, ui.LeftTop, 1, 1)
	s.gstack.Add(s.album, nil, ui.South, false, ui.LeftTop, false, ui.LeftTop, 1, 1)
	s.gstack.Add(s.artist, nil, ui.South, false, ui.LeftTop, false, ui.LeftTop, 1, 1)
	s.gstack.Add(s.year, nil, ui.South, false, ui.LeftTop, false, ui.LeftTop, 1, 1)
	s.vstack = ui.NewVerticalStack(ui.Space(), s.gstack)
}

func updateTab(s *song, c songAPIType) {
	s.title.SetText("Titre : " + c.Titre)
	s.album.SetText("Album: " + c.Titrealbum)
	s.artist.SetText("Artiste : " + c.Interpretemorceau)
	s.year.SetText("Année : " + c.Anneeeditionmusique)
}

func initGui() {
	var currentSong song
	var previousSong song
	var previousSong2 song
	var nextSong song
	var nextSong2 song

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
	go updateGui(ct, &currentSong, &previousSong, &previousSong2, &nextSong, &nextSong2, nt, ntc)
	prc.OnToggled(func() {
		if !prc.Checked() {
			ct.stop()
		} else {
			ct = closableTicker{
				ticker: time.NewTicker(1 * time.Minute),
				halt:   make(chan bool, 1),
			}
			go updateGui(ct, &currentSong, &previousSong, &previousSong2, &nextSong, &nextSong2, nt, ntc)
		}
	})

	createTab(&currentSong, current.Current.songAPIType)
	createTab(&previousSong, current.Previous1.songAPIType)
	createTab(&previousSong2, current.Previous2.songAPIType)
	createTab(&nextSong, current.Next1.songAPIType)
	createTab(&nextSong2, current.Next2.songAPIType)
	tabstack := ui.NewTab()
	tabstack.Append("Current", currentSong.vstack)
	tabstack.Append("Previous", previousSong.vstack)
	tabstack.Append("Next", nextSong.vstack)
	tabstack.Append("Settings", ui.NewVerticalStack(ntc, prc))
	tabstack.Append("Credits", ui.NewLabel("Depado 2015"))

	window = ui.NewWindow("GoFIP", 400, 300, tabstack)
	window.OnClosing(func() bool {
		ui.Stop()
		return true
	})
	window.Show()
}

func main() {
	fetchLatest()
	go ui.Do(initGui)
	err := ui.Go()
	if err != nil {
		panic(err)
	}
}
