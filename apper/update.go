package apper

import (
	"fmt"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/renderer"
	"log"
	"time"
)

var totTime float64 = 0

func (a *App) Update(r *renderer.Renderer, dt time.Duration) {
	a.FrameRater.Start()

	a.A.Gls().Clear(
		gls.COLOR_BUFFER_BIT | gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT,
	)

	a.Cam.UpdatePos()

	// Currently will wait five seconds before
	//updating the simulation (if there is one)
	var t = int(totTime) > 5
	if a.Sim != nil && t {
		a.Sim.Update(dt)
	}

	err := r.Render(a.Scene, a.Cam.Self)
	if err != nil {
		log.Fatalf("error rendering the scene: %v", err)
	}

	totTime += dt.Seconds()

	// Print total time each frame
	fmt.Print(fmt.Sprintf("\r%f", totTime))

	a.FrameRater.Wait()
}

func (a *App) Run() {
	a.A.Run(a.Update)
}
