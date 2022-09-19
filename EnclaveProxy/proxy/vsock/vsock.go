package vsock_cli

import (
	"github.com/mdlayher/vsock"
	log "github.com/sirupsen/logrus"
)

func Req(in, out interface{}) {
	var num uint32 = 10
	var port uint32 = 5000
	c, err := vsock.Dial(num, port, &vsock.Config{})
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()

	if _, err := c.Write([]byte("hello enclave")); err != nil {
		log.Fatal(err)
	}

	b := make([]byte, 16)
	n, err := c.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(b[:n]))

}
