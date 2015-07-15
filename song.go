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
