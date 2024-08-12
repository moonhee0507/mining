package app

import (
	"bufio"
	"fmt"
	"mining/config"
	"mining/repository"
	"mining/service"
	"os"
	"time"

	"github.com/inconshreveable/log15"
)

type App struct {
	config *config.Config

	service *service.Service
	repository *repository.Repository

	log log15.Logger
}

func NewApp(config *config.Config) {
	a := &App {
		config: config,
		log: log15.New("module", "app"),
	}

	var err error

	if a.repository, err = repository.NewRepository(config); err != nil {
		panic(err)
	} else {
		a.service = service.NewService(config, a.repository)

		a.log.Info("Module Started", "time", time.Now().Unix())

		sc := bufio.NewScanner(os.Stdin)

		sc.Scan()
		fmt.Println(sc.Text())
	}
}