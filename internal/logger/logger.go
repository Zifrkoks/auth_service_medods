package logger

import "fmt"

func Log(text string) {

	fmt.Println(text)
}
func LogImportant(text string) {
	fmt.Println("===================================")
	fmt.Println(text)
	fmt.Println("===================================")
}
