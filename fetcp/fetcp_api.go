// +build windows

package fetcp

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

//sys	FETCP_Connect(cHostAdr *byte, iPortNr int) (result int, err error) [failretval<0] = fetcp.FETCP_Connect
//sys	FETCP_Disconnect(iSocketHnd int) (result int, err error) [failretval!=0] = fetcp.FETCP_DisConnect
//sys	FETCP_Detect(cHostAdr *byte, iPortNr int) (result int, err error) [failretval<0] = fetcp.FETCP_Detect
//sys	FETCP_GetSocketState(cHostAdr *byte, iPortNr int) (result int, err error) [failretval<0] = fetcp.FETCP_GetSocketState
//sys	FETCP_GetSocketList(iNext int) (result int, err error) [failretval<0] = fetcp.FETCP_GetSocketList
//sys	FETCP_GetDLLVersion(cVersion *byte) = fetcp.FETCP_GetDLLVersion
//sys	FETCP_GetErrorText(iErrorCode int, cErrorText *byte) (result int, err error) [failretval<0] = fetcp.FETCP_GetErrorText
//sys	FETCP_GetLastError(iSocketHnd int, iErrorCode *int, cErrorText *byte) (result int, err error) [failretval<0] = fetcp.FETCP_GetLastError
//sys	FETCP_GetSocketHnd(cHostAdr *byte, iPortNr int) (result int, err error) [failretval<0] = fetcp.FETCP_GetSocketHnd
//sys	FETCP_GetSocketPara(iSocketHnd int, cPara *byte, cValue *byte) (result int, err error) [failretval<0] = fetcp.FETCP_GetSocketPara
//sys	FETCP_SetSocketPara(iSocketHnd int, cPara *byte, cValue *byte) (result int, err error) [failretval<0] = fetcp.FETCP_SetSocketPara
//sys	FETCP_Transceive(iSocketHnd int, cSendProt *byte, iSendLen int, recvBuf *byte, recvBufLen int) (result int, err error) [failretval<0] = fetcp.FETCP_Transceive
//sys	FETCP_Transmit(iSocketHnd int, cSendProt *byte, iSendLen int) (result int, err error) [failretval<0] = fetcp.FETCP_Transmit
//sys	FETCP_Receive(iSocketHnd int, cRecProt *byte, cRecLec int) (result int, err error) [failretval<0] = fetcp.FETCP_Receive
