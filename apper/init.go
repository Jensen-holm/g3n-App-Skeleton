package apper

import (
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/util/helper"
)

func (a *App) Init() {

	if a.Help {
		a.Add2Scene(helper.NewAxes(10))
	}

	a.Orbit = camera.NewOrbitControl(a.Cam)

}

func (a *App) Add2Scene(nodes ...core.INode) {
	for _, n := range nodes {
		a.Scene.Add(n)
	}
}

func (a *App) ToggleHelp(h bool) {
	if h != a.Help {
		a.Help = h
	}
}

func (a *App) AddSurface() {

}

func (a *App) AddLight() {

}
