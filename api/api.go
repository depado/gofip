package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/0xAX/notificator"
)

// URL is the URL of the JSON API of fip. Exported as api.URL
const URL = "http://www.fipradio.fr/sites/default/files/import_si/si_titre_antenne/FIP_player_current.json"

// FIP is the struct representing the response of the API URL
type FIP struct {
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
		Song `json:"song"`
	} `json:"current"`
	Previous2 struct {
		Song `json:"song"`
	} `json:"previous2"`
	Previous1 struct {
		Song `json:"song"`
	} `json:"previous1"`
	Next1 struct {
		Song `json:"song"`
	} `json:"next1"`
	Next2 struct {
		Song `json:"song"`
	} `json:"next2"`
}

// Update updates an instance of FIP to the latest available data
func (f *FIP) Update() error {
	res, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(f)
	return err
}

// Song is a shorthand for data contained in FIP. Representing a single song.
type Song struct {
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

// Notify pushes a notification on the desktop of the user.
func (s *Song) Notify(nt *notificator.Notificator) {
	title := strings.Title(strings.ToLower(s.Titre))
	artist := strings.Title(strings.ToLower(s.Interpretemorceau))
	album := strings.Title(strings.ToLower(s.Titrealbum))
	nt.Push(title, fmt.Sprintf("%s (%s)", artist, album), "")
}
