package main

import (
	"fmt"
	"strings"
)

func main() {
  text:="Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software."

	var countUpperCase int
	for len(text) > 0 {
		spaceIndex := strings.Index(text, " ")
    if(spaceIndex == -1){
      spaceIndex = len(text)-1
    }
		word := strings.Trim(text[:spaceIndex], " ,.")    
		if word == strings.Title(word) {
			countUpperCase++
     fmt.Println(word)
		}
		text = text[spaceIndex+1:]
	}
  fmt.Println(countUpperCase)
}
