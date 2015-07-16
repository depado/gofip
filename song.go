package main

import "github.com/andlabs/ui"

type song struct {
	title  ui.Label
	album  ui.Label
	artist ui.Label
	year   ui.Label
	vstack ui.Control
	gstack ui.Grid
}

func (s *song) createTab(c songAPIType) {
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

func (s *song) updateTab(c songAPIType) {
	s.title.SetText("Titre : " + c.Titre)
	s.album.SetText("Album: " + c.Titrealbum)
	s.artist.SetText("Artiste : " + c.Interpretemorceau)
	s.year.SetText("Année : " + c.Anneeeditionmusique)
}
