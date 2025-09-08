package main

import (
	"/opt/Project/RIP/LAB1/internal/api/"
	"log"
)

func main() {
	log.Println("Application start!")
	api.StartServer()
	log.Println("Application terminated!")
}
