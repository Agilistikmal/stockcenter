package cli

import "fmt"

func Scan(text string) string {
	fmt.Println("##", text)
	var input string
	fmt.Print("> ")
	fmt.Scanln(&input)
	return input
}
