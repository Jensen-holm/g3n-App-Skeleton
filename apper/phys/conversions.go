package phys

import "github.com/g3n/engine/math32"

const mileFt = 5280

func Kg2G(kg float32) float32 {
	return kg * 1000
}

func Mph2Fps(mph float32) float32 {
	ftHr := mph * mileFt
	return (ftHr / (60)) / 60
}

func Cm2Ft(cm float32) float32 {
	return (cm * 0.393701) / 12
}

func Deg2Rad(deg float32) float32 {
	return deg * (math32.Pi / 180)
}
