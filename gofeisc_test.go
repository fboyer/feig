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

// This test does not require a live environment to be run
type FeiscSuite struct{}

var _ = gc.Suite(&FeiscSuite{})

func (f *FeiscSuite) TestGetDllVersion(c *gc.C) {
	version := GetDllVersion()
	c.Assert(version, gc.HasLen, 8)
}

// These tests require a live environment with at least 1 RFID tag present on the antenna
type LiveFeiscSuite struct {
	readerHnd, socketHnd int
}

var _ = gc.Suite(&LiveFeiscSuite{0, 0})

func (l *LiveFeiscSuite) SetUpSuite(c *gc.C) {
	if *live {
		c.Skip("-live not provided")
	}
	l.socketHnd, _ = gofetcp.Connect("192.168.23.253", 4001)
	gofetcp.SetSocketParam(l.socketHnd, "Timeout", "10000")
	gofetcp.SetSocketParam(l.socketHnd, "CharTimeout", "100")
}

func (l *LiveFeiscSuite) SetUpTest(c *gc.C) {
	l.readerHnd, _ = NewReader(l.socketHnd)
	SetReaderParam(l.readerHnd, "FrameSupport", "ADVANCED")
	time.Sleep(1 * time.Second)
}

func (l *LiveFeiscSuite) TearDownTest(c *gc.C) {
	if l.readerHnd != 0 {
		DeleteReader(l.readerHnd)
	}
}

func (l *LiveFeiscSuite) TearDownSuite(c *gc.C) {
	if l.socketHnd != 0 {
		gofetcp.Disconnect(l.socketHnd)
	}
}

func (l *LiveFeiscSuite) TestNewReader(c *gc.C) {
	c.Assert(l.readerHnd, jc.GreaterThan, 0)
}

func (l *LiveFeiscSuite) TestDeleteReader(c *gc.C) {
	result, _ := DeleteReader(l.readerHnd)
	c.Assert(result, gc.Equals, 0)
	l.readerHnd = 0
}

func (l *LiveFeiscSuite) TestGetLastError(c *gc.C) {
	GetReaderParam(l.readerHnd, "UnknownParameter")
	errorCode, _, result, _ := GetLastError(l.readerHnd)
	c.Assert(result, gc.Equals, 0)
	c.Assert(errorCode, gc.Equals, FEISC_ERR_UNKNOWN_PARAMETER)
}

func (l *LiveFeiscSuite) TestGetLastState(c *gc.C) {
	result, _, _ := GetLastState(l.readerHnd)
	c.Assert(result, gc.Equals, 0)
}

func (l *LiveFeiscSuite) TestGetReaderParam(c *gc.C) {
	value, result, _ := GetReaderParam(l.readerHnd, "FrameSupport")
	c.Assert(result, gc.Equals, 0)
	c.Assert(value, gc.Equals, "ADVANCED")
}

func (l *LiveFeiscSuite) TestSetReaderParam(c *gc.C) {
	result, _ := SetReaderParam(l.readerHnd, "FrameSupport", "STANDARD")
	value, _, _ := GetReaderParam(l.readerHnd, "FrameSupport")
	c.Assert(result, gc.Equals, 0)
	c.Assert(value, gc.Equals, "STANDARD")
}

func (l *LiveFeiscSuite) TestResetCpu(c *gc.C) {
	result, _ := ResetCpu(l.readerHnd, 255)
	c.Assert(result, gc.Equals, 0)
}

func (l *LiveFeiscSuite) TestGetSoftVersion(c *gc.C) {
	version, result, _ := GetSoftVersion(l.readerHnd, 255, 1)
	c.Assert(result, gc.Equals, 0)
	c.Assert(version, gc.HasLen, 14)
}

func (l *LiveFeiscSuite) TestResetRf(c *gc.C) {
	result, _ := ResetRf(l.readerHnd, 255)
	c.Assert(result, gc.Equals, 0)
}

func (l *LiveFeiscSuite) TestSetRfOnOff(c *gc.C) {
	result, _ := SetRfOnOff(l.readerHnd, 255, 1)
	c.Assert(result, gc.Equals, 0)
	result, _ = SetRfOnOff(l.readerHnd, 255, 0)
	c.Assert(result, gc.Equals, 0)
}

// func (l *LiveFeiscSuite) TestReadConfBlock(c C*) {
// }

// func (l *LiveFeiscSuite) TestWriteConfBlock(c C*) {
// }

// func (l *LiveFeiscSuite) TestSaveConfBlock(c C+) {
// }

// func (l *LiveFeiscSuite) TestResetConfBlock(c C*) {
// }

func (l *LiveFeiscSuite) TestSendIsoCmd(c *gc.C) {
	_, _, result, _ := SendIsoCmd(l.readerHnd, 255, "0100", 4, 1)
	c.Assert(result, gc.Equals, 0)
}
