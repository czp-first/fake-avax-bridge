package socket_cli

import (
	"encoding/json"
	"net"
	"os"

	log "github.com/sirupsen/logrus"
)

func Req(in, out interface{}) {
	log.Infof("%+v", in)
	conn, err := net.Dial("tcp", os.Getenv("EnclaveURL"))
	if err != nil {
		log.Fatalf("conn server failed, err:%v\n", err)
	}
	defer conn.Close()

	encode := json.NewEncoder(conn)

	err = encode.Encode(&in)
	if err != nil {
		log.Fatalf("encode req error is :%v\n", err)
	}

	decode := json.NewDecoder(conn)
	err = decode.Decode(&out)

	if err != nil {
		log.Fatalf("encode resp error is :%v\n", err)
	}
	log.Infof("%+v", out)
}
