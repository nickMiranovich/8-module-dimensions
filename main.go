package main

import (
	"fmt"
	"myFirstModule/electronic"
)

func main() {
	iphone := electronic.NewApplePhone("9")
	samsung := electronic.NewAndroidPhone("Galaxy", "X6")
	nokia := electronic.NewStantionPhone("Nokia", "6310", 17)
	printCharacteristics(iphone)
	printCharacteristics(samsung)
	printCharacteristics(nokia)

}

func printCharacteristics(p electronic.Phone) {
	fmt.Print(p.Brand(), " ", p.Model(), " ", p.Type())
	switch p.Type() {
	case "stantion":
		p, ok := p.(electronic.StationPhone)
		if ok {
			fmt.Println(" ", p.ButtonsCount())
		}
	case "smartphone":
		p, ok := p.(electronic.Smartphone)
		if ok {
			fmt.Println(" ", p.OS())
		}
	}
}
