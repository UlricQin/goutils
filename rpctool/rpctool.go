package rpctool

import (
	"time"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func DialTimeout(network, address string, timeout time.Duration) (*rpc.Client, error) {
	conn, err := net.DialTimeout(network, address, timeout)
	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), err
}

