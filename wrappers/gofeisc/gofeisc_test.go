package gofeisc

import (
	"github.com/fboyer/gofeig/wrappers/gofetcp/gofetcp"
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type S struct {
	readerHnd, socketHnd int
}

func (s *S) SetUpSuite(c *C) {
	s.socketHnd, _ = gofetcp.Connect("192.168.20.112", 4001)
}

func (s *S) SetUpTest(c *C) {
	s.readerHnd, _ = NewReader(s.socketHnd)
}

func (s *S) TearDownTest(c *C) {
	if s.readerHnd != 0 {
		DeleteReader(s.readerHnd)
	}
}

func (s *S) TearDownSuite(c *C) {
	if s.socketHnd != 0 {
		gofetcp.Disconnect(s.socketHnd)
	}
}

func (s *S) TestNewReader(c *C) {
	c.Assert(s.readerHnd, Equals, 1)
}

func (s *S) TestDeleteReader(c *C) {
	result, _ := DeleteReader(s.readerHnd)
	c.Assert(result, Equals, 0)
	s.readerHnd = 0
}

func (s *S) TestGetDllVersion(c *C) {
	version := GetDllVersion()
	c.Assert(version, HasLen, 7)
}

func (s *S) TestGetLastError(c *C) {
	readerHnd, _ := NewReader(-1)
	errorCode, _, result, _ := GetLastError(readerHnd)
	c.Assert(result, Equals, 0)
	c.Assert(errorCode, Equals, FEISC_ERR_HND_IS_NEGATIVE)
}

func (s *S) TestGetLastStatus(c *C) {
	status, _, _ := GetLastStatus(s.readerHnd)
	c.Assert(status, Equals, 0)
}

func (s *S) TestGetReaderParam(c *C) {
	value, result, _ := GetReaderParam(s.readerHnd, "FrameSupport")
	c.Assert(result, Equals, 0)
	c.Assert(value, Not(Equals), "ADVANCED")
}

func (s *S) TestSetReaderParam(c *C) {
	result, _ := SetReaderParam(s.readerHnd, "FrameSupport", "ADVANCED")
	value, _, _ := GetReaderParam(s.readerHnd, "FrameSupport")
	c.Assert(result, Equals, 0)
	c.Assert(value, Equals, "ADVANCED")
}

func (s *S) TestResetCpu(c C*) {
	status, _ := ResetCpu(s.readerHnd)
	c.Assert(status, Equals, 0)
}

func (s *S) TestGetSoftVersion(c *C) {
	version, status, _ := GetSoftVersion(s.readerHnd, 255, 1)
	c.Assert(status, Equals, 0)
	c.Assert(version, HasLen, 10)
}

func (s *S) TestResetRf(c C*) {
	status, _ := ResetRf(s.readerHnd, 255)
	c.Assert(status, Equals, 0)
}

func (s *S) TestSetRfOnOff(c C*) {
	status, _ := SetRfOnOff(s.readerHnd, 255, 1)
	c.Assert(status, Equals, 0)
	status, _ = SetRfOnOff(s.readerHnd, 255, 0)
	c.Assert(status, Equals, 0)
}

// func (s *S) TestReadConfBlock(c C*) {
// }

// func (s *S) TestWriteConfBlock(c C*) {
// }

// func (s *S) TestSaveConfBlock(c C+) {
// }

// func (s *S) TestResetConfBlock(c C*) {
// }

func (s *S) TestSendIsoCmd(c C*) {
	respData, respLen, result, _ := SendIsoCmd(s.readerHnd, 255, "0100", 4, 1)
	c.Assert(status, Equals, 0)
}
