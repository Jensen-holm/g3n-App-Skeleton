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
	Mass float32

	Velo *math32.Vector3

	Geom *geometry.Geometry
	Mat  *material.Standard
	Mesh *graphic.Mesh
	Body *object.Body
	Pos  *math32.Vector3
}

func NewSphere(x, y, z, r, m float32, c string) *Sphere {
	var s = geometry.NewSphere(
		float64(r),
		10,
		10,
	)

	// TODO : make the material more customizable
	mat := material.NewStandard(
		math32.NewColor(c),
	)
	mesh := graphic.NewMesh(s, mat)
	b := object.NewBody(mesh)
	b.SetPosition(x, y, z)
	v := math32.NewVector3(0, 0, 0)
	return &Sphere{
		R:    r,
		Geom: s,
		Mat:  mat,
		Mesh: mesh,
		Body: b,
		Velo: v,
		Mass: m,
		Pos: &math32.Vector3{
			X: x,
			Y: y,
			Z: z,
		},
	}
}

func (s *Sphere) SetVelo(vX, vY, vZ float32) {
	s.Velo = &math32.Vector3{
		X: vX,
		Y: vY,
		Z: vZ,
	}
}

func (s *Sphere) Update(newX, newY, newZ float32) {
	s.Body.SetPosition(newX, newY, newZ)
	s.Pos.X = newX
	s.Pos.Y = newY
	s.Pos.Z = newZ
}

func (s *Sphere) UpdateVelo(deltaX, deltaY, deltaZ float32) {
	s.Velo = math32.NewVector3(
		s.Velo.X+deltaX,
		s.Velo.Y+deltaY,
		s.Velo.Z+deltaZ,
	)
}
