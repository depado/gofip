package main

import (
	"log"
	"net/http"
	"strings"

	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"

	"github.com/andlabs/ui"
)

type song struct {
	title  ui.Label
	album  ui.Label
	artist ui.Label
	year   ui.Label
	cover  *image.RGBA
	covera ui.Area
	vstack ui.Control
	api    *songAPIType
}

type areaHandler struct {
	img *image.RGBA
}

func (a *areaHandler) Paint(rect image.Rectangle) *image.RGBA {
	return a.img.SubImage(rect).(*image.RGBA)
}

func (a *areaHandler) Mouse(me ui.MouseEvent)  {}
func (a *areaHandler) Key(ke ui.KeyEvent) bool { return false }

func (s *song) updateCover() {
	var url string
	var err error
	var src image.Image

	url = s.api.Visuel.Small
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	src, _, err = image.Decode(resp.Body)
	s.cover = image.NewRGBA(src.Bounds())
	draw.Draw(s.cover, s.cover.Rect, src, image.ZP, draw.Src)
	s.covera = ui.NewArea(100, 100, &areaHandler{s.cover})
}

func (s *song) updateLabels() {
	s.title = ui.NewLabel("Titre : " + strings.Title(strings.ToLower(s.api.Titre)))
	s.album = ui.NewLabel("Album : " + strings.Title(strings.ToLower(s.api.Titrealbum)))
	s.artist = ui.NewLabel("Artiste : " + strings.Title(strings.ToLower(s.api.Interpretemorceau)))
	s.year = ui.NewLabel("Année : " + s.api.Anneeeditionmusique)
}

func (s *song) createTab() {
	s.updateLabels()
	s.updateCover()
	// Inside Grid
	igrid := ui.NewGrid()
	igrid.Add(s.title, nil, ui.South, false, ui.LeftTop, false, ui.LeftTop, 1, 1)
	igrid.Add(s.album, nil, ui.South, false, ui.LeftTop, false, ui.LeftTop, 1, 1)
	igrid.Add(s.artist, nil, ui.South, false, ui.LeftTop, false, ui.LeftTop, 1, 1)
	igrid.Add(s.year, nil, ui.South, false, ui.LeftTop, false, ui.LeftTop, 1, 1)
	// Outside Grid
	ogrid := ui.NewGrid()
	ogrid.Add(s.covera, nil, ui.South, false, ui.LeftTop, true, ui.LeftTop, 1, 1)
	ogrid.Add(igrid, nil, ui.East, false, ui.LeftTop, true, ui.LeftTop, 1, 1)
	s.vstack = ui.NewVerticalStack(ogrid)
}

func (s *song) updateTab() {
	s.updateLabels()
	s.updateCover()
}

func updateTabs(songs ...*song) {
	for _, c := range songs {
		c.updateTab()
	}
}

func createTabs(songs ...*song) {
	for _, c := range songs {
		c.createTab()
	}
}
