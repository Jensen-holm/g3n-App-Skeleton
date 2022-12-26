package apper

import (
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/math32"
)

type Cam struct {
	Self  *camera.Camera
	Orbit *camera.OrbitControl
	Pos   *math32.Vector3
}

func NewCam() *Cam {
	cam := camera.New(1)
	cam.SetPosition(2, 2, 2)
	return &Cam{
		Self: cam,
		Pos: &math32.Vector3{
			X: 2,
			Y: 2,
			Z: 2,
		},
	}
}

// UpdatePos -> Ran each frame, and detects user movement
// and updates the Pos map inside the Cam struct
func (c *Cam) UpdatePos() {
	p := c.Self.Position()
	if &p != c.Pos {
		c.Pos = &p
	}
}

// Move -> Ran when the user presses certain keys to
// move them around the simulation, updates the Pos map
func (c *Cam) Move(x, y, z float32) {

	// move the camera to those coords
	// the Pos vector will be updated in the
	// UpdatePos function
	c.Self.SetPosition(x, y, z)
}
