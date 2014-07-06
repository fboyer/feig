package fetcp

import . "gopkg.in/check.v1"

type FetcpSuite struct {
	socketHnd int
}

var _ = Suite(&FetcpSuite{})

func (f *FetcpSuite) SetUpSuite(c *C) {
	if !*live {
		c.Skip("-live not provided")
	}
}

func (f *FetcpSuite) SetUpTest(c *C) {
	f.socketHnd = Connect(*hostAddr, *port)
	if f.socketHnd < 0 {
		c.Skip("Unable to get a new socket handle")
	}
}

func (f *FetcpSuite) TearDownTest(c *C) {
	if f.socketHnd > 0 {
		Disconnect(f.socketHnd)
	}
}

func (f *FetcpSuite) TestConnect(c *C) {
	c.Assert(f.socketHnd > 0, Equals, true)
}

func (f *FetcpSuite) TestDisconnect(c *C) {
	c.Assert(Disconnect(f.socketHnd), Equals, 0)
}

func (f *FetcpSuite) TestGetTcpDllVersion(c *C) {
	c.Assert(GetTcpDllVersion(), HasLen, 8)
}

// TODO: Issue with the errorCode return value
func (f *FetcpSuite) TestGetTcpLastError(c *C) {
	SetSocketParam(f.socketHnd, "Unknown", "Unknown")
	_, errText, result := GetTcpLastError(f.socketHnd)
	c.Assert(result, Equals, 0)
	//c.Assert(errCode, Equals, FETCP_ERR_UNKNOWN_PARAMETER)
	c.Assert(errText, Equals, "FETCP: (-1250) unknown parameter")
}

func (f *FetcpSuite) TestGetSocketParam(c *C) {
	value, result := GetSocketParam(f.socketHnd, "Timeout")
	c.Assert(result, Equals, 0)
	c.Assert(value, Equals, "3000")
}

func (f *FetcpSuite) TestSetSocketParam(c *C) {
	result := SetSocketParam(f.socketHnd, "Timeout", "10000")
	c.Assert(result, Equals, 0)

	value, result := GetSocketParam(f.socketHnd, "Timeout")
	c.Assert(result, Equals, 0)
	c.Assert(value, Equals, "10000")
}
