package main

import app "App/apper"

func main() {

	a := app.NewApp(true)
	Init(a)
	a.Run()

}

func Init(a *app.App) {

	a.AddBg(.5, .75, 2, .5)
	a.Add2Scene(
		a.Cam.Self,
		app.Light("white", 1, 100, 100, 100),
		app.Light("white", 1, -100, 100, -100),

	)

}
