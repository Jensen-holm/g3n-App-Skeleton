package main

import (
	"math/rand"

	"github.com/Jensen-holm/g3n-Wrapper/apper"
	"github.com/Jensen-holm/g3n-Wrapper/apper/model"
	"github.com/Jensen-holm/g3n-Wrapper/apper/phys"
)

// main -> demonstration of how you can use
// the g3n-Wrapper/apper package in your project
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
	sim.SetGravity(0, -32.2, 0)

	ground := model.NewPlane(10000, 10000, 90, "slategray", false)
	sim.SetPlane(ground)

	l1 := apper.Light("white", 1, 100, 100, 100)
	l2 := apper.Light("white", 1, -100, 100, -100)
	l3 := apper.Light("white", 1, -100, -100, -100)
	l4 := apper.Light("white", 1, 100, -100, 100)

	a.Add2Scene(
		a.Cam.Self,
		l1, l2, l3, l4,

		// If you want an invisible plane,
		// create the ground but don't add
		// its mesh to the scene
		ground.Mesh,
	)

	for i := 0; i < 20; i++ {
		ball := model.NewSphere(float32(rand.Intn(50)), float32(rand.Intn(50)), float32(rand.Intn(50)), float32(rand.Intn(50)), float32(rand.Intn(50)), "blue")
		sim.AddSphere(ball)
		ball.ApplyForce(float32(rand.Intn(500)), float32(rand.Intn(500)), float32(rand.Intn(500)))
		a.Add2Scene(ball.Mesh)
	}

}
