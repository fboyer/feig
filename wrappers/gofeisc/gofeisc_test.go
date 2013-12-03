package gofeisc

import (
	"github.com/fboyer/gofeig/wrappers/gofetcp"
	gc "launchpad.net/gocheck"
	jc "launchpad.net/juju-core/testing/checkers"
	"testing"
	"time"
)

func Test(t *testing.T) {
	gc.TestingT(t)
}

type FeiscSuite struct {
	readerHnd, socketHnd int
}

var _ = gc.Suite(&FeiscSuite{0, 0})

// To be tested with a live setup with at least 1 RFID tag present on the antenna
func (s *FeiscSuite) SetUpSuite(c *gc.C) {
	s.socketHnd, _ = gofetcp.Connect("192.168.23.253", 4001)
	gofetcp.SetSocketParam(s.socketHnd, "Timeout", "10000")
	gofetcp.SetSocketParam(s.socketHnd, "CharTimeout", "100")
}

func (s *FeiscSuite) SetUpTest(c *gc.C) {
	s.readerHnd, _ = NewReader(s.socketHnd)
	SetReaderParam(s.readerHnd, "FrameSupport", "ADVANCED")
	time.Sleep(1 * time.Second)
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
	c.Assert(version, gc.HasLen, 8)
}

func (s *FeiscSuite) TestGetLastError(c *gc.C) {
	GetReaderParam(s.readerHnd, "UnknownParameter")
	errorCode, _, result, _ := GetLastError(s.readerHnd)
	c.Assert(result, gc.Equals, 0)
	c.Assert(errorCode, gc.Equals, FEISC_ERR_UNKNOWN_PARAMETER)
}

func (s *FeiscSuite) TestGetLastState(c *gc.C) {
	result, _, _ := GetLastState(s.readerHnd)
	c.Assert(result, gc.Equals, 0)
}

func (s *FeiscSuite) TestGetReaderParam(c *gc.C) {
	value, result, _ := GetReaderParam(s.readerHnd, "FrameSupport")
	c.Assert(result, gc.Equals, 0)
	c.Assert(value, gc.Equals, "ADVANCED")
}

func (s *FeiscSuite) TestSetReaderParam(c *gc.C) {
	result, _ := SetReaderParam(s.readerHnd, "FrameSupport", "STANDARD")
	value, _, _ := GetReaderParam(s.readerHnd, "FrameSupport")
	c.Assert(result, gc.Equals, 0)
	c.Assert(value, gc.Equals, "STANDARD")
}

func (s *FeiscSuite) TestResetCpu(c *gc.C) {
	result, _ := ResetCpu(s.readerHnd, 255)
	c.Assert(result, gc.Equals, 0)
}

func (s *FeiscSuite) TestGetSoftVersion(c *gc.C) {
	version, result, _ := GetSoftVersion(s.readerHnd, 255, 1)
	c.Assert(result, gc.Equals, 0)
	c.Assert(version, gc.HasLen, 14)
}

func (s *FeiscSuite) TestResetRf(c *gc.C) {
	result, _ := ResetRf(s.readerHnd, 255)
	c.Assert(result, gc.Equals, 0)
}

func (s *FeiscSuite) TestSetRfOnOff(c *gc.C) {
	result, _ := SetRfOnOff(s.readerHnd, 255, 1)
	c.Assert(result, gc.Equals, 0)
	result, _ = SetRfOnOff(s.readerHnd, 255, 0)
	c.Assert(result, gc.Equals, 0)
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
	_, _, result, _ := SendIsoCmd(s.readerHnd, 255, "0100", 4, 1)
	c.Assert(result, gc.Equals, 0)
}
