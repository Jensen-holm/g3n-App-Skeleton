package main

import (
	app "App/apper"
	"App/apper/model"
	"App/apper/phys"
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

	sim := phys.NewSim()
	a.AddSim(sim)
	sim.AddConstantForce(0, -.98, 0)

	ball := model.NewSphere(3, 1, 1, 3, "purple")
	ball2 := model.NewSphere(3, 0, 0, 3, "green")

	ball.SetVelo(2, 0, 0)
	ball2.SetVelo(2, 0, 0)

	sim.AddSphere(ball, ball2)

	ground := model.NewPlane(500, 500, 90, "slategray", false)

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
