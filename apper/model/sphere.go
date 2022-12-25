package model

import (
	"github.com/g3n/engine/experimental/physics/object"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

type Sphere struct {
	R    float32
	Geom *geometry.Geometry
	Mat  *material.Standard
	Mesh *graphic.Mesh
	Body *object.Body
	Pos  *math32.Vector3
}

func NewSphere(x, y, z, r, vX, vY, vZ float32, c string) *Sphere {
	var s = geometry.NewSphere(
		float64(r),
		100,
		100,
	)

	// TODO : make the material more customizable
	mat := material.NewStandard(
		math32.NewColor(c),
	)
	mesh := graphic.NewMesh(s, mat)
	b := object.NewBody(mesh)
	b.SetPosition(x, y, z)
	b.SetVelocity(math32.NewVector3(vX, vY, vZ))
	return &Sphere{
		R:    r,
		Geom: s,
		Mat:  mat,
		Mesh: mesh,
		Body: b,
		Pos: &math32.Vector3{
			X: x,
			Y: y,
			Z: z,
		},
	}
}

func (s *Sphere) UpdatePos() {
	p := s.Mesh.Position()
	s.Pos = &p
}
