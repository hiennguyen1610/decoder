package main

import (
	"fmt"
	"os"
	"hiennguyen.io/decode"
)

func main(){
	filePath := "./message.txt"
	if len(os.Args) == 2 {
		filePath = os.Args[1]
	}

	text := decode.Decode(filePath)
	if text != nil && len(text) > 0 {
		fmt.Println("OUTPUT: ")
		for _, value := range text {
			fmt.Println(value)
		}
	}
	fmt.Print("Press 'Enter' to close...")
	fmt.Scanln() 
}
