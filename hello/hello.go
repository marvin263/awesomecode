package main


import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main(){
	log.SetPrefix("greetings...")
	log.SetFlags(0)
	
	// A slice of names
	names :=[]string{"Gladys", "Samantha", "Darrin"}

	// Request a greeting messages.
	messages, err := greetings.Hellos(names)
        
	//If an error was returned, print it to the console and exit the program
	if err != nil {
		fmt.Println("Received messages and err")
		fmt.Println("Fatal makes program exit!!!")
		log.Fatal(err)
	}

	// In no error was returned, print the returned messages to the console
	fmt.Println(messages)
}


