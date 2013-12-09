package gofeig

import (
	gc "launchpad.net/gocheck"
	jc "launchpad.net/juju-core/testing/checkers"
	"testing"
)

func Test(t *testing.T) {
	gc.TestingT(t)
}

var live = flag.Bool("live", false, "Include live tests")
var hostAddr = flag.String("hostAddr", "", "Define the host address for the RFID reader")
var port = flag.Int("port", 4001, "Define the port number for the RFID reader")

// This test does not require a live environment to be run
type FetcpSuite struct{}

var _ = gc.Suite(&FetcpSuite{})

func (f *FetcpSuite) TestGetDllVersion(c *gc.C) {
	version := GetDllVersion()
	c.Assert(version, gc.HasLen, 8)
}

// These tests require a live environment
type LiveFetcpSuite struct {
	socketHnd int
}

var _ = gc.Suite(&LiveFetcpSuite{0})

func (l *LiveFetcpSuite) SetUpSuite() {
	if *live {
		c.Skip("-live not provided")
	}
}

func (l *LiveFetcpSuite) SetUpTest(c *gc.C) {
	l.socketHnd, _ = Connect("192.168.23.253", 4001)
}

func (l *LiveFetcpSuite) TearDownTest(c *gc.C) {
	if l.socketHnd != 0 {
		Disconnect(l.socketHnd)
	}
}

func (l *LiveFetcpSuite) TestConnect(c *gc.C) {
	c.Assert(l.socketHnd, jc.GreaterThan, 0)
}

func (l *LiveFetcpSuite) TestDisconnect(c *gc.C) {
	result, _ := Disconnect(l.socketHnd)
	c.Assert(result, gc.Equals, 0)
	l.socketHnd = 0
}

func (l *LiveFetcpSuite) TestGetLastError(c *gc.C) {
	GetSocketParam(l.socketHnd, "Unknown")
	errorCode, _, result, _ := GetLastError(l.socketHnd)
	c.Assert(result, gc.Equals, 0)
	c.Assert(errorCode, gc.Equals, FETCP_ERR_UNKNOWN_PARAMETER)
}

func (l *LiveFetcpSuite) TestGetSocketParam(c *gc.C) {
	value, result, _ := GetSocketParam(l.socketHnd, "Timeout")
	c.Assert(result, gc.Equals, 0)
	c.Assert(value, gc.Equals, "3000")
}

func (l *LiveFetcpSuite) TestSetSocketParam(c *gc.C) {
	result, _ := SetSocketParam(l.socketHnd, "Timeout", "10000")
	value, _, _ := GetSocketParam(l.socketHnd, "Timeout")
	c.Assert(result, gc.Equals, 0)
	c.Assert(value, gc.Equals, "10000")
}
