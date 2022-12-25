package apper

import (
	"github.com/g3n/engine/experimental/physics"
	"time"
)

type Sim struct {
	Self *physics.Simulation
}

func NewSim() {
}

// UpdateSim -> supposed to be called each frame
func (s *Sim) UpdateSim(a *App, deltaTime time.Duration) {
	s.Self.Step(float32(deltaTime))
}

func (s *Sim) NewForce() {
}
