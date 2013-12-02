package gofetcp

import (
	"bytes"
	"github.com/fboyer/gofeig/native/fetcp"
)

const (
	// Error Codes
	FETCP_ERR_NEWSOCKET_FAILURE      = -1200
	FETCP_ERR_EMPTY_LIST             = -1201
	FETCP_ERR_POINTER_IS_NULL        = -1202
	FETCP_ERR_NO_MEMORY              = -1203
	FETCP_ERR_SERVER_NOT_FOUND       = -1205
	FETCP_ERR_NO_CONNECTION          = -1211
	FETCP_ERR_SERVER_ADDR            = -1212
	FETCP_ERR_UNKNOWN_HND            = -1220
	FETCP_ERR_HND_IS_NULL            = -1221
	FETCP_ERR_HND_IS_NEGATIVE        = -1222
	FETCP_ERR_NO_HND_FOUND           = -1223
	FETCP_ERR_TIMEOUT                = -1230
	FETCP_ERR_RECEIVE_PROCESS        = -1232
	FETCP_ERR_CHANGE_PORT_PARA       = -1236
	FETCP_ERR_TRANSMIT_PROCESS       = -1237
	FETCP_ERR_GET_CONNECTION_STATE   = -1238
	FETCP_ERR_UNKNOWN_PARAMETER      = -1250
	FETCP_ERR_PARAMETER_OUT_OF_RANGE = -1251
	FETCP_ERR_ODD_PARAMETERSTRING    = -1252
	FETCP_ERR_UNKNOWN_ERRORCODE      = -1254
	FETCP_ERR_BUFFER_OVERFLOW        = -1270

	// TCP States
	FETCP_STATE_CLOSED      = 1
	FETCP_STATE_LISTEN      = 2
	FETCP_STATE_SYN_SENT    = 3
	FETCP_STATE_SYN_RCVD    = 4
	FETCP_STATE_ESTABLISHED = 5
	FETCP_STATE_FIN_WAIT1   = 6
	FETCP_STATE_FIN_WAIT2   = 7
	FETCP_STATE_CLOSE_WAIT  = 8
	FETCP_STATE_CLOSING     = 9
	FETCP_STATE_LAST_ACK    = 10
	FETCP_STATE_TIME_WAIT   = 11
	FETCP_STATE_DELETE_TCB  = 12
)

func Connect(hostAddr string, port int) (socketHnd int, err error) {
	socketHnd, err = fetcp.FETCP_Connect(&([]byte(hostAddr))[0], port)
	return
}

func Disconnect(socketHnd int) (result int, err error) {
	result, err = fetcp.FETCP_Disconnect(socketHnd)
	return
}

func GetDllVersion() (version string) {
	ver := make([]byte, 256)
	fetcp.FETCP_GetDLLVersion(&ver[0])
	n := bytes.Index(ver, []byte{0})
	version = string(ver[:n])
	return
}

func GetLastError(socketHnd int) (errorCode int, errorText string, result int, err error) {
	text := make([]byte, 256)
	result, err = fetcp.FETCP_GetLastError(socketHnd, &errorCode, &text[0])
	n := bytes.Index(text, []byte{0})
	errorText = string(text[:n])
	return
}

func GetSocketParam(socketHnd int, param string) (value string, result int, err error) {
	val := make([]byte, 128)
	result, err = fetcp.FETCP_GetSocketPara(socketHnd, &([]byte(param))[0], &val[0])
	n := bytes.Index(val, []byte{0})
	value = string(val[:n])
	return
}

func SetSocketParam(socketHnd int, param string, value string) (result int, err error) {
	result, err = fetcp.FETCP_SetSocketPara(socketHnd, &([]byte(param))[0], &([]byte(value))[0])
	return
}
