package main

import "fmt"

func main() {

	x := [10]float64{1, 2, 3, 4}

	fmt.Println(x)

	fmt.Println(avg(x[:]))
}

func avg(x []float64) float64 {

	var total, count float64

	for _, value := range x {
		if value > 0 {
			count++
		}

		total += value
	}

	return total / count

}
