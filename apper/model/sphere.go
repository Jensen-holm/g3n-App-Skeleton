package model

import (
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
}

func NewSphere(r float32, c string) *Sphere {
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
	return &Sphere{
		R:    r,
		Geom: s,
		Mat:  mat,
		Mesh: mesh,
	}
}
