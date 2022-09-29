package socket_cli

import (
	"encoding/json"
	"net"
	"os"

	log "github.com/sirupsen/logrus"
)

func Req(in, out interface{}) {
	log.Infof("request: %+v", in)
	conn, err := net.Dial("tcp", os.Getenv("EnclaveURL"))
	if err != nil {
		log.Fatalf("conn server failed, err:%v", err)
	}
	defer conn.Close()

	encode := json.NewEncoder(conn)

	err = encode.Encode(&in)
	if err != nil {
		log.Fatalf("error is :%v", err)
	}

	decode := json.NewDecoder(conn)
	err = decode.Decode(&out)

	if err != nil {
		log.Fatalf("error is :%v", err)
	}
	log.Infof("response: %+v", out)
}
