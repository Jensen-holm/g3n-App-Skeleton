package apper

import (
	"github.com/g3n/engine/app"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/util"
)

// the apper package serves as a wrapper
// for the g3n application package

type App struct {
	A          *app.Application
	Scene      *core.Node
	Cam        *Cam
	Sim        *Sim
	FrameRater *util.FrameRater
	Help       bool
}

func NewApp(helper bool) *App {
	var a = &App{
		A:          app.App(),
		Scene:      core.NewNode(),
		Cam:        NewCam(),
		FrameRater: util.NewFrameRater(60),
		Help:       helper,
	}
	a.Prep()
	return a
}
