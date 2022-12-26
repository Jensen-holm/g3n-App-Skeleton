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

func (s *Sim) Update(time time.Duration) {
	s.UpdateObjs(time)
	//s.CheckCollisions()
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
		s.ApplyGravity(sp, time)
		s.CheckCollisions()
	}
}

func (s *Sim) ApplyGravity(sphere *model.Sphere, time time.Duration) {
	// force = mass * acceleration
	a := s.Gravity.Y * float32(time.Seconds())
	sphere.Velo.Y = (sphere.Mass * a) + sphere.OuterForce.Y
}

func (s *Sim) CheckCollisions() {
	if s.Plane == nil {
		return
	}
	// check if they have hit the ground
	// if so, for now we will just keep it on the ground
	for _, sphere := range s.Spheres {
		if sphere.Pos.Y-sphere.R <= s.Plane.LocY {
			sphere.ApplyForce(
				sphere.Velo.X,
				-sphere.Velo.Y,
				sphere.Velo.Z,
			)
		}
	}
}
