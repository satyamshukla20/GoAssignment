package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)
func get_keys(entries map[string]bool) (keys []string){
	for k, _ := range entries {
		keys = append(keys, k)
	}
	return
}
func get_word() string{
   resp, err := http.Get("https://random-word-api.herokuapp.com/word?number=5")
   if err!=nil{
	return "elephant"
   }
   defer resp.Body.Close()
   body, err := ioutil.ReadAll(resp.Body)
   var words []string
   err = json.Unmarshal(body, &words)
   if err!=nil{
	return "elephant"
   }
   for _, word := range words {
	if len(word) > 4 && len(word) < 9 {
		return word
	}
   }
   return words[0]
}
func main(){
  // variable declarations
  word := get_word()

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
			break
		}
    		// Console display
		fmt.Println("\n")
		 fmt.Println(placeholder) // render the placeholder
		 fmt.Printf("chances: %d\n", chances) // render the chances left
		// fmt.Printf() // show the letters or words guessed till now.

	
		fmt.Println(get_keys(entries))
		fmt.Printf("Guess a letter or the word: ")

    		// take the input
		str := ""
		fmt.Scanln(&str)
		if len(str)==len(word) && str==word{
			fmt.Println("you win")
			break
		}
    		// compare and update entries, placeholder and chances.
			// check for duplicates
			_, ok := entries[str]
			if ok {
				continue
			}
			// update entries
		   entries[str]=true

		   found :=false
		   for i, v := range word {
			if str==string(v){
				placeholder[i]= string(v)
				found = true
			}
		   }
		   if !found {
			chances -=1
		   }
  	}
}