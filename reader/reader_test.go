package reader

import (
	"testing"
)

var (
	r = Reader{hostAddr: "192.168.20.112", port: 4001, busAddr: 255}
)

func TestReaderConnect(t *testing.T) {
	if err := Connect(hostAddr, 4001); err != nil {
		t.Errorf("Connect(%v) = %v", r.hostAddr, r.socketHnd)
	}
}

// func TestDisconnect(t *testing.T) {
// To-Do
// }
