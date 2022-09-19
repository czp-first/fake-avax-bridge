package proxy

import (
	"os"

	socket_cli "EnclaveProxy/proxy/socket"
	vsock_cli "EnclaveProxy/proxy/vsock"
)

func Req(in, out interface{}) {
	mode := os.Getenv("ProxyMode")
	if mode == "vsock" {
		vsock_cli.Req(in, out)
	} else {
		socket_cli.Req(in, out)
	}
}
