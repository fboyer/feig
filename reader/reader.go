package reader

import (
	"bytes"
	"github.com/fboyer/gofeig/feisc"
	"github.com/fboyer/gofeig/fetcp"
	"regexp"
)

type Reader struct {
	busAddr                                byte
	port, socketHnd, readerHnd             int
	hostAddr, tcpDllVersion, iscDllVersion string
}

func NewReader(hostAddr string, port int) *Reader {
	return nil
}

var (
	validIpAddressRegex = regexp.MustCompile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	validHostnameRegex  = regexp.MustCompile(`^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$`)
)

// FETCP
func (r *Reader) Connect() (err error) {
	r.socketHnd, err = fetcp.FETCP_Connect(&([]byte(r.hostAddr))[0], r.port)
	return
}

func (r *Reader) Disconnect() (status int, err error) {
	status, err = fetcp.FETCP_Disconnect(r.socketHnd)
	return
}

func (r *Reader) GetTcpDllVersion() {
	ver := make([]byte, 256)
	fetcp.FETCP_GetDLLVersion(&ver[0])
	n := bytes.Index(ver, []byte{0})
	r.tcpDllVersion = string(ver[:n])
	return
}

func (r *Reader) GetTcpLastError() (errorCode int, errorText string, result int, err error) {
	text := make([]byte, 256)
	result, err = fetcp.FETCP_GetLastError(r.socketHnd, &errorCode, &text[0])
	n := bytes.Index(text, []byte{0})
	errorText = string(text[:n])
	return
}

func (r *Reader) GetSocketParam(param string) (value string, result int, err error) {
	val := make([]byte, 128)
	result, err = fetcp.FETCP_GetSocketPara(r.socketHnd, &([]byte(param))[0], &val[0])
	n := bytes.Index(val, []byte{0})
	value = string(val[:n])
	return
}

func (r *Reader) SetSocketParam(param string, value string) (result int, err error) {
	result, err = fetcp.FETCP_SetSocketPara(r.socketHnd, &([]byte(param))[0], &([]byte(value))[0])
	return
}

// FEISC
func (r *Reader) NewReader() (err error) {
	r.readerHnd, err = feisc.FEISC_NewReader(r.socketHnd)
	return
}

func (r *Reader) DeleteReader() (result int, err error) {
	result, err = feisc.FEISC_DeleteReader(r.readerHnd)
	return
}

func (r *Reader) GetIscDllVersion() {
	ver := make([]byte, 256)
	feisc.FEISC_GetDLLVersion(&ver[0])
	n := bytes.Index(ver, []byte{0})
	r.iscDllVersion = string(ver[:n])
	return
}

func (r *Reader) GetIscLastError() (errorCode int, errorText string, result int, err error) {
	text := make([]byte, 256)
	result, err = feisc.FEISC_GetLastError(r.readerHnd, &errorCode, &text[0])
	n := bytes.Index(text, []byte{0})
	errorText = string(text[:n])
	return
}

func (r *Reader) GetLastStatus() (status int, statusText string, err error) {
	text := make([]byte, 128)
	status, err = feisc.FEISC_GetLastStatus(r.readerHnd, &text[0])
	n := bytes.Index(text, []byte{0})
	statusText = string(text[:n])
	return
}

func (r *Reader) GetReaderParam(param string) (value string, result int, err error) {
	val := make([]byte, 128)
	result, err = feisc.FEISC_GetReaderPara(r.readerHnd, &([]byte(param))[0], &val[0])
	n := bytes.Index(val, []byte{0})
	value = string(val[:n])
	return
}

func (r *Reader) SetReaderParam(param string, value string) (result int, err error) {
	result, err = feisc.FEISC_SetReaderPara(r.readerHnd, &([]byte(param))[0], &([]byte(value))[0])
	return
}

func (r *Reader) ResetCPU() (status int, err error) {
	status, err = feisc.FEISC_0x63_CPUReset(r.readerHnd, r.busAddr)
	return
}

func (r *Reader) GetSoftVersion(dataFormat int) (version string, status int, err error) {
	var ver []byte
	if dataFormat == 0 {
		ver = make([]byte, 8)
	} else {
		ver = make([]byte, 15)
	}
	status, err = feisc.FEISC_0x65_SoftVersion(r.readerHnd, r.busAddr, &ver[0], dataFormat)
	n := bytes.Index(ver, []byte{0})
	version = string(ver[:n])
	return
}

func (r *Reader) ResetRF() (status int, err error) {
	status, err = feisc.FEISC_0x69_RFReset(r.readerHnd, r.busAddr)
	return
}

func (r *Reader) SetRFOnOff(rfState byte) (status int, err error) {
	status, err = feisc.FEISC_0x6A_RFOnOff(r.readerHnd, r.busAddr, byte(rfState))
	return
}

func (r *Reader) ReadConfBlock(configAddr byte, dataFormat int) (configBlock string, status int, err error) {
	var block []byte
	if dataFormat == 0 {
		block = make([]byte, 15)
	} else {
		block = make([]byte, 29)
	}
	status, err = feisc.FEISC_0x80_ReadConfBlock(r.readerHnd, r.busAddr, configAddr, &block[0], dataFormat)
	n := bytes.Index(block, []byte{0})
	configBlock = string(block[:n])
	return
}

func (r *Reader) WriteConfBlock(configAddr byte, configBlock string, dataFormat int) (status int, err error) {
	status, err = feisc.FEISC_0x81_WriteConfBlock(r.readerHnd, r.busAddr, configAddr, &([]byte(configBlock))[0], dataFormat)
	return
}

func (r *Reader) SaveConfBlock(configAddr byte) (status int, err error) {
	status, err = feisc.FEISC_0x82_SaveConfBlock(r.readerHnd, r.busAddr, configAddr)
	return
}

func (r *Reader) ResetConfBlock(configAddr byte) (status int, err error) {
	status, err = feisc.FEISC_0x83_ResetConfBlock(r.readerHnd, r.busAddr, configAddr)
	return
}

func (r *Reader) SendISOCmd(reqData string, reqLen int, dataFormat int) (respData string, respLen int, status int, err error) {
	resp := make([]byte, 1024)
	status, err = feisc.FEISC_0xB0_ISOCmd(r.readerHnd, r.busAddr, &([]byte(reqData))[0], reqLen, &resp[0], &respLen, dataFormat)
	n := bytes.Index(resp, []byte{0})
	respData = string(resp[:n])
	return
}
