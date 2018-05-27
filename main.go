package main

import "fmt"

func main() {

	fmt.Println()
	fmt.Println(" --- for loop ---")

	text := "hello"
	printOneCharPerLine(text)

	fmt.Println()
	fmt.Println(" --- If statement ---")

	evenOrOdd(5, 11)
	evenOrOdd(2, 4)

	fmt.Println()
	fmt.Println(" --- Switch statement ---")

	fmt.Println(nameOfWeekDay(1))
	fmt.Println(nameOfWeekDay(3))
	fmt.Println(nameOfWeekDay(8))

	fmt.Println()
	fmt.Println(" --- Exercises ---")

	numbersDivisibleby(3, 10)
	numbersDivisibleby(5, 56)

	fmt.Println()
	fizzBuzz()

}

func numbersDivisibleby(divisor, upperLimit int) {

	isFirstRun := true

	for i := 0; i <= upperLimit; i++ {

		if i%divisor == 0 {

			if !isFirstRun {
				fmt.Print(", ")
			}
			fmt.Print(i)
		}
		isFirstRun = false
	}
	fmt.Println()
}

func fizzBuzz() {
	for i := 1; i <= 15; i++ {
		fmt.Print(i)
		fmt.Print(" ")
		if i%3 == 0 {
			fmt.Print("Fizz")
		}
		if i%5 == 0 {
			fmt.Print("Buzz")
		}
		fmt.Println()
	}
}

func nameOfWeekDay(day int8) string {
	var dayName string
	switch day {
	case 1:
		dayName = "Monday"
	case 2:
		dayName = "Tuesday"
	case 3:
		dayName = "Wednesday"
	case 4:
		dayName = "Thursday"
	case 5:
		dayName = "Friday"
	case 6:
		dayName = "Saturday"
	case 7:
		dayName = "Sunday"
	default:
		dayName = "<unknown>"
	}
	return dayName
}

func evenOrOdd(start, end int) {

	var msg string
	for i := start; i <= end; i++ {
		if i%2 == 0 {
			msg = "even"
		} else {
			msg = "odd"
		}
		fmt.Println(i, "is", msg)
	}
}

func printOneCharPerLine(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c %s", str[i], "\n")
	}
}
