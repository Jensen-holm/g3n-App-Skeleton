package apper

import (
	"fmt"
	"log"
	"time"

	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/renderer"
)

func (a *App) Update(r *renderer.Renderer, dt time.Duration) {
	a.FrameRater.Start()

	a.A.Gls().Clear(
		gls.COLOR_BUFFER_BIT | gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT,
	)

	// this will be used in future to get
	// the position of the camera, and move it
	// according to keyboard controls
	a.Cam.UpdatePos()

	// Currently will wait five seconds before
	//updating the simulation (if there is one)
	var t = int(a.A.RunTime().Seconds()) >= 5
	if a.Sim != nil && t {
		a.Sim.Update(dt)
	}

	err := r.Render(a.Scene, a.Cam.Self)
	if err != nil {
		log.Fatalf("error rendering the scene: %v", err)
	}

	// Print total time each frame
	fmt.Print(fmt.Sprintf("\rTotal Time: %f", a.A.RunTime().Seconds()))

	a.FrameRater.Wait()
}

// Run -> if the user set their own custom update function,
// use that one, else run the default one that we have above
func (a *App) Run() {

	if a.CustomUpdate != nil {
		a.A.Run(a.CustomUpdate)
	} else {
		a.A.Run(a.Update)
	}

}

// GetFps -> for the user to use should they choose to
// create and use their own custom update function
func GetFps(dt time.Duration, iters int) float32 {
	return float32(iters) / float32(dt.Seconds())
}
