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

	ball := model.NewSphere(10, 10, 10, 3, "purple")
	sim := app.NewSim(a.Scene)
	sim.AddBody(ball.Body)

	a.Add2Scene(
		a.Cam.Self,
		app.Light("white", 1, 100, 100, 100),
		app.Light("white", 1, -100, 100, -100),
		ball.Mesh,
	)
}
