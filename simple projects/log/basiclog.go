package main

import "log"

func main() {
	log.Println("Hello world!")
	log.Fatal("This logs a message then calls os.Exit(1)")
	log.Panic("This logs a message then calls panic()")
	log.Println("This will not work after fatal or panic")

}
