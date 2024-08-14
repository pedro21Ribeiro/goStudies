package main

import (
	"fmt"
	"log"

	"example.com/greeting"
)

func main() {
	log.SetPrefix("message: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Adriana")

	if (err != nil){
		log.Fatal(err)
	}else{
		fmt.Println(message)
	}
	

	names := []string{"Mariane", "João", "Pedro"}
	
	messages, err :=greetings.Hellos(names)

	if err!=nil {
		log.Fatal(err)
	}
	
	
	for _, message := range messages{
		fmt.Println(message)
	}
}