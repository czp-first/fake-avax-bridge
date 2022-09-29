package main

import (
	"flag"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/czp-first/fake-avax-bridge/EnclaveProxy/core"
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
		log.Info("Initializing env...")
		err := godotenv.Load(envFile)
		if err != nil {
			log.Fatalf("Fail initialize env: %v", err)
			return
		}
		log.Info("Successfully initialize env")
	}

	ctx := core.NewEnclaveProxyContext()
	err := ctx.Init()
	if err != nil {
		log.Fatal(err)
		return
	}
	core.NewServer(ctx)

}
