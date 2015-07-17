package main

import (
	"net/http"
	"strings"

	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"

	"github.com/andlabs/ui"
)

type song struct {
	api    *songAPIType
	title  ui.Label
	album  ui.Label
	artist ui.Label
	year   ui.Label
	cover  *image.RGBA
	stack  ui.Control
}

// Updates the cover of the s song. Downloads the image, decodes it and draws it
// over the previous image.
func (s *song) generateCover(create bool) (err error) {
	var url string
	var src image.Image

	url = s.api.Visuel.Small
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	src, _, err = image.Decode(resp.Body)
	if create {
		s.cover = image.NewRGBA(src.Bounds())
	}
	draw.Draw(s.cover, s.cover.Rect, src, image.ZP, draw.Src)
	return
}

// Updates the labels for the s song.
// create defines whether or not we're creating the elements.
func (s *song) generateLabels(create bool) {
	if create {
		s.title = ui.NewLabel("Titre : " + strings.Title(strings.ToLower(s.api.Titre)))
		s.album = ui.NewLabel("Album : " + strings.Title(strings.ToLower(s.api.Titrealbum)))
		s.artist = ui.NewLabel("Artiste : " + strings.Title(strings.ToLower(s.api.Interpretemorceau)))
		s.year = ui.NewLabel("Année : " + s.api.Anneeeditionmusique)
	} else {
		s.title.SetText("Titre : " + strings.Title(strings.ToLower(s.api.Titre)))
		s.album.SetText("Album : " + strings.Title(strings.ToLower(s.api.Titrealbum)))
		s.artist.SetText("Artiste : " + strings.Title(strings.ToLower(s.api.Interpretemorceau)))
		s.year.SetText("Année : " + s.api.Anneeeditionmusique)
	}
}

// Creates the tab for the specified song. Calls updateLabels and updateCover
// with the create argument to true. Defines two grids : igrid and ogrid
// igrid contains all the labels
// ogrid groups the area containing the image handler and the igrid
func (s *song) createTab() {
	s.generateLabels(true)
	s.generateCover(true)
	// Inside Grid
	igrid := ui.NewGrid()
	igrid.Add(s.title, nil, ui.South, false, ui.LeftTop, false, ui.LeftTop, 1, 1)
	igrid.Add(s.album, nil, ui.South, false, ui.LeftTop, false, ui.LeftTop, 1, 1)
	igrid.Add(s.artist, nil, ui.South, false, ui.LeftTop, false, ui.LeftTop, 1, 1)
	igrid.Add(s.year, nil, ui.South, false, ui.LeftTop, false, ui.LeftTop, 1, 1)
	// Outside Grid
	ogrid := ui.NewGrid()
	ogrid.Add(ui.NewArea(100, 100, &areaHandler{s.cover}), nil, ui.South, false, ui.LeftTop, true, ui.LeftTop, 1, 1)
	ogrid.Add(igrid, nil, ui.East, false, ui.LeftTop, true, ui.LeftTop, 1, 1)
	s.stack = ui.NewVerticalStack(ogrid)
}

// Updates the s song's tab.
func (s *song) updateTab() {
	s.generateLabels(false)
	s.generateCover(false)
}

// Update all the tabs of all the songs passed as arguments.
func updateTabs(songs ...*song) {
	for _, c := range songs {
		go c.updateTab()
	}
}

// Creates all the tabs of all the songs passed as arguments.
func createTabs(songs ...*song) {
	for _, c := range songs {
		c.createTab()
	}
}
