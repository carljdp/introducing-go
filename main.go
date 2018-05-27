package main

import "fmt"

func main() {

	fmt.Println(" --- for loop ---")

	text := "hello"
	printOneCharPerLine(text)

	fmt.Println(" --- If statement ---")

	evenOrOdd(5, 11)
	evenOrOdd(2, 4)

	fmt.Println(" --- Switch statement ---")

	fmt.Println(nameOfWeekDay(1))
	fmt.Println(nameOfWeekDay(3))
	fmt.Println(nameOfWeekDay(8))

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
		dayName = ""
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
