package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andlabs/ui"
)

const (
	fipURL = "http://www.fipradio.fr/sites/default/files/import_si/si_titre_antenne/FIP_player_current.json"
)

var window ui.Window
var current fipAPIType
var currentSong song

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

func updateGui(c, p, n *ui.Control) {
	fetchLatest()

}

func create(s song, c songAPIType) {
	s.title = ui.NewLabel("Titre : " + c.Titre)
	s.album = ui.NewLabel("Album : " + current.Current.songAPIType.Titrealbum)
	s.artist = ui.NewLabel("Artiste : " + current.Current.songAPIType.Interpretemorceau)
	s.vstack = ui.NewVerticalStack(currentSong.title, currentSong.album, currentSong.artist)
}

func createCurrent() {
	currentSong.title = ui.NewLabel("Titre : " + current.Current.songAPIType.Titre)
	currentSong.album = ui.NewLabel("Album : " + current.Current.songAPIType.Titrealbum)
	currentSong.artist = ui.NewLabel("Artiste : " + current.Current.songAPIType.Interpretemorceau)
	currentSong.vstack = ui.NewVerticalStack(currentSong.title, currentSong.album, currentSong.artist)
}

func updateCurrent() {
	currentSong.title.SetText(current.Current.songAPIType.Titre)

}

func initGui() {
	createCurrent()
	tabstack := ui.NewTab()
	tabstack.Append("Current", currentSong.vstack)
	tabstack.Append("Previous", ui.Space())
	tabstack.Append("Next", ui.Space())
	tabstack.Append("Settings", ui.Space())
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
	// nf := notificator.New(notificator.Options{
	// 	DefaultIcon: "icon/default.png",
	// 	AppName:     "GoFip",
	// })
}
