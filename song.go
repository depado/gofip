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
	gstack ui.Grid
	api    *songAPIType
}

type cover struct {
	icon ui.ImageIndex
}

type areaHandler struct {
	img *image.RGBA
}

func (a *areaHandler) Paint(rect image.Rectangle) *image.RGBA {
	return a.img.SubImage(rect).(*image.RGBA)
}

func (a *areaHandler) Mouse(me ui.MouseEvent)  {}
func (a *areaHandler) Key(ke ui.KeyEvent) bool { return false }

func (s *song) dlConvertCover() {
	var url string
	var err error
	var src image.Image
	var img *image.RGBA

	url = s.api.Visuel.Small
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	src, _, err = image.Decode(resp.Body)
	img = image.NewRGBA(src.Bounds())
	draw.Draw(img, img.Rect, src, image.ZP, draw.Src)
	s.cover = img
}

func (s *song) updateCover() {
	s.dlConvertCover()
	s.covera = ui.NewArea(100, 100, &areaHandler{s.cover})
}

func (s *song) createTab() {
	s.title = ui.NewLabel("Titre : " + strings.Title(strings.ToLower(s.api.Titre)))
	s.album = ui.NewLabel("Album : " + strings.Title(strings.ToLower(s.api.Titrealbum)))
	s.artist = ui.NewLabel("Artiste : " + strings.Title(strings.ToLower(s.api.Interpretemorceau)))
	s.year = ui.NewLabel("Année : " + s.api.Anneeeditionmusique)
	s.updateCover()
	ggrid := ui.NewSimpleGrid(1, s.title, s.album, s.artist, s.year)
	s.gstack = ui.NewGrid()
	s.gstack.Add(s.covera, nil, ui.South, false, ui.LeftTop, true, ui.LeftTop, 1, 1)
	s.gstack.Add(ggrid, nil, ui.East, false, ui.LeftTop, false, ui.LeftTop, 1, 1)
	// s.gstack.Add(s.album, nil, ui.South, false, ui.LeftTop, false, ui.LeftTop, 1, 1)
	// s.gstack.Add(s.artist, nil, ui.South, false, ui.LeftTop, false, ui.LeftTop, 1, 1)
	// s.gstack.Add(s.year, nil, ui.South, false, ui.LeftTop, false, ui.LeftTop, 1, 1)
	s.vstack = ui.NewVerticalStack(s.gstack)
}

func (s *song) updateTab() {
	s.title.SetText("Titre : " + strings.Title(strings.ToLower(s.api.Titre)))
	s.album.SetText("Album: " + strings.Title(strings.ToLower(s.api.Titrealbum)))
	s.artist.SetText("Artiste : " + strings.Title(strings.ToLower(s.api.Interpretemorceau)))
	s.year.SetText("Année : " + s.api.Anneeeditionmusique)
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
