package gofetcp

import (
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type S struct {
	socketHnd int
}

func (s *S) SetUpTest(c *C) {
	s.socketHnd, _ = Connect("192.168.20.112", 4001)
}

func (s *S) TearDownTest(c *C) {
	if s.socketHnd != 0 {
		Disconnect(s.socketHnd)
	}
}

func (s *S) TestConnect(c *C) {
	c.Assert(s.socketHnd, Equals, 1)
}

func (s *S) TestDisconnect(c *C) {
	result, _ := Disconnect(s.socketHnd)
	c.Assert(result, Equals, 0)
	s.socketHnd = 0
}

func (s *S) TestGetDllVersion(c *C) {
	version := GetDllVersion()
	c.Assert(version, HasLen, 7)
}

func (s *S) TestGetLastError(c *C) {
	socketHnd, _ := Connect("127.0.0.1", 4001)
	errorCode, _, result, _ := GetLastError(socketHnd)
	c.Assert(result, Equals, 0)
	c.Assert(errorCode, Equals, FETCP_ERR_TIMEOUT)
}

func (s *S) TestGetSocketParam(c *C) {
	value, result, _ := GetSocketParam(socketHnd, "Timeout")
	c.Assert(result, Equals, 0)
	c.Assert(value, Equals, "3000")
}

func (s *S) TestSetSocketParam(c *C) {
	result, _ := SetSocketParam(socketHnd, "Timeout", "10000")
	value, _, _ := GetSocketParam(socketHnd, "Timeout")
	c.Assert(result, Equals, 0)
	c.Assert(value, Equals, "10000")
}
