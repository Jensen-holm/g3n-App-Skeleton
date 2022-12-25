package main

import (
	app "App/apper"
	stadium "App/apper/model/models"
)

func main() {

	a := app.NewApp(true)
	Init(a)
	a.Run()

}

func Init(a *app.App) {

	a.AddBg(.5, .75, 2, .5)
	a.Add2Scene(
		// camera & lights
		a.Cam.Self,
		app.Light("white", 1, 100, 100, 100),
		app.Light("white", 1, -100, 100, -100),

		// objects
		stadium.Soccer(),
	)

}
