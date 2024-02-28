package main

import (
	"github.com/anoopmoothedath/hello-world/services"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("Started Hello-World Program")

	log.Println("Hello, World!")
	log.Println("This line is being printed from feature branch")

	services.PrintString("This is being printed from a new directory services/PrintString function")
}
