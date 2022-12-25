package apper

import (
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/experimental/physics"
	"github.com/g3n/engine/experimental/physics/object"
	"time"
)

type Sim struct {
	Self *physics.Simulation
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

// UpdateSim -> supposed to be called each frame
func (s *Sim) UpdateSim(a *App, deltaTime time.Duration) {
	s.Self.Step(float32(deltaTime))
}

func (s *Sim) NewForce() {
}
