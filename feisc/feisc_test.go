package feisc

import . "gopkg.in/check.v1"

type FeiscSuite struct {
	readerHnd int
	socketHnd int
}

var _ = Suite(&FeiscSuite{})

func (f *FeiscSuite) SetUpSuite(c *C) {
	if !*live {
		c.Skip("-live not provided")
	}
	f.socketHnd = Connect(*hostAddr, *port)
	if f.socketHnd < 0 {
		c.Skip("Unable to get a new socket handle")
	}
}

func (f *FeiscSuite) SetUpTest(c *C) {
	f.readerHnd = NewReader(f.socketHnd)
	if f.readerHnd < 0 {
		c.Skip("Unable to get a new reader handle")
	}
	result := SetReaderParam(f.readerHnd, "FrameSupport", "ADVANCED")
	if result < 0 {
		c.Skip("Unable to set the FrameSupport to ADVANCED")
	}
}

func (f *FeiscSuite) TearDownTest(c *C) {
	if f.readerHnd > 0 {
		DeleteReader(f.readerHnd)
	}
}

func (f *FeiscSuite) TearDownSuite(c *C) {
	if f.socketHnd > 0 {
		Disconnect(f.socketHnd)
	}
}

func (f *FeiscSuite) TestNewReader(c *C) {
	c.Assert(f.readerHnd > 0, Equals, true)
}

func (f *FeiscSuite) TestDeleteReader(c *C) {
	c.Assert(DeleteReader(f.readerHnd), Equals, 0)
}

func (f *FeiscSuite) TestGetIscDllVersion(c *C) {
	c.Assert(GetIscDllVersion(), HasLen, 8)
}

// TODO: Issue with the errorCode return value
func (f *FeiscSuite) TestGetIscLastError(c *C) {
	SetReaderParam(f.readerHnd, "Unknown", "Unknown")
	_, errText, result := GetIscLastError(f.readerHnd)
	c.Assert(result, Equals, 0)
	//c.Assert(errCode, Equals, FEISC_ERR_UNKNOWN_PARAMETER)
	c.Assert(errText, Equals, "FEISC: (-4050) unknown parameter")
}

func (f *FeiscSuite) TestGetLastState(c *C) {
	stateText, result := GetLastState(f.readerHnd)
	c.Assert(result, Equals, 0)
	c.Assert(stateText, Equals, "OK")
}

func (f *FeiscSuite) TestGetReaderParam(c *C) {
	value, result := GetReaderParam(f.readerHnd, "FrameSupport")
	c.Assert(result, Equals, 0)
	c.Assert(value, Equals, "ADVANCED")
}

func (f *FeiscSuite) TestSetReaderParam(c *C) {
	result := SetReaderParam(f.readerHnd, "FrameSupport", "STANDARD")
	c.Assert(result, Equals, 0)
	value, result := GetReaderParam(f.readerHnd, "FrameSupport")
	c.Assert(result, Equals, 0)
	c.Assert(value, Equals, "STANDARD")
}

func (f *FeiscSuite) TestResetCpu(c *C) {
	result := ResetCpu(f.readerHnd, 255)
	c.Assert(result, Equals, 0)
}

func (f *FeiscSuite) TestGetSoftVersion(c *C) {
	version, result := GetSoftVersion(f.readerHnd, 255, 0)
	c.Assert(result, Equals, 0)
	c.Assert(version, HasLen, 8)
}
