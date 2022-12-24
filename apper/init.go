package apper

import "github.com/g3n/engine/core"

func (a *App) Init() {

	if a.Help {

	}

}

func (a *App) Add2Scene(nodes ...*core.Node) {
	for _, n := range nodes {
		a.Scene.Add(n)
	}
}

func (a *App) ToggleHelp(h bool) {
	if h != a.Help {
		a.Help = h
	}
}
