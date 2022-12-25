package apper

import (
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/experimental/physics"
	"github.com/g3n/engine/experimental/physics/object"
	"github.com/g3n/engine/math32"
	"time"
)

type Sim struct {
	Self *physics.Simulation
	//Anim *texture.Animator
}

func NewSim(scene *core.Node) *Sim {
	return &Sim{
		Self: physics.NewSimulation(scene),
	}
}

func (s *Sim) AddBody(bodies ...*object.Body) {
	for i, b := range bodies {
		s.Self.AddBody(b, "obj#"+string(rune(i)))
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

func (s *Sim) AddAttractor() {
}

// UpdateSim -> supposed to be called each frame
func (s *Sim) UpdateSim(deltaTime time.Duration) {
	s.Self.Step(
		float32(deltaTime.Microseconds()),
	)
}

func (s *Sim) NewForce() {
}
