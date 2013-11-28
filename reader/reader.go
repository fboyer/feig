package reader

import (
	"bytes"
	"github.com/fboyer/gofeig/feisc"
	"github.com/fboyer/gofeig/fetcp"
)

type Reader struct {
	hostAddr                   string
	port, socketHnd, readerHnd int
}

// FETCP
func (r *Reader) TCPConnect() (err error) {
	r.socketHnd, err = FETCP_Connect(&([]byte(hostAddr))[0], port)
	return
}

func (r *Reader) TCPDisconnect(socketHnd int) (status int, err error) {
	result, err = FETCP_DisConnect(socketHnd)
	return
}

func (r *Reader) TCPDetect(hostAddr string, port int) (result int, err error) {
	// To-Do
	return 0, nil
}

func (r *Reader) TCPGetSocketState(hostAddr string, port int) (status int, err error) {
	// To-Do
	return 0, nil
}

func (r *Reader) TCPGetSocketList(nextSocketHnd int) (socketHnd int, err error) {
	// To-Do
	return 0, nil
}

func (r *Reader) TCPGetDLLVersion() (version string, err error) {
	ver := make([]byte, 256)
	err = FETCP_GetDLLVersion(&ver[0])
	n := bytes.Index(ver, []byte{0})
	version = string(ver[:n])
	return
}

func (r *Reader) TCPGetErrorText(errorCode int) (errorText string, err error) {
	text := make([]byte, 256)
	_, err = FETCP_GetErrorText(errorCode, &text[0])
	n := bytes.Index(text, []byte{0})
	errorText = string(text[:n])
	return
}

func (r *Reader) TCPGetLastError(socketHnd int) (errorCode int, errorText string, result int, err error) {
	text := make([]byte, 256)
	result, err = FETCP_GetLastError(socketHnd, &errorCode, &text[0])
	n := bytes.Index(text, []byte{0})
	errorText = string(text[:n])
	return
}

func (r *Reader) TCPGetSocketHnd(hostAddr string, port int) (socketHnd int, err error) {
	socketHnd, err = FETCP_GetSocketHnd(&([]byte(hostAddr))[0], port)
	return
}

func (r *Reader) TCPGetSocketParam(socketHnd int, param string) (value string, result int, err error) {
	val := make([]byte, 128)
	result, err = FETCP_GetSocketPara(socketHnd, &([]byte(param))[0], &val[0])
	n := bytes.Index(val, []byte{0})
	value = string(val[:n])
	return
}

func (r *Reader) TCPSetSocketParam(socketHnd int, param string, value string) (result int, err error) {
	result, err = FETCP_SetSocketPara(socketHnd, &([]byte(param))[0], &([]byte(value))[0])
	return
}

func (r *Reader) TCPTransceive(socketHnd int, sendBuf string, sendBufLen int, recvBufLen int) (recvBuf string, err error) {
	// To-Do
	return "", nil
}

func (r *Reader) TCPTransmit(socketHnd int, sendBuf string, sendBufLen int) (err error) {
	// To-Do
	return nil
}

func (r *Reader) TCPReceive(socketHnd int, recvBufLen int) (recvBuf string, err error) {
	// To-Do
	return "", nil
}

// FEISC
func (r *Reader) ISCNewReader(portHnd int) (readerHnd int, err error) {
	readerHnd, err = FEISC_NewReader(portHnd)
	return
}

func (r *Reader) ISCDeleteReader(readerHnd int) (result int, err error) {
	result, err = FEISC_DeleteReader(readerHnd)
	return
}

func (r *Reader) ISCGetReaderList(nextReaderHnd int) (readerHnd int, err error) {
	readerHnd, err = FEISC_GetReaderList(nextReaderHnd)
	return
}

func (r *Reader) ISCGetDLLVersion() (version string, err error) {
	ver := make([]byte, 256)
	err := FEISC_GetDLLVersion(&ver[0])
	n := bytes.Index(ver, []byte{0})
	version = string(ver[:n])
	return
}

func (r *Reader) ISCGetErrorText(errorCode int) (errorText string, result int, err error) {
	text := make([]byte, 256)
	result, err = FEISC_GetErrorText(errorCode, &text[0])
	n := bytes.Index(text, []byte{0})
	errorText = string(text[:n])
	return
}

func (r *Reader) ISCGetStatusText(status int) (statusText string, result int, err error) {
	text := make([]byte, 128)
	result, err := FEISC_GetStatusText(byte(status), &text[0])
	n := bytes.Index(text, []byte{0})
	statusText = string(text[:n])
	return
}

func (r *Reader) ISCGetReaderParam(readerHnd int, param string) (value string, result int, err error) {
	val := make([]byte, 128)
	result, err = FEISC_GetReaderPara(readerHnd, &([]byte(param))[0], &val[0])
	n := bytes.Index(text, []byte{0})
	value = string(val[:n])
	return
}

func (r *Reader) ISCSetReaderParam(readerHnd int, param string, value string) (result int, err error) {
	result, err = FEISC_SetReaderPara(readerHnd, &([]byte(param))[0], &([]byte(value))[0])
	return
}

func (r *Reader) ISCResetCPU(readerHnd int, busAddr byte) (status int, err error) {
	status, err = FEISC_0x63_CPUReset(readerHnd, busAddr)
	return
}

func (r *Reader) ISCGetSoftVersion(readerHnd int, busAddr byte, dataFormat int) (version string, status int, err error) {
	if dataFormat == 0 {
		ver := make([]byte, 8)
	} else {
		ver := make([]byte, 15)
	}
	status, err = FEISC_0x65_SoftVersion(readerHnd, busAddr, &ver[0], dataFormat)
	n := bytes.Index(ver, []byte{0})
	version = string(ver[:n])
	return
}

func (r *Reader) ISCResetRF(readerHnd int, busAddr byte) (status int, err error) {
	status, err = FEISC_0x69_RFReset(readerHnd, busAddr)
	return
}

func (r *Reader) ISCSetRFOnOff(readerHnd int, busAddr byte, rfState bool) (status int, err error) {
	status, err = FEISC_0x6A_RFOnOff(readerHnd, busAddr, byte(rfState))
	return
}

func (r *Reader) ISCReadConfBlock(readerHnd int, busAddr byte, configAddr byte, dataFormat int) (configBlock string, status int, err error) {
	if dataFormat == 0 {
		block := make([]byte, 15)
	} else {
		block := make([]byte, 29)
	}
	status, err = FEISC_0x80_ReadConfBlock(readerHnd, busAddr, configAddr, &block[0], dataFormat)
	n := bytes.Index(text, []byte{0})
	configBlock = string(block[:n])
	return
}

func (r *Reader) ISCWriteConfBlock(readerHnd int, busAddr byte, configAddr byte, configBlock string, dataFormat int) (status int, err error) {
	status, err = FEISC_0x81_WriteConfBlock(readerHnd, busAddr, byte(config), &([]byte(configBlock))[0], dataFormat)
	return
}

func (r *Reader) ISCSaveConfBlock(readerHnd int, busAddr byte, configAddr byte) (status int, err error) {
	status, err = FEISC_0x82_SaveConfBlock(readerHnd, busAddr, configAddr)
	return
}

func (r *Reader) ISCResetConfBlock(readerHnd int, busAddr byte, configAddr byte) (status int, err error) {
	status, err = FEISC_0x83_ResetConfBlock(readerHnd, busAddr, configAddr)
	return
}

func (r *Reader) ISCSendISOCmd(readerHnd int, busAddr byte, reqData string, reqLen int, dataFormat int) (respData string, respLen int, status int, err error) {
	resp := make([]byte, 1024)
	status, err = FEISC_0xB0_ISOCmd(readerHnd, busAddr, &([]byte(reqData))[0], reqLen, &resp[0], &respLen, dataFormat)
	n := bytes.Index(resp, []byte{0})
	respData = string(resp[:n])
	return
}
