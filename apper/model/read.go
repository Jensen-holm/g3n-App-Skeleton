package model

import (
	"github.com/g3n/engine/core"
	obj2 "github.com/g3n/engine/loader/obj"
)

func Read(obj, mtl string) (*core.Node, error) {
	dec, err := obj2.Decode(
		obj, mtl,
	)
	group, err := dec.NewGroup()
	return group, err
}
