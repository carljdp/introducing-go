package main

import "fmt"

func main() {

	text := "hello"
	printOneCharperLine(text)

	evenOrOdd(5, 11)

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

func printOneCharperLine(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c %s", str[i], "\n")
	}
}
