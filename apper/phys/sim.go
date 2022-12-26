package phys

import (
	"App/apper/model"
	"github.com/g3n/engine/math32"
	"time"
)

type Sim struct {
	ConstantForces []*math32.Vector3

	Spheres []*model.Sphere
	Planes  []*model.Plane
}

func NewSim() *Sim {
	return new(Sim)
}

func (s *Sim) AddSphere(spheres ...*model.Sphere) {
	for _, sp := range spheres {
		s.Spheres = append(s.Spheres, sp)
	}
}

func (s *Sim) AddConstantForce(x, y, z float32) {
	s.ConstantForces = append(
		s.ConstantForces,
		math32.NewVector3(x, y, z),
	)
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
		dtX := initVelo.X * float32(time.Seconds())
		dtY := initVelo.Y * float32(time.Seconds())
		dtZ := initVelo.Z * float32(time.Seconds())

		sp.Update(
			dtX+initPos.X,
			dtY+initPos.Y,
			dtZ+initPos.Z,
		)
		//s.UpdateVelo(initVelo, time)
	}
}

// UpdateVelo -> updates the velocity of objects
// according to the forces in the simulation and acceleration.
// does so each frame
func (s *Sim) UpdateVelo(initVelo *math32.Vector3, time time.Duration) {
}
