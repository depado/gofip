package main

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
		Song struct {
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
		} `json:"song"`
	} `json:"previous2"`
	Previous1 struct {
		Song struct {
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
		} `json:"song"`
	} `json:"previous1"`
	Next1 struct {
		Song struct {
			Starttime           int    `json:"startTime"`
			Endtime             int    `json:"endTime"`
			ID                  string `json:"id"`
			Titre               string `json:"titre"`
			Titrealbum          string `json:"titreAlbum"`
			Interpretemorceau   string `json:"interpreteMorceau"`
			Anneeeditionmusique string `json:"anneeEditionMusique"`
			Visuel              struct {
				Small  string `json:"small"`
				Medium string `json:"medium"`
			} `json:"visuel"`
			Lien string `json:"lien"`
		} `json:"song"`
	} `json:"next1"`
	Next2 struct {
		Song struct {
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
		} `json:"song"`
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
