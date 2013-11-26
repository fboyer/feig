package feigreader

import (
	"feig/feisc"
	"feig/fetcp"
)

// FETCP
func Connect(hostAddr string, port int) (socketHnd int, err error) {
	socketHnd, err = FETCP_Connect(&([]byte(hostAddr))[0], uint32(port))
	return
}

func Disconnect(socketHnd int) (status int, err error) {
	result, err = FETCP_DisConnect(uint32(socketHnd))
	return
}

func Detect(hostAddr string, port int) (result int, err error) {
	// To-Do
	return 0, nil
}

func GetSocketState(hostAddr string, port int) (status int, err error) {
	// To-Do
	return 0, nil
}

func GetSocketList(nextSocketHnd int) (socketHnd int, err error) {
	// To-Do
	return 0, nil
}

// Name collision with FEISC
func GetDLLVersion() (version string, err error) {
	ver := make([]byte, 256)
	err := FETCP_GetDLLVersion(&ver[0])
	version = string(ver[:])
	return
}

// Name collision with FEISC
func GetErrorText(errorCode int) (errorText string, err error) {
	text := make([]byte, 256)
	_, err = FETCP_GetErrorText(uint32(errorCode), &text[0])
	errorText = string(text[:])
	return
}

func GetLastError(socketHnd int) (errorCode int, errorText string, result int, err error) {
	text := make([]byte, 256)
	result, err = FETCP_GetLastError(uint32(socketHnd), &errorCode, &text[0])
	errorText = string(text[:])
	return
}

func GetSocketHnd(hostAddr string, port int) (socketHnd int, err error) {
	socketHnd, err = FETCP_GetSocketHnd(&([]byte(hostAddr))[0], uint32(port))
	return
}

func GetSocketParam(socketHnd int, param string) (value string, result int, err error) {
	val := make([]byte, 128)
	result, err = FETCP_GetSocketPara(uint32(socketHnd), &([]byte(param))[0], &val[0])
	value = string(val[:])
	return
}

func SetSocketParam(socketHnd int, param string, value string) (result int, err error) {
	result, err = FETCP_SetSocketPara(uint32(socketHnd), &([]byte(param))[0], &([]byte(value))[0])
	return
}

func Transceive(socketHnd int, sendBuf string, sendBufLen int, recvBufLen int) (recvBuf string, err error) {
	// To-Do
	return "", nil
}

func Transmit(socketHnd int, sendBuf string, sendBufLen int) (err error) {
	// To-Do
	return nil
}

func Receive(socketHnd int, recvBufLen int) (recvBuf string, err error) {
	// To-Do
	return "", nil
}

// FEISC
func NewReader(portHnd int) (readerHnd int, err error) {
	readerHnd, err = FEISC_NewReader(uint32(portHnd))
	return
}

func DeleteReader(readerHnd int) (result int, err error) {
	result, err = FEISC_DeleteReader(uint32(readerHnd))
	return
}

func GetReaderList(nextReaderHnd int) (readerHnd int, err error) {
	readerHnd, err = FEISC_GetReaderList(uint32(nextReaderHnd))
	return
}

// Name collision with FETCP
func GetDLLVersion() (version string, err error) {
	ver := make([]byte, 256)
	err := FEISC_GetDLLVersion(&ver[0])
	version = string(ver[:])
	return
}

// Name collision with FETCP
func GetErrorText(errorCode int) (errorText string, result int, err error) {
	text := make([]byte, 256)
	result, err = FEISC_GetErrorText(uint32(errorCode), &text[0])
	errorText = string(text[:])
	return
}

func GetStatusText(status int) (statusText string, result int, err error) {
	text := make([]byte, 128)
	result, err := FEISC_GetStatusText(byte(status), &text[0])
	statusText = string(text[:])
	return
}

func GetReaderParam(readerHnd int, param string) (value string, result int, err error) {
	val := make([]byte, 128)
	result, err = FEISC_GetReaderPara(uint32(readerHnd), &([]byte(param))[0], &val[0])
	value = string(val[:])
	return
}

func SetReaderParam(readerHnd int, param string, value string) (result int, err error) {
	result, err = FEISC_SetReaderPara(uint32(readerHnd), &([]byte(param))[0], &([]byte(value))[0])
	return
}

func ResetCPU(readerHnd int, busAddr byte) (status int, err error) {
	status, err = FEISC_0x63_CPUReset(uint32(readerHnd), busAddr)
	return
}

func GetSoftVersion(readerHnd int, busAddr byte, dataFormat int) (version string, status int, err error) {
	if dataFormat == 0 {
		ver := make([]byte, 8)
	} else {
		ver := make([]byte, 15)
	}
	status, err = FEISC_0x65_SoftVersion(uint32(readerHnd), busAddr, &ver[0], uint32(dataFormat()))
	version = string(ver[:])
	return
}

func ResetRF(readerHnd int, busAddr byte) (status int, err error) {
	status, err = FEISC_0x69_RFReset(uint32(readerHnd), busAddr)
	return
}

func SetRFOnOff(readerHnd int, busAddr byte, rfState bool) (status int, err error) {
	status, err = FEISC_0x6A_RFOnOff(uint32(readerHnd), busAddr, byte(rfState))
	return
}

func ReadConfBlock(readerHnd int, busAddr byte, configAddr byte, dataFormat int) (configBlock string, status int, err error) {
	if dataFormat == 0 {
		block := make([]byte, 15)
	} else {
		block := make([]byte, 29)
	}
	status, err = FEISC_0x80_ReadConfBlock(uint32(readerHnd), busAddr, configAddr, &block[0], uint32(dataFormat))
	configBlock = string(block[:])
	return
}

func WriteConfBlock(readerHnd int, busAddr byte, configAddr byte, configBlock string, dataFormat int) (status int, err error) {
	status, err = FEISC_0x81_WriteConfBlock(uint32(readerHnd), busAddr, byte(config), &([]byte(configBlock))[0], byte(dataFormat))
	return
}

func SaveConfBlock(readerHnd int, busAddr byte, configAddr byte) (status int, err error) {
	status, err = FEISC_0x82_SaveConfBlock(uint32(readerHnd), busAddr, configAddr)
	return
}

func ResetConfBlock(readerHnd int, busAddr byte, configAddr byte) (status int, err error) {
	status, err = FEISC_0x83_ResetConfBlock(uint32(readerHnd), busAddr, configAddr)
	return
}

func SendISOCmd(readerHnd int, busAddr byte, reqData string, reqLen int, iDataFormat uint32) (respData string, respLen int, status int, err error) {
	resp := make([]byte, 1024)
	status, err = FEISC_0xB0_ISOCmd(uint32(readerHnd), busAddr, &([]byte(reqData))[0], uint32(reqLen), &resp[0], &respLen)
	respData = string(resp[:])
	return
}
