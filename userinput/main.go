package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var welcome string = "Welcome to user input"
	fmt.Println((welcome))
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Rating:")

	//comma ok || err err
	// first part is try & second part is catch
	input, err := reader.ReadString('\n')
	fmt.Println("Thanks for rating,", input)
	fmt.Printf("Type: %T\n", input)
	fmt.Printf("Type: %T\n", err)

}
