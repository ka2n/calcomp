package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ka2n/calcomp"
)

func main() {
	var productFlag string
	var kcalFlag float64

	flag.StringVar(&productFlag, "product", "powderV3.1", "product")
	flag.Float64Var(&kcalFlag, "kcal", 100, "target kcal")
	flag.Parse()

	if kcalFlag <= 0.0 {
		fmt.Println("[Error] kcal must be positive integer")
		os.Exit(1)
	}

	product, ok := calcomp.PRODUCTS[productFlag]
	if !ok {
		fmt.Println("[Error] unknown product")

		fmt.Println("Supported products:")
		for k := range calcomp.PRODUCTS {
			fmt.Printf("- %s", k)
		}
		os.Exit(1)
	}

	mlPKCal := product.ML / product.KCal
	gPKcal := product.G / product.KCal

	ml := mlPKCal * kcalFlag
	g := gPKcal * kcalFlag

	fmt.Printf("kcal: %0.1f, ml: %0.1f, g: %0.1f\n", kcalFlag, ml, g)
}
