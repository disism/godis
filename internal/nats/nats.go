package nats

import (
	"github.com/nats-io/nats.go"
	"os"
)

func Connect() (*nats.Conn, error) {
	addr := os.Getenv("NATS_ADDR")
	if addr == "" {
		addr = nats.DefaultURL
	}

	nc, err := nats.Connect(addr)
	if err != nil {
		return nil, err
	}

	return nc, nil
}
