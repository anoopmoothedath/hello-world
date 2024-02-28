package main

import "log"

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("Started Hello-World Program")

	log.Println("Hello, World!")
	log.Println("This line is being printed from feature branch")
}
