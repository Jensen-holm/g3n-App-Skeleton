package phys

import (
	"App/apper/model"
	"github.com/g3n/engine/math32"
	"time"
)

type Sim struct {
	Gravity *math32.Vector3
	Plane   *model.Plane

	Spheres []*model.Sphere
}

func NewSim() *Sim {
	return new(Sim)
}

func (s *Sim) AddSphere(spheres ...*model.Sphere) {
	for _, sp := range spheres {
		s.Spheres = append(s.Spheres, sp)
	}
}

func (s *Sim) SetPlane(p *model.Plane) {
	s.Plane = p
}

func (s *Sim) SetGravity(x, y, z float32) {
	s.Gravity = math32.NewVector3(x, y, z)
}

// Init -> if there are forces in the sim,
// initialize them by changing the velocity of
// the objects in the simulation
func (s *Sim) Init() {
}

// UpdateObjs -> to be run each frame
func (s *Sim) UpdateObjs(time time.Duration) {
	for _, sp := range s.Spheres {

		initVelo := sp.Velo
		initPos := sp.Pos

		// calculate distance traveled since last frame
		dtX := (initVelo.X + sp.OuterForce.X) * float32(time.Seconds())
		dtY := (initVelo.Y + sp.OuterForce.Y) * float32(time.Seconds())
		dtZ := (initVelo.Z + sp.OuterForce.Z) * float32(time.Seconds())

		sp.Update(
			dtX+initPos.X,
			dtY+initPos.Y,
			dtZ+initPos.Z,
		)
		s.ApplyGravity(sp)
	}
}

func (s *Sim) ApplyGravity(sphere *model.Sphere) {
	// force = mass * acceleration
	sphere.Velo.Y = (s.Gravity.Y * sphere.Mass) + sphere.OuterForce.Y
}

func (s *Sim) CheckCollisions() {

	// check if they have hit the ground
	for _, sphere := range s.Spheres {

		// check if the sphere is above the surface
		// not what this is doing ...
		if sphere.Pos.Y <= s.Plane.W {
			// shoot it back up

		}

		// if not, then it could not hit the surface

	}

}
