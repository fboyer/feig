package gofeig

import (
	"bytes"
	"github.com/fboyer/gofeig/native/feisc"
)

const (
	// Error codes
	FEISC_ERR_NEWREADER_FAILURE      = -4000
	FEISC_ERR_EMPTY_LIST             = -4001
	FEISC_ERR_POINTER_IS_NULL        = -4002
	FEISC_ERR_NO_MORE_MEM            = -4003
	FEISC_ERR_UNKNOWN_COMM_PORT      = -4004
	FEISC_ERR_UNSUPPORTED_FUNCTION   = -4005
	FEISC_ERR_NO_USB_SUPPORT         = -4006
	FEISC_ERR_OLD_FECOM              = -4007
	FEISC_ERR_NO_VALUE               = -4010
	FEISC_ERR_UNKNOWN_HND            = -4020
	FEISC_ERR_HND_IS_NULL            = -4021
	FEISC_ERR_HND_IS_NEGATIVE        = -4022
	FEISC_ERR_NO_HND_FOUND           = -4023
	FEISC_ERR_PORTHND_IS_NEGATIVE    = -4024
	FEISC_ERR_HND_UNVALID            = -4025
	FEISC_ERR_PROTLEN                = -4030
	FEISC_ERR_CHECKSUM               = -4031
	FEISC_ERR_BUSY_TIMEOUT           = -4032
	FEISC_ERR_UNKNOWN_STATUS         = -4033
	FEISC_ERR_NO_RECPROTOCOL         = -4034
	FEISC_ERR_CMD_BYTE               = -4035
	FEISC_ERR_TRANSCEIVE             = -4036
	FEISC_ERR_REC_BUS_ADR            = -4037
	FEISC_ERR_UNKNOWN_PARAMETER      = -4050
	FEISC_ERR_PARAMETER_OUT_OF_RANGE = -4051
	FEISC_ERR_ODD_PARAMETERSTRING    = -4052
	FEISC_ERR_UNKNOWN_ERRORCODE      = -4053
	FEISC_ERR_UNSUPPORTED_OPTION     = -4054
	FEISC_ERR_UNKNOWN_EPC_TYPE       = -4055
	FEISC_ERR_NO_PLUGIN              = -4060
	FEISC_ERR_PLUGIN_PRESENT         = -4061
	FEISC_ERR_UNKNOWN_PLUGIN_ID      = -4062
	FEISC_ERR_PI_BUILD_DATA          = -4063
	FEISC_ERR_PI_BUILD_FRAME         = -4064
	FEISC_ERR_PI_SPLIT_FRAME         = -4065
	FEISC_ERR_PI_SPLIT_DATA          = -4066
	FEISC_ERR_BUFFER_OVERFLOW        = -4070
	FEISC_ERR_TASK_STILL_RUNNING     = -4080
	FEISC_ERR_TASK_NOT_STARTED       = -4081
	FEISC_ERR_TASK_TIMEOUT           = -4082
	FEISC_ERR_TASK_SOCKET_INIT       = -4083
	FEISC_ERR_TASK_BUSY              = -4084
	FEISC_ERR_THREAD_CANCEL_ERROR    = -4085
	FEISC_ERR_CRYPT_LOAD_LIBRARY     = -4090
	FEISC_ERR_CRYPT_INIT             = -4091
	FEISC_ERR_CRYPT_AUTHENT_PROCESS  = -4092
	FEISC_ERR_CRYPT_ENCYPHER         = -4093
	FEISC_ERR_CRYPT_DECYPHER         = -4094

	// List of constants for the FEISC_EVENT_INIT structure
	FEISC_THREAD_ID     = 1
	FEISC_WND_HWND      = 2
	FEISC_CALLBACK      = 3
	FEISC_EVENT         = 4
	FEISC_CALLBACK_2    = 5
	FEISC_CALLBACK_4    = 6
	FEISC_PRT_EVENT     = 1
	FEISC_SNDPRT_EVENT  = 2
	FEISC_RECPRT_EVENT  = 3
	FEISC_SCANNER_EVENT = 4

	// List of constants for TaskID and for the FEISC_TASK_INIT structure
	FEISC_TASKID_FIRST_NEW_TAG      = 1
	FEISC_TASKID_EVERY_NEW_TAG      = 2
	FEISC_TASKID_NOTIFICATION       = 3
	FEISC_TASKID_SAM_COMMAND        = 4
	FEISC_TASKID_COMMAND_QUEUE      = 5
	FEISC_TASKID_MAX_EVENT          = 6
	FEISC_TASKID_PEOPLE_COUNTER     = 7
	FEISC_TASKCB_1                  = 1
	FEISC_TASKCB_2                  = 2
	FEISC_TASKCB_3                  = 3
	FEISC_TASK_CHANNEL_TYPE_AS_OPEN = 1
	FEISC_TASK_CHANNEL_TYPE_NEW_TCP = 5
)

type FEISC_EVENT_INIT struct {
}

type FEISC_TASK_INIT struct {
}

// FEISC
func NewReader(socketHnd int) (readerHnd int) {
	readerHnd = feisc.FEISC_NewReader(socketHnd)
	return
}

func DeleteReader(readerHnd int) (result int) {
	result = feisc.FEISC_DeleteReader(readerHnd)
	return
}

func GetIscDllVersion() (version string) {
	ver := make([]byte, 256)
	feisc.FEISC_GetDLLVersion(&ver[0])
	n := bytes.Index(ver, []byte{0})
	version = string(ver[:n])
	return
}

func GetIscLastError(readerHnd int) (errorCode int, errorText string, result int) {
	text := make([]byte, 256)
	result = feisc.FEISC_GetLastError(readerHnd, &errorCode, &text[0])
	n := bytes.Index(text, []byte{0})
	errorText = string(text[:n])
	return
}

func GetLastState(readerHnd int) (result int, resultText string) {
	text := make([]byte, 128)
	result = feisc.FEISC_GetLastState(readerHnd, &text[0])
	n := bytes.Index(text, []byte{0})
	resultText = string(text[:n])
	return
}

func GetReaderParam(readerHnd int, param string) (value string, result int) {
	val := make([]byte, 128)
	result = feisc.FEISC_GetReaderPara(readerHnd, &([]byte(param))[0], &val[0])
	n := bytes.Index(val, []byte{0})
	value = string(val[:n])
	return
}

func SetReaderParam(readerHnd int, param string, value string) (result int) {
	result = feisc.FEISC_SetReaderPara(readerHnd, &([]byte(param))[0], &([]byte(value))[0])
	return
}

func ResetCpu(readerHnd int, busAddr byte) (result int) {
	result = feisc.FEISC_0x63_CPUReset(readerHnd, busAddr)
	return
}

func GetSoftVersion(readerHnd int, busAddr byte, dataFormat int) (version string, result int) {
	var ver []byte
	if dataFormat == 0 {
		ver = make([]byte, 8)
	} else {
		ver = make([]byte, 15)
	}
	result = feisc.FEISC_0x65_SoftVersion(readerHnd, busAddr, &ver[0], dataFormat)
	n := bytes.Index(ver, []byte{0})
	version = string(ver[:n])
	return
}

func ResetRf(readerHnd int, busAddr byte) (result int) {
	result = feisc.FEISC_0x69_RFReset(readerHnd, busAddr)
	return
}

func SetRfOnOff(readerHnd int, busAddr byte, rfState byte) (result int) {
	result = feisc.FEISC_0x6A_RFOnOff(readerHnd, busAddr, byte(rfState))
	return
}

func ReadConfBlock(readerHnd int, busAddr byte, configAddr byte, dataFormat int) (configBlock string, result int) {
	var block []byte
	if dataFormat == 0 {
		block = make([]byte, 15)
	} else {
		block = make([]byte, 29)
	}
	result = feisc.FEISC_0x80_ReadConfBlock(readerHnd, busAddr, configAddr, &block[0], dataFormat)
	n := bytes.Index(block, []byte{0})
	configBlock = string(block[:n])
	return
}

func WriteConfBlock(readerHnd int, busAddr byte, configAddr byte, configBlock string, dataFormat int) (result int) {
	result = feisc.FEISC_0x81_WriteConfBlock(readerHnd, busAddr, configAddr, &([]byte(configBlock))[0], dataFormat)
	return
}

func SaveConfBlock(readerHnd int, busAddr byte, configAddr byte) (result int) {
	result = feisc.FEISC_0x82_SaveConfBlock(readerHnd, busAddr, configAddr)
	return
}

func ResetConfBlock(readerHnd int, busAddr byte, configAddr byte) (result int) {
	result = feisc.FEISC_0x83_ResetConfBlock(readerHnd, busAddr, configAddr)
	return
}

func SendIsoCmd(readerHnd int, busAddr byte, reqData string, reqLen int, dataFormat int) (respData string, respLen int, result int) {
	resp := make([]byte, 1024)
	result = feisc.FEISC_0xB0_ISOCmd(readerHnd, busAddr, &([]byte(reqData))[0], reqLen, &resp[0], &respLen, dataFormat)
	n := bytes.Index(resp, []byte{0})
	respData = string(resp[:n])
	return
}
