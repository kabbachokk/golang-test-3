package main

import (
	"log"

	"ip.com/app/delivery/console"
	"ip.com/config"
)

func main() {
	c, err := config.GetConfig()
	if err != nil {
		log.Print(err)
	}
	log.Print(c)
	console.Execute()
}
