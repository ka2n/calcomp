package calcomp

type powderProduct struct {
	ML   float64
	G    float64
	KCal float64
}

// PRODUCTS is a list of all products available
var PRODUCTS = map[string]powderProduct{
	"powderV3.1": powderProduct{ML: 55.0, KCal: 100.0, G: 23.0},
}
