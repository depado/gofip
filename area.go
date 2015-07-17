package main

import (
	"image"

	"github.com/andlabs/ui"
)

// Defines a new ui.AreaHandler implementation
type areaHandler struct {
	img *image.RGBA
}

// Satisfies the ui.AreaHandler interface
func (a *areaHandler) Paint(rect image.Rectangle) *image.RGBA {
	return a.img.SubImage(rect).(*image.RGBA)
}

// Satisfies the ui.AreaHandler interface
func (a *areaHandler) Mouse(me ui.MouseEvent) {}

// Satisfies the ui.AreaHandler interface
func (a *areaHandler) Key(ke ui.KeyEvent) bool { return false }
