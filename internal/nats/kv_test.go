package nats

import "testing"

func TestKVPut(t *testing.T) {
	var v = []byte(`
[server]
addr = ""

[redis]
addr = "10.143.24.84:6379"
password = ""

[jwt]
secret = "1234567890.disism.com.saikan.0987654321"

[ipfs]
addr = "http://127.0.0.1:5001"
`)
	if err := KVPut("conf", v); err != nil {
		t.Errorf("KVPut() error = %v", err)
		return
	}
}
