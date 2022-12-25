package model

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/math32"
)

type Plane struct {
	G    *geometry.Geometry
	W    float32
	H    float32
	A    float32
	LocX float32
	LocY float32
	locZ float32
}

// NewPlane -> Takes in width (w) height (h) and angle (a, in degrees)
// and creates a plane struct with that information
func NewPlane(w, h, a float32) *Plane {
	return &Plane{
		G: geometry.NewPlane(w, h),
		A: Deg2Rad(a),
		W: w,
		H: h,
	}
}

func Deg2Rad(deg float32) float32 {
	return deg * (math32.Pi / 180)
}
