package apper

import (
	"github.com/Jensen-holm/g3n-Wrapper/apper/phys"
	"github.com/g3n/engine/app"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/experimental/physics/object"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/util"
	"time"
)

type App struct {
	A            *app.Application
	Scene        *core.Node
	Cam          *Cam
	Sim          *phys.Sim
	FrameRater   *util.FrameRater
	Objs         []*object.Body
	CustomUpdate func(renderer2 *renderer.Renderer, dt time.Duration)
	Help         bool
}

func NewApp(helper bool) *App {
	a := &App{
		A:          app.App(),
		Scene:      core.NewNode(),
		Cam:        NewCam(),
		FrameRater: util.NewFrameRater(60),
		Help:       helper,
	}
	a.Prep()
	return a
}

func (a *App) SetUpdateFunc(customUpdate func(renderer *renderer.Renderer, dt time.Duration)) {
	a.CustomUpdate = customUpdate
}
