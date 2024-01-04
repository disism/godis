package nats

import (
	"context"
	"github.com/nats-io/nats.go/jetstream"
	"time"
)

func KVPut(name string, value []byte) error {
	nc, err := Connect()
	if err != nil {
		return err
	}
	defer nc.Drain()

	js, err := jetstream.New(nc)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	kv, err := js.CreateKeyValue(ctx, jetstream.KeyValueConfig{
		Bucket: "viper",
	})
	if err != nil {
		return err
	}

	_, err = kv.Put(ctx, name, value)
	if err != nil {
		return err
	}

	return nil
}
