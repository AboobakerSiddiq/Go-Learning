package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Our Shop")
	fmt.Println("Please rate us")

	reader := bufio.NewReader(os.Stdin)

	rating, _ := reader.ReadString('\n')

	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 32)

	if err != nil {
		fmt.Println("Not a valid Rating:", err)
	} else {
		fmt.Println("New Rating:", numRating)
	}

}
