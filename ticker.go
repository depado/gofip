package main

import "time"

type closableTicker struct {
	ticker *time.Ticker
	halt   chan bool
}

func (ct *closableTicker) stop() {
	ct.ticker.Stop()
	close(ct.halt)
}
