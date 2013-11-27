// +build windows

package feisc

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

// Administration functions for Reader Objects
//sys	FEISC_NewReader(iPortHnd int) (result int, err error) [failretval<0] = feisc.FEISC_NewReader
//sys	FEISC_DeleteReader(iReaderHnd int) (result int, err error) [failretval<0] = feisc.FEISC_DeleteReader
//sys	FEISC_GetReaderList(iNext int) (result int, err error) [failretval<0] = feisc.FEISC_GetReaderList
//sys	FEISC_GetReaderPara(iReaderHnd int, cPara *byte, cValue *byte) (result int, err error) [failretval<0] = feisc.FEISC_GetReaderPara
//sys	FEISC_SetReaderPara(iReaderHnd int, cPara *byte, cValue *byte) (result int, err error) [failretval<0] = feisc.FEISC_SetReaderPara
//sys	FEISC_GetDLLVersion(cVersion *byte) (err error) = feisc.FEISC_GetDLLVersion
//sys	FEISC_GetErrorText(iErrorCode int, cErrorText *byte) (result int, err error) [failretval<0] = feisc.FEISC_GetErrorText
//sys	FEISC_GetStatusText(ucStatus byte, cStatusText *byte) (result int, err error) [failretval<0] = feisc.FEISC_GetStatusText
// To-Do: Add missing API calls
// FEISC_AddEventHandler(iReaderHnd int, pInit *FEISC_EVENT_INIT) (result int, err erro) [failretval<0] = feisc.FEISC_AddEventHandler
// FEISC_DelEventHandler(iReaderHnd int, pInit *FEISC_EVENT_INIT) (result int, err erro) [failretval<0] = feisc.FEISC_DelEventHandler

// Functions for Plug-in objects to connect alternative port types
// To-Do: Add missing API calls

// Protocol functions
// To-Do: Add missing API calls

// Query function
// To-Do: Add missing API calls

// General communication functions
// To-Do: Add missing API calls

// Special communication functions
//sys	FEISC_0x63_CPUReset(iReaderHnd int, cBusAdr byte) (result int, err error) [failretval<0] = feisc.FEISC_0x63_CPUReset
//sys	FEISC_0x65_SoftVersion(iReaderHnd int, cBusAdr byte, cVersion *byte, iDataFormat int) (result int, err error) [failretval<0] = feisc.FEISC_0x65_SoftVersion
//sys	FEISC_0x69_RFReset(iReaderHnd int, cBusAdr byte) (result int, err error) [failretval<0] = feisc.FEISC_0x69_RFReset
//sys	FEISC_0x6A_RFOnOff(iReaderHnd int, cBusAdr byte, cRF byte) (result int, err error) [failretval<0] = feisc.FEISC_0x6A_RFOnOff
//sys	FEISC_0x80_ReadConfBlock(iReaderHnd int, cBusAdr byte, cConfAdr byte, cConfBlock *byte, iDataFormat int) (result int, err error) [failretval<0] = feisc.FEISC_0x80_ReadConfBlock
//sys	FEISC_0x81_WriteConfBlock(iReaderHnd int, cBusAdr byte, cConfAdr byte, cConfBlock *byte, iDataFormat int) (result int, err error) [failretval<0] = feisc.FEISC_0x81_WriteConfBlock
//sys	FEISC_0x82_SaveConfBlock(iReaderHnd int, cBusAdr byte, cConfAdr byte) (result int, err error) [failretval<0] = feisc.FEISC_0x82_SaveConfBlock
//sys	FEISC_0x83_ResetConfBlock(iReaderHnd int, cBusAdr byte, cConfAdr byte) (result int, err error) [failretval<0] = feisc.FEISC_0x83_ResetConfBlock
//sys	FEISC_0xB0_ISOCmd(iReaderHnd int, cBusAdr byte, cReqData *byte, iReqLen int, cRspData *byte, iRspLen *int, iDataFormat int) (result int, err error) [failretval<0] = feisc.FEISC_0xB0_ISOCmd

// Special functions for asynchronous tasks
// To-Do: Add missing API calls
// FEISC_StartAsyncTask(iReaderHnd int, iTaskID int, pInit *FEISC_TASK_INIT, pInput *void) (result int, err error) [failretval<0] = feisc.FEISC_StartAsyncTask
// FEISC_CancelAsyncTask(iReaderHnd int) (result int, err error) [failretval<0] = feisc.FEISC_CancelAsyncTask
// FEISC_TriggerAsyncTask(iReaderHnd int) (result int, err error) [failretval<0] = feisc.FEISC_TriggerAsyncTask
