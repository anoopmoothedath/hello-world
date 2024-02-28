package main

import (
	"github.com/anoopmoothedath/hello-world/services"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("Started Hello-World Program")

	services.PrintString("This is being printed from a new directory services/PrintString function")
}
