package ui

import (
	"github.com/Jensen-holm/g3n-Wrapper/apper"
	"github.com/g3n/engine/gui"
)

type UI struct {
	Title string
}

func NewUI(title string, a *apper.App) {
	w, h := a.A.GetSize()
	win := gui.NewWindow(float32(w), float32(h))
	win.SetTitle(title)
}
