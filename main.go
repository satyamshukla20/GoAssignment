package main

import (
	"fmt"
	"strings"
)

func main(){
  // variable declarations
  word := "mamoth"

	// lookup for entries made by the user.
	entries := map[string]bool {}
  
  	// list of "_" corrosponding to the number of letters in the word. [ _ _ _ _ _ ]
	placeholder := []string{}


	for i:=0; i<len(word); i++{
		placeholder = append(placeholder, "_")
	}

    chances:=8
	for {
		// evaluate a loss! If user guesses a wrong letter or the wrong word, they lose a chance.
		userInput:= strings.Join(placeholder,"")
	    if chances ==0 && userInput!=word {
			fmt.Println("game over, try again")
			break
		}
		// evaluate a win!
	    if userInput==word{
			fmt.Println("you win")
		}
    		// Console display
		fmt.Println("\n")
		 fmt.Println(placeholder) // render the placeholder
		// fmt.Printf() // render the chances left
		// fmt.Printf() // show the letters or words guessed till now.

		keys:=[] string{}
		for k, _ := range entries {
			keys = append(keys, k)
		}
		fmt.Println(keys)
		fmt.Printf("Guess a letter or the word: ")

    		// take the input

		str := ""
		fmt.Scanln(&str)
    		// compare and update entries, placeholder and chances.
		entries[str]=true
  	}
}