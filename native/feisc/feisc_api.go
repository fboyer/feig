package feisc

// Administration functions for Reader Objects
//sys	FEISC_NewReader(iPortHnd int) (result int, err error) [failretval<0] = feisc.FEISC_NewReader
//sys	FEISC_DeleteReader(iReaderHnd int) (result int, err error) [failretval<0] = feisc.FEISC_DeleteReader
//sys	FEISC_GetReaderList(iNext int) (result int, err error) [failretval<0] = feisc.FEISC_GetReaderList
//sys	FEISC_GetReaderPara(iReaderHnd int, cPara *byte, cValue *byte) (result int, err error) [failretval<0] = feisc.FEISC_GetReaderPara
//sys	FEISC_SetReaderPara(iReaderHnd int, cPara *byte, cValue *byte) (result int, err error) [failretval<0] = feisc.FEISC_SetReaderPara
//sys	FEISC_GetDLLVersion(cVersion *byte) = feisc.FEISC_GetDLLVersion
//sys	FEISC_GetErrorText(iErrorCode int, cErrorText *byte) (result int, err error) [failretval<0] = feisc.FEISC_GetErrorText
//sys	FEISC_GetStatusText(ucStatus byte, cStatusText *byte) (result int, err error) [failretval<0] = feisc.FEISC_GetStatusText
// To-Do: Add missing API calls
// FEISC_AddEventHandler(iReaderHnd int, pInit *FEISC_EVENT_INIT) (result int, err error) [failretval<0] = feisc.FEISC_AddEventHandler
// FEISC_DelEventHandler(iReaderHnd int, pInit *FEISC_EVENT_INIT) (result int, err error) [failretval<0] = feisc.FEISC_DelEventHandler

// Functions for Plug-in objects to connect alternative port types
// To-Do: Add missing API calls

// Protocol functions
// To-Do: Add missing API calls

// Query function
//sys	 FEISC_GetLastStatus(iReaderHnd int, cStatusText *byte) (result int, err error) [failretval<0] = feisc.FEISC_GetLastStatus
//sys	 FEISC_GetLastError(iReaderHnd int, iErrorCode *int, cErrorText *byte) (result int, err error) [failretval<0] = feisc.FEISC_GetLastError
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
