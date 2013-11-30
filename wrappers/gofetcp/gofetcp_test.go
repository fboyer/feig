package gofetcp

import (
	"testing"
)

var hostAddr string = "192.168.20.112"
var port int = 4001
var dllVersion = "02.02.00"

func TestConnect(t *testing.T) {
	if socketHnd, err := Connect(hostAddr, port); socketHnd < 0 {
		t.Errorf("Connect(%v, %v) = %v", hostAddr, port, socketHnd)
	}
	Disconnect(socketHnd)
}

func TestDisconnect(t *testing.T) {
	socketHnd, err := Connect(hostAddr, port)
	if result, _ := Disconnect(socketHnd); result != 0 {
		t.Errorf("Disconnect(%v) = %v", socketHnd, result)
	}
}

func TestGetDllVersion(t *testing.T) {
	version := GetDllVersion()
	if version != dllVersion {
		t.Errorf("GetDllVersion() = %v, wants %v", version, dllVersion)
	}
}

func TestGetLastError(t *testing.T) {
	socketHnd, _ := Connect("127.0.0.1", port)
	if errorCode, errorText, result, err := GetLastError(socketHnd); errorCode != FETCP_ERR_TIMEOUT {
		t.Errorf("GetLastError(%v) = %v, wants %v", socketHnd, errorCode, FETCP_ERR_TIMEOUT)
	}
	Disconnect(socketHnd)
}

func TestGetSocketParam(t *testing.T) {
	socketHnd, _ := Connect(hostAddr, port)
	if value, result, err := GetSocketParam(socketHnd, "Timeout"); value != "3000" {
		t.Errorf("GetSocketParam(%v, \"Timeout\") = %v, wants 3000", socketHnd, value)
	}
	Disconnect(socketHnd)
}

func TestSetSocketParam(t *testing.T) {
	socketHnd, _ := Connect(hostAddr, port)
	if result, err := SetSocketParam(socketHnd, "Timeout", "10000"); result != 0 {
		t.Errorf("SetSocketParam(%v, \"Timeout\", \"10000\") = %v, wants 0", socketHnd, result)
	} else {
		if value, result, err := GetSocketParam(socketHnd, "Timeout"); value != "10000" {
			t.Errorf("GetSocketParam(%v, \"Timeout\") = %v, expected 10000", socketHnd, value)
		}
	}
}
