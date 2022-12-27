package main

import (
	"github.com/Jensen-holm/g3n-App-Skeleton/apper"
	"github.com/Jensen-holm/g3n-App-Skeleton/apper/model"
	"github.com/Jensen-holm/g3n-App-Skeleton/apper/phys"
)

func main() {
	a := apper.NewApp(true)
	Init(a)
	a.Run()
}

// Init -> This is where the apps
// are customized using apper wrapper
func Init(a *apper.App) {
	a.AddBg(.5, .75, 2, .5)

	sim := phys.NewSim()
	a.AddSim(sim)
	sim.SetGravity(0, -9.8, 0)

	ball := model.NewSphere(0, 100, 0, 3, 145, "red")
	ball.ApplyForce(50, 50, 50)
	sim.AddSphere(ball)

	ground := model.NewPlane(10000, 10000, 90, "slategray", false)
	sim.SetPlane(ground)

	l1 := apper.Light("white", 1, 100, 100, 100)
	l2 := apper.Light("white", 1, -100, 100, -100)
	l3 := apper.Light("white", 1, -100, -100, -100)
	l4 := apper.Light("white", 1, 100, -100, 100)

	a.Add2Scene(
		a.Cam.Self,
		l1, l2, l3, l4,
		ground.Mesh,
		ball.Mesh,
	)
}
