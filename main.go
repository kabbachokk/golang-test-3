package main

import (
	"log"
	"os"

	"ip.com/cmd"
	"ip.com/config"
	"ip.com/db"
)

func main() {
	conf, err := config.GetConfig()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	log.Print(conf)

	conn, err := db.NewMysqlConn(conf)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	defer conn.Close()

	if err := cmd.Execute(); err != nil {
		log.Print(err)
		os.Exit(1)
	}

	os.Exit(0)
}
