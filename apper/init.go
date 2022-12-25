package apper

import (
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/util/helper"
)

func (a *App) Init() {

	a.Cam.Orbit = camera.NewOrbitControl(a.Cam.Self)
	a.Add2Scene(a.Cam.Self)

	a.A.Gls().ClearColor(.5, .75, 2, .5)
	if a.Help {
		a.Add2Scene(helper.NewAxes(10))
	}

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
