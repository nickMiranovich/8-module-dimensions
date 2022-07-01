package main

import (
	"fmt"
	"myFirstModule/dimensions"
)

func main() {

	u1 := dimensions.NewUnit(12.1, "inch")
	u2 := dimensions.NewUnit(3.1, "inch")
	u3 := dimensions.NewUnit(24.6, "cm")
	d1 := dimensions.NewDimensions(u1, u2, u3, "inch")
	d2 := dimensions.NewDimensions(u1, u2, u3, "cm")
	printDimensions(d1, "inch")
	printDimensions(d2, "cm")
}

func printDimensions(d dimensions.Dimensions, u string) {
	fmt.Printf("IN %v dimensions: height: %v, length: %v, width: %v\n", u, d.Height().Value, d.Length().Value, d.Width().Value)
}
