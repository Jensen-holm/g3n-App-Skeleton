package model

import (
	"github.com/g3n/engine/experimental/physics/object"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

type Plane struct {
	G    *geometry.Geometry
	Mat  *material.Standard
	Mesh *graphic.Mesh
	Body *object.Body

	W    float32
	H    float32
	A    float32
	LocX float32
	LocY float32
	locZ float32
}

// NewPlane -> Takes in width (w) height (h) and angle (a, in degrees)
// and creates a plane struct with that information
func NewPlane(w, h, a float32, color string, wire bool) *Plane {
	p := geometry.NewPlane(w, h)
	mat := material.NewStandard(
		math32.NewColor(color),
	)
	mat.SetWireframe(wire)

	// double-sided by default
	mat.SetSide(material.SideDouble)
	mesh := graphic.NewMesh(p, mat)
	mesh.RotateX(Deg2Rad(a))
	b := object.NewBody(mesh)
	return &Plane{
		G:    p,
		A:    a,
		Mat:  mat,
		Mesh: mesh,
		Body: b,
		W:    w,
		H:    h,
	}
}

func Deg2Rad(deg float32) float32 {
	return deg * (math32.Pi / 180)
}
