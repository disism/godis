package conf

import (
	"github.com/spf13/viper"
	"testing"
)

func TestConfig(t *testing.T) {
	addr := viper.GetString("redis.addr")
	t.Logf("addr: %s", addr)
}
