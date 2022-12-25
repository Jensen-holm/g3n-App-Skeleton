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
		0,
		0,
		1,
		"brown",
	)

	ball2 := model.NewSphere(
		20,
		20,
		20,
		3,
		0,
		0,
		1,
		"green",
	)

	a.Sim = app.NewSim(a.Scene)
	a.Sim.AddSphere(ball, ball2)

	a.Add2Scene(
		a.Cam.Self,
		app.Light("white", 1, 100, 100, 100),
		app.Light("white", 1, -100, 100, -100),
		ball.Mesh,
		ball2.Mesh,
	)
}
