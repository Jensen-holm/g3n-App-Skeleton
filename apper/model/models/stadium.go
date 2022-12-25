package stadium

import (
	"App/apper/model"
	"github.com/g3n/engine/core"
	"log"
	"os"
)

func ModDir() string {
	wd, _ := os.Getwd()
	return wd + "/apper/model/models/"
}

func Soccer() *core.Node {
	g, err := model.Read(
		ModDir()+"Soccer.obj",
		ModDir()+"Soccer.mtl",
	)

	if err != nil {
		log.Fatalf("error reading the soccer obj | mtl files: %v", err)
	}

	return g
}
