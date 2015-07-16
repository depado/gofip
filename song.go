package main

import (
	"strings"

	"github.com/andlabs/ui"
)

type song struct {
	title  ui.Label
	album  ui.Label
	artist ui.Label
	year   ui.Label
	vstack ui.Control
	gstack ui.Grid
	api    *songAPIType
}

func (s *song) createTab() {
	s.title = ui.NewLabel("Titre : " + strings.Title(strings.ToLower(s.api.Titre)))
	s.album = ui.NewLabel("Album : " + strings.Title(strings.ToLower(s.api.Titrealbum)))
	s.artist = ui.NewLabel("Artiste : " + strings.Title(strings.ToLower(s.api.Interpretemorceau)))
	s.year = ui.NewLabel("Année : " + s.api.Anneeeditionmusique)
	s.gstack = ui.NewGrid()
	s.gstack.Add(s.title, nil, ui.South, false, ui.LeftTop, false, ui.LeftTop, 1, 1)
	s.gstack.Add(s.album, nil, ui.South, false, ui.LeftTop, false, ui.LeftTop, 1, 1)
	s.gstack.Add(s.artist, nil, ui.South, false, ui.LeftTop, false, ui.LeftTop, 1, 1)
	s.gstack.Add(s.year, nil, ui.South, false, ui.LeftTop, false, ui.LeftTop, 1, 1)
	s.vstack = ui.NewVerticalStack(ui.Space(), s.gstack)
}

func (s *song) updateTab() {
	s.title.SetText("Titre : " + strings.Title(strings.ToLower(s.api.Titre)))
	s.album.SetText("Album: " + strings.Title(strings.ToLower(s.api.Titrealbum)))
	s.artist.SetText("Artiste : " + strings.Title(strings.ToLower(s.api.Interpretemorceau)))
	s.year.SetText("Année : " + s.api.Anneeeditionmusique)
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
