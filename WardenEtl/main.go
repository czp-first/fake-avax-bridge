package main

import (
	"flag"
	"time"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"WardenEtl/core"
)

var (
	envFile string
)

func init() {
	flag.StringVar(&envFile, "e", "dev.env", "env file")
}

func main() {

	flag.Parse()

	log.Info("Initializing env...")

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Fail initialize env: %v\n", err)
		return
	}
	log.Info("Successfully initialize env")

	ctx, err := core.NewWardenContext()
	if err != nil {
		panic("create context error!")
	}

	ctx.Init()

	go ctx.SeeEthBlock()

	// go ctx.SeeDxchainBlock()
	// go ctx.ConsumeOffboardTxn()

	// go ctx.MonitorOffboard()
	// go ctx.ConfirmOffboard()

	for {
		time.Sleep(5 * time.Second)
	}

}
