package apper

import (
	"fmt"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/renderer"
	"log"
	"time"
)

func (a *App) Update(r *renderer.Renderer, dt time.Duration) {
	a.FrameRater.Start()

	a.A.Gls().Clear(
		gls.COLOR_BUFFER_BIT | gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT,
	)

	if a.Sim != nil {
		a.Sim.UpdateSim(dt)
		a.Sim.UpdateObjsPos()
		fmt.Println(a.Sim.Spheres[0].Pos)
	}

	a.Cam.UpdatePos()
	//a.Cam.Self.LookAt(a.Sim.Spheres[1].Pos, a.Sim.Spheres[1].Pos)

	err := r.Render(a.Scene, a.Cam.Self)
	if err != nil {
		log.Fatalf("error rendering the scene: %v", err)
	}

	a.FrameRater.Wait()
}

func (a *App) Run() {
	a.A.Run(a.Update)
}
