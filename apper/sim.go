package apper

import (
	"App/apper/model"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/experimental/physics"
	"github.com/g3n/engine/math32"
	"time"
)

type Sim struct {
	Self    *physics.Simulation
	Spheres []*model.Sphere
	Planes  []*model.Plane
	//Anim *texture.Animator
}

func NewSim(scene *core.Node) *Sim {
	return &Sim{
		Self: physics.NewSimulation(scene),
	}
}

func (s *Sim) AddSphere(spheres ...*model.Sphere) {
	for i, sp := range spheres {
		s.Self.AddBody(sp.Body, "obj#"+string(rune(i)))
		s.Spheres = append(s.Spheres, sp)
	}
}

func (s *Sim) AddPlane(planes ...*model.Plane) {
	for _, p := range planes {
		s.Planes = append(s.Planes, p)
		s.Self.AddBody(p.Body, "plane")
	}
}

func (s *Sim) AddConstForce(x, y, z float32) {
	s.Self.AddForceField(physics.NewConstantForceField(
		&math32.Vector3{
			X: x,
			Y: y,
			Z: z,
		},
	))
}

// UpdateSim -> supposed to be called each frame
func (s *Sim) UpdateSim(deltaTime time.Duration) {
	s.Self.Step(float32(deltaTime.Seconds()))
}

func (s *Sim) UpdateObjsPos() {
	for _, s := range s.Spheres {
		s.UpdatePos()
	}
}
