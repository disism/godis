package conf

import (
	"fmt"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

const DefaultNATSAddr = "127.0.0.1:4222"

func ConnectAndReadRemoteConf() {
	if err := viper.AddRemoteProvider("nats", fmt.Sprintf("nats://%s", DefaultNATSAddr), "conf"); err != nil {
		panic(fmt.Errorf("fatal errors connect remote: %w", err))
	}
	viper.SetConfigType("toml")
	if err := viper.ReadRemoteConfig(); err != nil {
		panic(fmt.Errorf("fatal errors config file: %w", err))
	}
}
