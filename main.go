package main

import "fmt"

func main() {

	var (
		input, output float64
	)

	fmt.Print("enter a temp in F: ")
	fmt.Scanf("%f", &input)
	output = fToC(input)
	fmt.Println("Temp in C", output)

	fmt.Print("Enter feet to conver to meters: ")
	fmt.Scanf("%f", &input)
	output = feetToMeters(input)
	fmt.Println("Meters:", output)
}

func fToC(fTemp float64) float64 {
	cTemp := (fTemp - 32) * (5.0 / 9)
	return cTemp
}

func feetToMeters(ft float64) float64 {
	meters := ft * 0.3048
	return meters
}
