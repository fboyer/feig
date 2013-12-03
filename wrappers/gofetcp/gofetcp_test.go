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
var _ = Suite(&S{0})

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
	c.Assert(version, HasLen, 8)
}

func (s *S) TestGetLastError(c *C) {
	GetSocketParam(s.socketHnd, "Unknown")
	errorCode, _, result, _ := GetLastError(s.socketHnd)
	c.Assert(result, Equals, 0)
	c.Assert(errorCode, Equals, FETCP_ERR_UNKNOWN_PARAMETER)
}

func (s *S) TestGetSocketParam(c *C) {
	value, result, _ := GetSocketParam(s.socketHnd, "Timeout")
	c.Assert(result, Equals, 0)
	c.Assert(value, Equals, "3000")
}

func (s *S) TestSetSocketParam(c *C) {
	result, _ := SetSocketParam(s.socketHnd, "Timeout", "10000")
	value, _, _ := GetSocketParam(s.socketHnd, "Timeout")
	c.Assert(result, Equals, 0)
	c.Assert(value, Equals, "10000")
}
