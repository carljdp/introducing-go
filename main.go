package main

import "fmt"

func main() {

	text := "hello"

	printOneCharperLine(text)

}

func printOneCharperLine(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c %s", str[i], "\n")
	}
}
