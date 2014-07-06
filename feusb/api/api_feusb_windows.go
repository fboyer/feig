// +build windows
package api

//sys	FEUSB_GetDLLVersion(char* cVersion) = feusb.FEUSB_GetDLLVersion
//sys	FEUSB_GetErrorText(int iError, char* cText) (result int) = feusb.FEUSB_GetErrorText
//sys	FEUSB_GetLastError(int iDevHnd , int* iErrorCode, char* cErrorText) (result int) = feusb.FEUSB_GetLastError
//sys	FEUSB_Scan(int iScanOpt, FEUSB_SCANSEARCH* pSearchOpt) (result int) = feusb.FEUSB_Scan
//sys	FEUSB_ScanAndOpen(int iScanOpt, FEUSB_SCANSEARCH* pSearchOpt) (result int) = feusb.FEUSB_ScanAndOpen
//sys	FEUSB_GetScanListPara(int iIndex, char* cPara, char* cValue) (result int) = feusb.FEUSB_GetScanListPara
//sys	FEUSB_GetScanListSize() (result int) = feusb.FEUSB_GetScanListSize
//sys	FEUSB_ClearScanList() (result int) = feusb.FEUSB_ClearScanList
//sys	FEUSB_AddEventHandler(int iDevHnd, FEUSB_EVENT_INIT* pInit) (result int) = feusb.FEUSB_AddEventHandler
//sys	FEUSB_DelEventHandler(int iDevHnd, FEUSB_EVENT_INIT* pInit) (result int) = feusb.FEUSB_DelEventHandler
//sys	FEUSB_OpenDevice(long nDeviceID) (result int) = feusb.FEUSB_OpenDevice
//sys	FEUSB_CloseDevice(int iDevHnd) (result int) = feusb.FEUSB_CloseDevice
//sys	FEUSB_GetDeviceList(int iNext) (result int) = feusb.FEUSB_GetDeviceList
//sys	FEUSB_GetDeviceHnd(long nDeviceID) (result int) = feusb.FEUSB_GetDeviceHnd
//sys	FEUSB_GetDevicePara(int iDevHnd, char* cPara, char* cValue) (result int) = feusb.FEUSB_GetDevicePara
//sys	FEUSB_SetDevicePara(int iDevHnd, char* cPara, char* cValue) (result int) = feusb.FEUSB_SetDevicePara
//sys	FEUSB_Transceive(int iDevHnd, char* cInterface, int iDir, UCHAR* cSendData, int iSendLen, UCHAR* cRecData, int iRecLen) (result int) = feusb.FEUSB_Transceive
//sys	FEUSB_Transmit(int iDevHnd, char* cInterface, UCHAR* cSendData, int iSendLen) (result int) = feusb.FEUSB_Transmit
//sys	FEUSB_Receive(int iDevHnd, char* cInterface, UCHAR* cRecData, int iRecLen) (result int) = feusb.FEUSB_Receive
