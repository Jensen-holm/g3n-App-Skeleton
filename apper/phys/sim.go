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

//	func (s *Sim) UpdateObjsPos() {
//		for _, sp := range s.Spheres {
//			sp.UpdatePos()
//		}
//	}
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

// Update -> to be run each frame
func (s *Sim) Update(time time.Duration) {
	for _, sp := range s.Spheres {

		initVelo := sp.Velo
		initPos := sp.Pos

		// calculate distance traveled since last frame
		distTraveledX := initVelo.X * float32(time.Seconds())
		distTraveledY := initVelo.Y * float32(time.Seconds())
		distTraveledZ := initVelo.Z * float32(time.Seconds())

		// update the objects position in space
		//sp.Pos.X = distTraveledX + initVelo.X
		//sp.Pos.Y = distTraveledY + initVelo.Y
		//sp.Pos.Z = distTraveledZ + initVelo.Z

		sp.Update(
			distTraveledX+initPos.X,
			distTraveledY+initPos.Y,
			distTraveledZ+initPos.Z,
		)

		// update the velocity based on acceleration ??
		//for _, f := range s.ConstantForces {
		//}
	}
}
