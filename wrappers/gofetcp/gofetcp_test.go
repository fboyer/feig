package gofetcp

import (
	gc "launchpad.net/gocheck"
	jc "launchpad.net/juju-core/testing/checkers"
	"testing"
)

func Test(t *testing.T) {
	gc.TestingT(t)
}

type FetcpSuite struct {
	socketHnd int
}

var _ = gc.Suite(&FetcpSuite{0})

// To be tested with a live setup
func (s *FetcpSuite) SetUpTest(c *gc.C) {
	s.socketHnd, _ = Connect("192.168.23.253", 4001)
}

func (s *FetcpSuite) TearDownTest(c *gc.C) {
	if s.socketHnd != 0 {
		Disconnect(s.socketHnd)
	}
}

func (s *FetcpSuite) TestConnect(c *gc.C) {
	c.Assert(s.socketHnd, jc.GreaterThan, 0)
}

func (s *FetcpSuite) TestDisconnect(c *gc.C) {
	result, _ := Disconnect(s.socketHnd)
	c.Assert(result, gc.Equals, 0)
	s.socketHnd = 0
}

func (s *FetcpSuite) TestGetDllVersion(c *gc.C) {
	version := GetDllVersion()
	c.Assert(version, gc.HasLen, 8)
}

func (s *FetcpSuite) TestGetLastError(c *gc.C) {
	GetSocketParam(s.socketHnd, "Unknown")
	errorCode, _, result, _ := GetLastError(s.socketHnd)
	c.Assert(result, gc.Equals, 0)
	c.Assert(errorCode, gc.Equals, FETCP_ERR_UNKNOWN_PARAMETER)
}

func (s *FetcpSuite) TestGetSocketParam(c *gc.C) {
	value, result, _ := GetSocketParam(s.socketHnd, "Timeout")
	c.Assert(result, gc.Equals, 0)
	c.Assert(value, gc.Equals, "3000")
}

func (s *FetcpSuite) TestSetSocketParam(c *gc.C) {
	result, _ := SetSocketParam(s.socketHnd, "Timeout", "10000")
	value, _, _ := GetSocketParam(s.socketHnd, "Timeout")
	c.Assert(result, gc.Equals, 0)
	c.Assert(value, gc.Equals, "10000")
}
