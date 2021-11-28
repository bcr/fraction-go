package main

import (
	"bcr/fraction"
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Print("? ")
		text, _ := reader.ReadString('\n')
		if text == "\n" {
			break
		}
		fmt.Print("= ")
		fmt.Println(fraction.Evaluate(text))
	}
}
