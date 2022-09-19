package main

import (
	"flag"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"EnclaveProxy/server"
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
		log.Fatalf("Fail initialize env: %v", err)
		return
	}
	log.Info("Successfully initialize env")

	server.NewServer()

}