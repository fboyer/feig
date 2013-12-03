package gofeisc

import (
	"github.com/fboyer/gofeig/wrappers/gofetcp/gofetcp"
	gc "launchpad.net/gocheck"
	jc "launchpad.net/juju-core/testing/checkers"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type FeiscSuite struct {
	readerHnd, socketHnd int
}

func (s *FeiscSuite) SetUpSuite(c *gc.C) {
	s.socketHnd, _ = gofetcp.Connect("192.168.20.112", 4001)
}

func (s *FeiscSuite) SetUpTest(c *gc.C) {
	s.readerHnd, _ = NewReader(s.socketHnd)
}

func (s *FeiscSuite) TearDownTest(c *gc.C) {
	if s.readerHnd != 0 {
		DeleteReader(s.readerHnd)
	}
}

func (s *FeiscSuite) TearDownSuite(c *gc.C) {
	if s.socketHnd != 0 {
		gofetcp.Disconnect(s.socketHnd)
	}
}

func (s *FeiscSuite) TestNewReader(c *gc.C) {
	c.Assert(s.readerHnd, jc.GreaterThan, 0)
}

func (s *FeiscSuite) TestDeleteReader(c *gc.C) {
	result, _ := DeleteReader(s.readerHnd)
	c.Assert(result, gc.Equals, 0)
	s.readerHnd = 0
}

func (s *FeiscSuite) TestGetDllVersion(c *gc.C) {
	version := GetDllVersion()
	c.Assert(version, gc.HasLen, 7)
}

func (s *FeiscSuite) TestGetLastError(c *gc.C) {
	readerHnd, _ := NewReader(-1)
	errorCode, _, result, _ := GetLastError(readerHnd)
	c.Assert(result, gc.Equals, 0)
	c.Assert(errorCode, gc.Equals, FEISC_ERR_HND_IS_NEGATIVE)
}

func (s *FeiscSuite) TestGetLastStatus(c *gc.C) {
	status, _, _ := GetLastStatus(s.readerHnd)
	c.Assert(status, gc.Equals, 0)
}

func (s *FeiscSuite) TestGetReaderParam(c *gc.C) {
	value, result, _ := GetReaderParam(s.readerHnd, "FrameSupport")
	c.Assert(result, Equals, 0)
	c.Assert(value, Not(Equals), "ADVANCED")
}

func (s *FeiscSuite) TestSetReaderParam(c *gc.C) {
	result, _ := SetReaderParam(s.readerHnd, "FrameSupport", "ADVANCED")
	value, _, _ := GetReaderParam(s.readerHnd, "FrameSupport")
	c.Assert(result, Equals, 0)
	c.Assert(value, Equals, "ADVANCED")
}

func (s *FeiscSuite) TestResetCpu(c *gc.C) {
	status, _ := ResetCpu(s.readerHnd)
	c.Assert(status, Equals, 0)
}

func (s *FeiscSuite) TestGetSoftVersion(c *gc.C) {
	version, status, _ := GetSoftVersion(s.readerHnd, 255, 1)
	c.Assert(status, Equals, 0)
	c.Assert(version, HasLen, 10)
}

func (s *FeiscSuite) TestResetRf(c *gc.C) {
	status, _ := ResetRf(s.readerHnd, 255)
	c.Assert(status, Equals, 0)
}

func (s *FeiscSuite) TestSetRfOnOff(c *gc.C) {
	status, _ := SetRfOnOff(s.readerHnd, 255, 1)
	c.Assert(status, Equals, 0)
	status, _ = SetRfOnOff(s.readerHnd, 255, 0)
	c.Assert(status, Equals, 0)
}

// func (s *FeiscSuite) TestReadConfBlock(c C*) {
// }

// func (s *FeiscSuite) TestWriteConfBlock(c C*) {
// }

// func (s *FeiscSuite) TestSaveConfBlock(c C+) {
// }

// func (s *FeiscSuite) TestResetConfBlock(c C*) {
// }

func (s *FeiscSuite) TestSendIsoCmd(c *gc.C) {
	respData, respLen, result, _ := SendIsoCmd(s.readerHnd, 255, "0100", 4, 1)
	c.Assert(status, Equals, 0)
}
