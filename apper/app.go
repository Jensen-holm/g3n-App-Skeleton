package apper

import (
	"github.com/g3n/engine/app"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/experimental/physics/object"
	"github.com/g3n/engine/util"
)

type App struct {
	A          *app.Application
	Scene      *core.Node
	Cam        *Cam
	Sim        *Sim
	FrameRater *util.FrameRater
	Objs       []*object.Body
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
