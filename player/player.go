package player

import (
	"log"

	"github.com/ziutek/gst"
)

// StreamURL is the URL of the FIP stream
const StreamURL = "http://audio.scdn.arkena.com/11016/fip-midfi128.mp3"

// Player is the main player structure
type Player struct {
	*gst.Element
}

// Pause pauses the Player
func (p *Player) Pause() {
	log.Println(p.SetState(gst.STATE_PAUSED), "Pausing")
	state, prev, x := P.GetState(500)
	log.Println(state, prev, x)
}

// Play starts the Player
func (p *Player) Play() {
	log.Println(p.SetState(gst.STATE_PLAYING), "Playing")
	state, prev, x := P.GetState(500)
	log.Println(state, prev, x)
}

// SetURI applies the given URI to the Player
func (p *Player) SetURI(uri string) {
	p.SetProperty("uri", uri)
}

// P is the actual player
var P Player

func init() {
	P.Element = gst.ElementFactoryMake("playbin", "player")
	if P.Element == nil {
		log.Fatal("ElementFactoryMake failed")
	}
	P.SetURI(StreamURL)
	state, prev, x := P.GetState(500)
	log.Println(state, prev, x)
}
