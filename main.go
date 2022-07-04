package main

import (
	"fmt"
	"myFirstModule/dimensions"
)

func main() {
	iphone := dimensions.NewApplePhone("9")
	samsung := dimensions.NewAndroidPhone("Galaxy", "X6")
	nokia := dimensions.NewStantionPhone("Nokia", "6310", 17)
	printCharacteristics(iphone)
	printCharacteristics(samsung)
	printCharacteristics(nokia)

}

func printCharacteristics(p dimensions.Phone) {
	fmt.Print(p.Brand(), " ", p.Model(), " ", p.Type())
	switch p.Type() {
	case "stantion":
		p, ok := p.(dimensions.StationPhone)
		if ok {
			fmt.Println(" ", p.ButtonsCount())
		}
	case "smartphone":
		p, ok := p.(dimensions.Smartphone)
		if ok {
			fmt.Println(" ", p.OS())
		}
	}
}
