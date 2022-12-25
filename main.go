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

	// Background
	a.AddBg(.5, .75, 2, .5)

	// making a ball
	ball := model.NewSphere(3, "purple")

	// Adding things to the scene
	a.Add2Scene(

		// camera & lights
		a.Cam.Self,
		app.Light("white", 1, 100, 100, 100),
		app.Light("white", 1, -100, 100, -100),

		// objects
		ball.Mesh,
	)
}
