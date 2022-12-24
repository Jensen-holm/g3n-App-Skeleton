package apper

import (
	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/util"
)

// the apper package serves as a wrapper
// for the g3n application package

type App struct {
	A          *app.Application
	Scene      *core.Node
	Cam        *camera.Camera
	Orbit      *camera.OrbitControl
	FrameRater *util.FrameRater
	Help       bool
}

func NewApp() *App {
	return &App{
		A:    app.App(),
		Help: false,
	}
}
