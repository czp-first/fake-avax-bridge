package main

import (
	"flag"

	"github.com/czp-first/fake-avax-bridge/WardenServer/core"

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
		}
		log.Info("Successfully initialize env")
	}

	ctx := core.NewWardenContext()

	err := ctx.Init()
	if err != nil {
		log.Fatal(err)
		return
	}

	go ctx.MonitorOnboard()
	log.Info("monitoring onboard...")
	go ctx.ConfirmOnboard()
	log.Info("confirm onboard...")

	log.Infoln("warden node successfully starts")

	core.NewServer(ctx)

}
