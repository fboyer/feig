// +build windows
package api

//sys	FETCP_Connect(cHostAdr *byte, iPortNr int) (result int) = fetcp.FETCP_Connect
//sys	FETCP_Disconnect(iSocketHnd int) (result int) [failretval!=0] = fetcp.FETCP_DisConnect
//sys	FETCP_Detect(cHostAdr *byte, iPortNr int) (result int) = fetcp.FETCP_Detect
//sys	FETCP_GetSocketState(cHostAdr *byte, iPortNr int) (result int) = fetcp.FETCP_GetSocketState
//sys	FETCP_GetSocketList(iNext int) (result int) = fetcp.FETCP_GetSocketList
//sys	FETCP_GetDLLVersion(cVersion *byte) = fetcp.FETCP_GetDLLVersion
//sys	FETCP_GetErrorText(iErrorCode int, cErrorText *byte) (result int) = fetcp.FETCP_GetErrorText
//sys	FETCP_GetLastError(iSocketHnd int, iErrorCode *int, cErrorText *byte) (result int) = fetcp.FETCP_GetLastError
//sys	FETCP_GetSocketHnd(cHostAdr *byte, iPortNr int) (result int) = fetcp.FETCP_GetSocketHnd
//sys	FETCP_GetSocketPara(iSocketHnd int, cPara *byte, cValue *byte) (result int) = fetcp.FETCP_GetSocketPara
//sys	FETCP_SetSocketPara(iSocketHnd int, cPara *byte, cValue *byte) (result int) = fetcp.FETCP_SetSocketPara
//sys	FETCP_Transceive(iSocketHnd int, cSendProt *byte, iSendLen int, recvBuf *byte, recvBufLen int) (result int) = fetcp.FETCP_Transceive
//sys	FETCP_Transmit(iSocketHnd int, cSendProt *byte, iSendLen int) (result int) = fetcp.FETCP_Transmit
//sys	FETCP_Receive(iSocketHnd int, cRecProt *byte, cRecLec int) (result int) = fetcp.FETCP_Receive
