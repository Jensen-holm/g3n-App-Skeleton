package apper

import (
	"github.com/Jensen-holm/g3n-Wrapper/apper/phys"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/util/helper"
)

func (a *App) Prep() {
	gui.Manager().Set(a.Scene)
	if a.Help {
		a.Add2Scene(helper.NewAxes(10))
	}
	// set up the keyboard controls for the camera

}

func (a *App) Add2Scene(nodes ...core.INode) {
	for _, n := range nodes {
		a.Scene.Add(n)
	}
}

func (a *App) AddBg(r, g, b, alpha float32) {
	a.A.Gls().ClearColor(
		r, g, b, alpha,
	)
}

func (a *App) AddSim(s *phys.Sim) {
	a.Sim = s
}

func Light(
	color string, alpha float32, posX, posY, posZ float32,
) *light.Directional {

	l := light.NewDirectional(
		math32.NewColor(color),
		alpha,
	)
	l.SetPosition(posX, posY, posZ)

	return l
}
