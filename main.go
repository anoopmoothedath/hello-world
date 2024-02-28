package main

import "log"

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("Started Hello-World Program")
}
