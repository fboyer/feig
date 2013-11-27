package fetcp

import (
	// "fmt"
	"bytes"
	"testing"
)

var hostAddr string = "192.168.20.112"
var port int = 4001

func TestFETCP_Connect(t *testing.T) {
	if socketHnd, _ := FETCP_Connect(&([]byte(hostAddr))[0], uint32(port)); socketHnd < 0 {
		t.Errorf("FETCP_Connect(%v, %v) = %v", hostAddr, port, socketHnd)
	}
}

func TestFETCP_Disconnect(t *testing.T) {
	if socketHnd, _ := FETCP_Connect(&([]byte(hostAddr))[0], uint32(port)); socketHnd < 0 {
		t.Errorf("FETCP_Connect(%v, %v) = %v\n", hostAddr, port, socketHnd)
	} else {
		if result, _ := FETCP_Disconnect(socketHnd); result != 0 {
			t.Errorf("FETCP_Disconnect(%v) = %v", socketHnd, result)
		}
	}
}

func TestFETCP_Detect(t *testing.T) {
	if result, _ := FETCP_Detect(&([]byte(hostAddr))[0], port); result != 0 {
		t.Errorf("FETCP_Detect(%v, %v) = %v", hostAddr, port, result)
	}
}

func TestFETCP_GetDLLVersion(t *testing.T) {
	ver := make([]byte, 256)
	FETCP_GetDLLVersion(&ver[0])
	n := bytes.Index(ver, []byte{0})
	version := string(ver[:n])
	if version != "02.02.00" {
		t.Errorf("FETCP_GetDLLVersion() = %v", version)
	}
}
