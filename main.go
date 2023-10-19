package main

import (
	"log"
	"os"

	"github.com/kabbachokk/golang-test-3/app/delivery/cli"
	"github.com/kabbachokk/golang-test-3/app/repository"
	"github.com/kabbachokk/golang-test-3/app/usecase"
	"github.com/kabbachokk/golang-test-3/cmd"
	"github.com/kabbachokk/golang-test-3/config"
	"github.com/kabbachokk/golang-test-3/db"
)

func main() {
	conf, err := config.GetConfig()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	conn, err := db.NewMysqlConn(conf)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	defer conn.Close()

	repo := repository.NewMysqlRepo(conn)
	uc := usecase.NewUseCase(repo)

	rootCmd := cmd.NewCmd()

	handlers := cli.NewCliHandlers(uc, rootCmd)
	cli.SetupHandlers(handlers)

	if err := rootCmd.Execute(); err != nil {
		log.Print(err)
		os.Exit(1)
	}

	os.Exit(0)
}
