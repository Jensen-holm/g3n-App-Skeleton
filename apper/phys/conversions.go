package phys

const mileFt = 5280

// G2Kg -> converts grams to kilograms
func G2Kg(g float32) float32 {
	return g / 1000
}

// Mph2Fps -> converts miles per hour
// into feet per second
func Mph2Fps(mph float32) float32 {
	ftHr := mph * mileFt
	return (ftHr / (60)) / 60
}
