package main

import (
	"log"
	"os"

	"ip.com/app/delivery/cli"
	"ip.com/app/repository"
	"ip.com/app/usecase"
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

	repo := repository.NewMysqlRepo(conn)
	uc := usecase.NewUseCase(repo)

	rc := cmd.NewCmd() //rootCmd

	handlers := cli.NewCliHandlers(uc, rc)
	cli.SetupHandlers(handlers)

	if err := rc.Execute(); err != nil {
		log.Print(err)
		os.Exit(1)
	}

	os.Exit(0)
}
