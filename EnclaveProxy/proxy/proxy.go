package proxy

import (
	"os"

	socket_cli "github.com/czp-first/fake-avax-bridge/EnclaveProxy/proxy/socket"
	vsock_cli "github.com/czp-first/fake-avax-bridge/EnclaveProxy/proxy/vsock"
)

func Req(in, out interface{}) {
	mode := os.Getenv("ProxyMode")
	if mode == "vsock" {
		vsock_cli.Req(in, out)
	} else {
		socket_cli.Req(in, out)
	}
}
