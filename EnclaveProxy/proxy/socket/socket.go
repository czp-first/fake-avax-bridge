package socket_cli

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	log "github.com/sirupsen/logrus"
)

func Req(in, out interface{}) {
	log.Infof("request: %+v", in)
	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%s", os.Getenv("EnclavePort")))
	if err != nil {
		log.Fatalf("conn server failed, err:%v\n", err)
	}
	defer conn.Close()

	encode := json.NewEncoder(conn)

	err = encode.Encode(&in)
	if err != nil {
		log.Fatalf("error is :%v\n", err)
	}

	decode := json.NewDecoder(conn)
	err = decode.Decode(&out)

	if err != nil {
		log.Fatalf("error is :%v\n", err)
	}
	log.Infof("response: %+v", out)
}
