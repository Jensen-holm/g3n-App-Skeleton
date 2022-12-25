package main

import app "App/apper"

func main() {

	a := app.NewApp()
	a.ToggleHelp(true)
	a.Init()
	a.Run()

}
