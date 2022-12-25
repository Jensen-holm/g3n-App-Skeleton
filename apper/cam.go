package apper

import "github.com/g3n/engine/camera"

type Cam struct {
	Self  *camera.Camera
	Orbit *camera.OrbitControl
	Pos   map[string]float32
}

func NewCam() *Cam {
	cam := camera.New(1)
	cam.SetPosition(2, 2, 2)
	return &Cam{
		Self: cam,
		Pos: map[string]float32{
			"X": 2,
			"Y": 2,
			"Z": 2,
		},
	}
}

// UpdatePos -> Ran each frame, and detects user movement
// and updates the Pos map inside the Cam struct
func (c *Cam) UpdatePos() {
}

// Move -> Ran when the user presses certain keys to
// move them around the simulation, updates the Pos map
func (c *Cam) Move() {
}
