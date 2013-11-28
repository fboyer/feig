package reader

import (
	"testing"
)

var (
	r = Reader{hostAddr: "192.168.20.112", port: 4001, busAddr: 255}
)

func TestConnect(t *testing.T) {
	if err := r.Connect(); err != nil {
		t.Errorf("r.Connect() = %v", r.socketHnd)
	}
	r.Disconnect()
}

func TestDisconnect(t *testing.T) {
	r.Connect()
	if status, err := r.Disconnect(); err != nil {
		t.Errorf("r.Disconnect() = %v", status)
	}
}
