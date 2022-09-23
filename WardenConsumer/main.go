package main

import (
	"flag"
	"time"

	"WardenConsumer/core"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var (
	envFile string
)

func init() {
	flag.StringVar(&envFile, "e", "", "env file")
}

func main() {
	flag.Parse()

	if envFile != "" {
		log.Infof("Initializing env..., envfile[%s]", envFile)
		err := godotenv.Load(envFile)
		if err != nil {
			log.Fatalf("Fail initialize env: %v", err)
			return
		}
		log.Info("Successfully initialize env")
	}

	ctx, err := core.NewConsumerContext()
	if err != nil {
		log.Fatal(err)
	}

	ctx.Init()

	go ctx.ConsumeOnboardTxn()
	go ctx.ConsumeSettings()

	for {
		time.Sleep(time.Second * 10)
	}
}
