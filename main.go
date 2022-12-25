package main

import (
	app "App/apper"
	"App/apper/model"
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

	ball := model.NewSphere(
		10,
		10,
		10,
		3,
		"brown",
	)

	ball2 := model.NewSphere(
		20,
		20,
		20,
		3,
		"green",
	)

	sim := app.NewSim(a.Scene)
	sim.AddConstForce(0, -1, 0)
	sim.AddBody(ball.Body, ball2.Body)

	a.Add2Scene(
		a.Cam.Self,
		app.Light("white", 1, 100, 100, 100),
		app.Light("white", 1, -100, 100, -100),
		ball.Mesh,
		ball2.Mesh,
	)
}
