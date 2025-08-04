package main

import "fmt"

func main() {
	var assets, liabilities, units float64

	fmt.Print("Enter total assets: ")
	fmt.Scanln(&assets)

	fmt.Print("Enter total liabilities: ")
	fmt.Scanln(&liabilities)

	fmt.Print("Enter outstanding units: ")
	fmt.Scanln(&units)

	nav := (assets - liabilities) / units
	fmt.Printf("Calculated NAV: %.4f\n", nav)
}
