package model

import (
	"fmt"
	"github.com/g3n/engine/core"
	obj2 "github.com/g3n/engine/loader/obj"
	"strings"
)

func IsFileType(p string, ft string) bool {

	isFile := strings.Contains(p, ".")
	if !isFile {
		return false
	}

	s := strings.Split(p, ".")
	isModel := s[len(s)-1] == ft
	if !isModel {
		return false
	}
	return true
}

func Read(obj, mtl string) (*core.Node, error) {
	if !IsFileType(obj, "obj") || IsFileType(mtl, "mtl") {
		return nil, fmt.Errorf(
			"invalid file type in one of the obj | mtl files",
		)
	}

	dec, err := obj2.Decode(
		obj, mtl,
	)
	group, err := dec.NewGroup()

	return group, err
}

func ReadAll(groups ...*core.Node) []*core.Node {
	objs := make([]*core.Node, 0)
	for _, g := range groups {
		objs = append(objs, g)
	}

	return objs
}
