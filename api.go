package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/0xAX/notificator"
)

type fipAPIType struct {
	Current struct {
		Emission struct {
			Starttime int    `json:"startTime"`
			Endtime   int    `json:"endTime"`
			ID        string `json:"id"`
			Titre     string `json:"titre"`
			Visuel    struct {
				Small string `json:"small"`
			} `json:"visuel"`
			Lien string `json:"lien"`
		} `json:"emission"`
		songAPIType `json:"song"`
	} `json:"current"`
	Previous2 struct {
		songAPIType `json:"song"`
	} `json:"previous2"`
	Previous1 struct {
		songAPIType `json:"song"`
	} `json:"previous1"`
	Next1 struct {
		songAPIType `json:"song"`
	} `json:"next1"`
	Next2 struct {
		songAPIType `json:"song"`
	} `json:"next2"`
}

type songAPIType struct {
	Starttime           int    `json:"startTime"`
	Endtime             int    `json:"endTime"`
	ID                  string `json:"id"`
	Titre               string `json:"titre"`
	Titrealbum          string `json:"titreAlbum"`
	Interpretemorceau   string `json:"interpreteMorceau"`
	Anneeeditionmusique string `json:"anneeEditionMusique"`
	Label               string `json:"label"`
	Visuel              struct {
		Small  string `json:"small"`
		Medium string `json:"medium"`
	} `json:"visuel"`
	Lien string `json:"lien"`
}

func (s *songAPIType) push(nt *notificator.Notificator) {
	title := strings.Title(strings.ToLower(s.Titre))
	artist := strings.Title(strings.ToLower(s.Interpretemorceau))
	album := strings.Title(strings.ToLower(s.Titrealbum))
	nt.Push(title, artist+" ("+album+")", "")
}

func fetchLatest() (current fipAPIType, err error) {
	res, err := http.Get(fipURL)
	if err != nil {
		return
	}
	defer res.Body.Close()
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&current)
	if err != nil {
		return
	}
	return
}
