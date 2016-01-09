package main

import (
	"log"

	"github.com/Depado/gofip/api"
	"github.com/Depado/gofip/player"
)

var current api.FIP

func main() {
	var err error

	p := player.P
	p.Play()

	if err = current.Update(); err != nil {
		log.Fatal(err)
	}
	log.Println(current)
}
