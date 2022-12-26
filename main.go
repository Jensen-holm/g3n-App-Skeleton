package main

import (
	app "App/apper"
	"App/apper/model"
	"github.com/g3n/engine/experimental/physics/object"
)

func main() {
	a := app.NewApp(true)
	Init(a)
	a.Run()
}

// Init -> This is where the apps
// are customized using apper
func Init(a *app.App) {
	a.AddBg(.5, .75, 2, .5)

	ball := model.NewSphere(0, 0, 0, 3, "purple")
	ball2 := model.NewSphere(0, 0, 0, 3, "green")
	ground := model.NewPlane(500, 500, 90, "slategray", false)

	ground.Body.SetBodyType(object.Static)

	a.Sim = app.NewSim(a.Scene)
	a.Sim.AddSphere(ball, ball2)
	a.Sim.AddConstForce(0, 0, 0)
	a.Sim.AddPlane(ground)

	l1 := app.Light("white", 1, 100, 100, 100)
	l2 := app.Light("white", 1, -100, 100, -100)
	l3 := app.Light("white", 1, -100, -100, -100)
	l4 := app.Light("white", 1, 100, -100, 100)

	a.Add2Scene(
		a.Cam.Self,
		l1, l2, l3, l4,
		ground.Mesh,
		ball.Mesh,
		ball2.Mesh,
	)
}
