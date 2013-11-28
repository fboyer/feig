package reader

import (
	"testing"
)

var (
	hostAddr = "192.168.23.90"
	port     = 4001
)

func TestConnect(t *testing.T) {
	const hostAddr = "192.168.23.90"
	if socketHnd, _ := Connect(hostAddr, 4001); socketHnd < 0 {
		t.Errorf("Connect(%v) = %v", hostAddr, socketHnd)
	}
}

// func TestDisconnect(t *testing.T) {
// To-Do
// }
