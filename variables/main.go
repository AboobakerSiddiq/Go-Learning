package main

import "fmt"

var userName string = "Siddiq"

func main() {
	var username string = "hitesh"
	fmt.Println(userName)
	fmt.Printf("Variables is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variables is of type: %T \n", isLoggedIn)

	var smallValue int = 25555
	fmt.Println(smallValue)
	fmt.Printf("Variables is of type: %T \n", smallValue)

	var smallFloat float64 = 255.859857458
	fmt.Println(smallFloat)
	fmt.Printf("Variables is of type: %T \n", smallFloat)

	var garbageInt int
	fmt.Println(garbageInt)
	fmt.Printf("Variables is of type: %T \n", garbageInt)

	//lexer taking care of the type
	var sampleTXT = "Hello Siddiq"
	fmt.Println(sampleTXT)
	fmt.Printf("Variables is of type: %T \n", sampleTXT)

	//No var style
	numberOfUser := 300000
	fmt.Println(numberOfUser)
}
