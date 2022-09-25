package main

import (
	"flag"
	"time"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/czp-first/fake-avax-bridge/WardenEtl/core"
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

	ctx, err := core.NewWardenContext()
	if err != nil {
		panic("create context error!")
	}

	ctx.Init()

	go ctx.SeeFromChainBlock()

	// go ctx.SeeDxchainBlock()

	for {
		time.Sleep(5 * time.Second)
	}

}
