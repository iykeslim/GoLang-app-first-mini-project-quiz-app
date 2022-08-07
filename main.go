package main

import "fmt"

func main() {
	fmt.Println("Welcoee to a Quiz Game Y'all")
	var name string
	fmt.Printf("Please Enter your name: \n")

	// getting user input
	fmt.Scan(&name)

	// using formatted string
	fmt.Printf("hello %v, You may start playing \n", name)

	fmt.Printf("Enter your age: \n")
	var age uint
	fmt.Scan(&age)

	// conditionals
	if age >= 15 {
		fmt.Println("you are old enough to play")
	} else {
		fmt.Println("sorry, you are not allowed to play")
		return
	}
	var score int
	answer_count := 2

	fmt.Printf("what is the difference between black and white? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	// concatination and logical operators
	if answer+" "+answer2 == "colour spelling" || answer+" "+answer2 == "COLOUR SPELLING" {
		fmt.Println("you are correct")
		score++
		fmt.Printf("Good %v, you are correct \n", name)
	} else {
		fmt.Printf("Sorry %v, you are wrong \n", name)
	}

	fmt.Println("how many characters does the word colour have")
	var response int
	fmt.Scan(&response)
	if response == 6 {
		fmt.Printf("%v, you are correct again \n", name)
		score++
	} else {
		fmt.Printf("soryy %v, it is incorrect \n", name)
	}

	fmt.Printf("%v, you scored %v out of %v \n", name, score, answer_count)

	//type conversion because int / int wll not return int sometimes -> error
	percent_score := (float64(score) / float64(answer_count)) * 100
	fmt.Printf("You got a score of: %v%%. %v \n", percent_score, name)
}
