/* Automatically generated nanopb constant definitions */
/* Generated by nanopb-0.3.9-dev at Tue Jan  9 09:52:13 2018. */

#include "fk-app.pb.h"

/* @@protoc_insertion_point(includes) */
#if PB_PROTO_HEADER_VERSION != 30
#error Regenerate this file with the current version of nanopb generator.
#endif



const pb_field_t fk_app_QueryCapabilities_fields[3] = {
    PB_FIELD(  1, UINT32  , SINGULAR, STATIC  , FIRST, fk_app_QueryCapabilities, version, version, 0),
    PB_FIELD(  2, UINT32  , SINGULAR, STATIC  , OTHER, fk_app_QueryCapabilities, callerTime, version, 0),
    PB_LAST_FIELD
};

const pb_field_t fk_app_ModuleCapabilities_fields[3] = {
    PB_FIELD(  1, UINT32  , SINGULAR, STATIC  , FIRST, fk_app_ModuleCapabilities, id, id, 0),
    PB_FIELD(  2, STRING  , SINGULAR, CALLBACK, OTHER, fk_app_ModuleCapabilities, name, id, 0),
    PB_LAST_FIELD
};

const pb_field_t fk_app_SensorCapabilities_fields[6] = {
    PB_FIELD(  1, UINT32  , SINGULAR, STATIC  , FIRST, fk_app_SensorCapabilities, id, id, 0),
    PB_FIELD(  2, UINT32  , SINGULAR, STATIC  , OTHER, fk_app_SensorCapabilities, module, id, 0),
    PB_FIELD(  3, STRING  , SINGULAR, CALLBACK, OTHER, fk_app_SensorCapabilities, name, module, 0),
    PB_FIELD(  4, UINT32  , SINGULAR, STATIC  , OTHER, fk_app_SensorCapabilities, frequency, name, 0),
    PB_FIELD(  5, STRING  , SINGULAR, CALLBACK, OTHER, fk_app_SensorCapabilities, unitOfMeasure, frequency, 0),
    PB_LAST_FIELD
};

const pb_field_t fk_app_Capabilities_fields[6] = {
    PB_FIELD(  1, UINT32  , SINGULAR, STATIC  , FIRST, fk_app_Capabilities, version, version, 0),
    PB_FIELD(  2, BYTES   , SINGULAR, CALLBACK, OTHER, fk_app_Capabilities, deviceId, version, 0),
    PB_FIELD(  3, STRING  , SINGULAR, CALLBACK, OTHER, fk_app_Capabilities, name, deviceId, 0),
    PB_FIELD(  4, MESSAGE , REPEATED, CALLBACK, OTHER, fk_app_Capabilities, modules, name, &fk_app_ModuleCapabilities_fields),
    PB_FIELD(  5, MESSAGE , REPEATED, CALLBACK, OTHER, fk_app_Capabilities, sensors, modules, &fk_app_SensorCapabilities_fields),
    PB_LAST_FIELD
};

const pb_field_t fk_app_TimeSpec_fields[4] = {
    PB_FIELD(  1, INT32   , SINGULAR, STATIC  , FIRST, fk_app_TimeSpec, fixed, fixed, 0),
    PB_FIELD(  2, INT32   , SINGULAR, STATIC  , OTHER, fk_app_TimeSpec, interval, fixed, 0),
    PB_FIELD(  3, INT32   , SINGULAR, STATIC  , OTHER, fk_app_TimeSpec, offset, interval, 0),
    PB_LAST_FIELD
};

const pb_field_t fk_app_Schedule_fields[5] = {
    PB_FIELD(  1, MESSAGE , SINGULAR, STATIC  , FIRST, fk_app_Schedule, second, second, &fk_app_TimeSpec_fields),
    PB_FIELD(  2, MESSAGE , SINGULAR, STATIC  , OTHER, fk_app_Schedule, minute, second, &fk_app_TimeSpec_fields),
    PB_FIELD(  3, MESSAGE , SINGULAR, STATIC  , OTHER, fk_app_Schedule, hour, minute, &fk_app_TimeSpec_fields),
    PB_FIELD(  4, MESSAGE , SINGULAR, STATIC  , OTHER, fk_app_Schedule, day, hour, &fk_app_TimeSpec_fields),
    PB_LAST_FIELD
};

const pb_field_t fk_app_Schedules_fields[5] = {
    PB_FIELD(  1, MESSAGE , SINGULAR, STATIC  , FIRST, fk_app_Schedules, readings, readings, &fk_app_Schedule_fields),
    PB_FIELD(  2, MESSAGE , SINGULAR, STATIC  , OTHER, fk_app_Schedules, transmission, readings, &fk_app_Schedule_fields),
    PB_FIELD(  3, MESSAGE , SINGULAR, STATIC  , OTHER, fk_app_Schedules, status, transmission, &fk_app_Schedule_fields),
    PB_FIELD(  4, MESSAGE , SINGULAR, STATIC  , OTHER, fk_app_Schedules, location, status, &fk_app_Schedule_fields),
    PB_LAST_FIELD
};

const pb_field_t fk_app_NetworkInfo_fields[3] = {
    PB_FIELD(  1, STRING  , SINGULAR, CALLBACK, FIRST, fk_app_NetworkInfo, ssid, ssid, 0),
    PB_FIELD(  2, STRING  , SINGULAR, CALLBACK, OTHER, fk_app_NetworkInfo, password, ssid, 0),
    PB_LAST_FIELD
};

const pb_field_t fk_app_NetworkSettings_fields[3] = {
    PB_FIELD(  1, INT32   , SINGULAR, STATIC  , FIRST, fk_app_NetworkSettings, createAccessPoint, createAccessPoint, 0),
    PB_FIELD(  2, MESSAGE , REPEATED, CALLBACK, OTHER, fk_app_NetworkSettings, networks, createAccessPoint, &fk_app_NetworkInfo_fields),
    PB_LAST_FIELD
};

const pb_field_t fk_app_Identity_fields[4] = {
    PB_FIELD(  1, STRING  , SINGULAR, CALLBACK, FIRST, fk_app_Identity, device, device, 0),
    PB_FIELD(  2, STRING  , SINGULAR, CALLBACK, OTHER, fk_app_Identity, stream, device, 0),
    PB_FIELD(  3, BYTES   , SINGULAR, CALLBACK, OTHER, fk_app_Identity, deviceId, stream, 0),
    PB_LAST_FIELD
};

const pb_field_t fk_app_ConfigureSensorQuery_fields[3] = {
    PB_FIELD(  1, UINT32  , SINGULAR, STATIC  , FIRST, fk_app_ConfigureSensorQuery, id, id, 0),
    PB_FIELD(  2, UINT32  , SINGULAR, STATIC  , OTHER, fk_app_ConfigureSensorQuery, frequency, id, 0),
    PB_LAST_FIELD
};

const pb_field_t fk_app_QueryDataSets_fields[1] = {
    PB_LAST_FIELD
};

const pb_field_t fk_app_DataSet_fields[8] = {
    PB_FIELD(  1, UINT32  , SINGULAR, STATIC  , FIRST, fk_app_DataSet, id, id, 0),
    PB_FIELD(  2, UINT32  , SINGULAR, STATIC  , OTHER, fk_app_DataSet, sensor, id, 0),
    PB_FIELD(  3, UINT64  , SINGULAR, STATIC  , OTHER, fk_app_DataSet, time, sensor, 0),
    PB_FIELD(  4, UINT32  , SINGULAR, STATIC  , OTHER, fk_app_DataSet, size, time, 0),
    PB_FIELD(  5, UINT32  , SINGULAR, STATIC  , OTHER, fk_app_DataSet, pages, size, 0),
    PB_FIELD(  6, UINT32  , SINGULAR, STATIC  , OTHER, fk_app_DataSet, hash, pages, 0),
    PB_FIELD(  7, STRING  , SINGULAR, CALLBACK, OTHER, fk_app_DataSet, name, hash, 0),
    PB_LAST_FIELD
};

const pb_field_t fk_app_DataSets_fields[2] = {
    PB_FIELD(  1, MESSAGE , REPEATED, CALLBACK, FIRST, fk_app_DataSets, dataSets, dataSets, &fk_app_DataSet_fields),
    PB_LAST_FIELD
};

const pb_field_t fk_app_DownloadDataSet_fields[3] = {
    PB_FIELD(  1, UINT32  , SINGULAR, STATIC  , FIRST, fk_app_DownloadDataSet, id, id, 0),
    PB_FIELD(  2, UINT32  , SINGULAR, STATIC  , OTHER, fk_app_DownloadDataSet, page, id, 0),
    PB_LAST_FIELD
};

const pb_field_t fk_app_Sample_fields[3] = {
    PB_FIELD(  1, UINT64  , SINGULAR, STATIC  , FIRST, fk_app_Sample, time, time, 0),
    PB_FIELD(  2, FLOAT   , SINGULAR, STATIC  , OTHER, fk_app_Sample, value, time, 0),
    PB_LAST_FIELD
};

const pb_field_t fk_app_DataSetData_fields[8] = {
    PB_FIELD(  1, UINT64  , SINGULAR, STATIC  , FIRST, fk_app_DataSetData, time, time, 0),
    PB_FIELD(  2, UINT32  , SINGULAR, STATIC  , OTHER, fk_app_DataSetData, page, time, 0),
    PB_FIELD(  3, UINT32  , SINGULAR, STATIC  , OTHER, fk_app_DataSetData, sensor, page, 0),
    PB_FIELD(  4, MESSAGE , REPEATED, CALLBACK, OTHER, fk_app_DataSetData, samples, sensor, &fk_app_Sample_fields),
    PB_FIELD(  5, FLOAT   , REPEATED, CALLBACK, OTHER, fk_app_DataSetData, floats, samples, 0),
    PB_FIELD(  6, BYTES   , SINGULAR, CALLBACK, OTHER, fk_app_DataSetData, data, floats, 0),
    PB_FIELD(  7, UINT32  , SINGULAR, STATIC  , OTHER, fk_app_DataSetData, hash, data, 0),
    PB_LAST_FIELD
};

const pb_field_t fk_app_LiveDataPoll_fields[2] = {
    PB_FIELD(  1, UINT32  , SINGULAR, STATIC  , FIRST, fk_app_LiveDataPoll, interval, interval, 0),
    PB_LAST_FIELD
};

const pb_field_t fk_app_LiveDataSample_fields[4] = {
    PB_FIELD(  1, UINT32  , SINGULAR, STATIC  , FIRST, fk_app_LiveDataSample, sensor, sensor, 0),
    PB_FIELD(  2, UINT64  , SINGULAR, STATIC  , OTHER, fk_app_LiveDataSample, time, sensor, 0),
    PB_FIELD(  3, FLOAT   , SINGULAR, STATIC  , OTHER, fk_app_LiveDataSample, value, time, 0),
    PB_LAST_FIELD
};

const pb_field_t fk_app_LiveData_fields[2] = {
    PB_FIELD(  1, MESSAGE , REPEATED, CALLBACK, FIRST, fk_app_LiveData, samples, samples, &fk_app_LiveDataSample_fields),
    PB_LAST_FIELD
};

const pb_field_t fk_app_QueryDataSet_fields[2] = {
    PB_FIELD(  1, UINT32  , SINGULAR, STATIC  , FIRST, fk_app_QueryDataSet, id, id, 0),
    PB_LAST_FIELD
};

const pb_field_t fk_app_EraseDataSet_fields[2] = {
    PB_FIELD(  1, UINT32  , SINGULAR, STATIC  , FIRST, fk_app_EraseDataSet, id, id, 0),
    PB_LAST_FIELD
};

const pb_field_t fk_app_File_fields[7] = {
    PB_FIELD(  1, UINT32  , SINGULAR, STATIC  , FIRST, fk_app_File, id, id, 0),
    PB_FIELD(  2, UINT64  , SINGULAR, STATIC  , OTHER, fk_app_File, time, id, 0),
    PB_FIELD(  3, UINT32  , SINGULAR, STATIC  , OTHER, fk_app_File, size, time, 0),
    PB_FIELD(  4, UINT32  , SINGULAR, STATIC  , OTHER, fk_app_File, pages, size, 0),
    PB_FIELD(  5, UINT32  , SINGULAR, STATIC  , OTHER, fk_app_File, version, pages, 0),
    PB_FIELD(  6, STRING  , SINGULAR, CALLBACK, OTHER, fk_app_File, name, version, 0),
    PB_LAST_FIELD
};

const pb_field_t fk_app_Files_fields[2] = {
    PB_FIELD(  1, MESSAGE , REPEATED, CALLBACK, FIRST, fk_app_Files, files, files, &fk_app_File_fields),
    PB_LAST_FIELD
};

const pb_field_t fk_app_DownloadFile_fields[5] = {
    PB_FIELD(  1, UINT32  , SINGULAR, STATIC  , FIRST, fk_app_DownloadFile, id, id, 0),
    PB_FIELD(  2, UINT32  , SINGULAR, STATIC  , OTHER, fk_app_DownloadFile, page, id, 0),
    PB_FIELD(  3, UINT32  , SINGULAR, STATIC  , OTHER, fk_app_DownloadFile, pageSize, page, 0),
    PB_FIELD(  4, BYTES   , SINGULAR, CALLBACK, OTHER, fk_app_DownloadFile, token, pageSize, 0),
    PB_LAST_FIELD
};

const pb_field_t fk_app_EraseFile_fields[2] = {
    PB_FIELD(  1, UINT32  , SINGULAR, STATIC  , FIRST, fk_app_EraseFile, id, id, 0),
    PB_LAST_FIELD
};

const pb_field_t fk_app_FileData_fields[6] = {
    PB_FIELD(  1, UINT32  , SINGULAR, STATIC  , FIRST, fk_app_FileData, page, page, 0),
    PB_FIELD(  2, BYTES   , SINGULAR, CALLBACK, OTHER, fk_app_FileData, data, page, 0),
    PB_FIELD(  3, UINT32  , SINGULAR, STATIC  , OTHER, fk_app_FileData, size, data, 0),
    PB_FIELD(  4, UINT32  , SINGULAR, STATIC  , OTHER, fk_app_FileData, hash, size, 0),
    PB_FIELD(  5, BYTES   , SINGULAR, CALLBACK, OTHER, fk_app_FileData, token, hash, 0),
    PB_LAST_FIELD
};

const pb_field_t fk_app_DeviceStatus_fields[6] = {
    PB_FIELD(  1, UINT32  , SINGULAR, STATIC  , FIRST, fk_app_DeviceStatus, uptime, uptime, 0),
    PB_FIELD(  2, FLOAT   , SINGULAR, STATIC  , OTHER, fk_app_DeviceStatus, batteryPercentage, uptime, 0),
    PB_FIELD(  3, FLOAT   , SINGULAR, STATIC  , OTHER, fk_app_DeviceStatus, batteryVoltage, batteryPercentage, 0),
    PB_FIELD(  4, UINT32  , SINGULAR, STATIC  , OTHER, fk_app_DeviceStatus, gpsHasFix, batteryVoltage, 0),
    PB_FIELD(  5, UINT32  , SINGULAR, STATIC  , OTHER, fk_app_DeviceStatus, gpsSatellites, gpsHasFix, 0),
    PB_LAST_FIELD
};

const pb_field_t fk_app_WireMessageQuery_fields[14] = {
    PB_FIELD(  1, UENUM   , SINGULAR, STATIC  , FIRST, fk_app_WireMessageQuery, type, type, 0),
    PB_FIELD(  2, MESSAGE , SINGULAR, STATIC  , OTHER, fk_app_WireMessageQuery, queryCapabilities, type, &fk_app_QueryCapabilities_fields),
    PB_FIELD(  3, MESSAGE , SINGULAR, STATIC  , OTHER, fk_app_WireMessageQuery, configureSensor, queryCapabilities, &fk_app_ConfigureSensorQuery_fields),
    PB_FIELD(  4, MESSAGE , SINGULAR, STATIC  , OTHER, fk_app_WireMessageQuery, queryDataSets, configureSensor, &fk_app_QueryDataSets_fields),
    PB_FIELD(  5, MESSAGE , SINGULAR, STATIC  , OTHER, fk_app_WireMessageQuery, queryDataSet, queryDataSets, &fk_app_QueryDataSet_fields),
    PB_FIELD(  6, MESSAGE , SINGULAR, STATIC  , OTHER, fk_app_WireMessageQuery, downloadDataSet, queryDataSet, &fk_app_DownloadDataSet_fields),
    PB_FIELD(  7, MESSAGE , SINGULAR, STATIC  , OTHER, fk_app_WireMessageQuery, eraseDataSet, downloadDataSet, &fk_app_EraseDataSet_fields),
    PB_FIELD(  8, MESSAGE , SINGULAR, STATIC  , OTHER, fk_app_WireMessageQuery, liveDataPoll, eraseDataSet, &fk_app_LiveDataPoll_fields),
    PB_FIELD(  9, MESSAGE , SINGULAR, STATIC  , OTHER, fk_app_WireMessageQuery, newSchedules, liveDataPoll, &fk_app_Schedules_fields),
    PB_FIELD( 10, MESSAGE , SINGULAR, STATIC  , OTHER, fk_app_WireMessageQuery, downloadFile, newSchedules, &fk_app_DownloadFile_fields),
    PB_FIELD( 11, MESSAGE , SINGULAR, STATIC  , OTHER, fk_app_WireMessageQuery, eraseFile, downloadFile, &fk_app_EraseFile_fields),
    PB_FIELD( 12, MESSAGE , SINGULAR, STATIC  , OTHER, fk_app_WireMessageQuery, networkSettings, eraseFile, &fk_app_NetworkSettings_fields),
    PB_FIELD( 13, MESSAGE , SINGULAR, STATIC  , OTHER, fk_app_WireMessageQuery, identity, networkSettings, &fk_app_Identity_fields),
    PB_LAST_FIELD
};

const pb_field_t fk_app_Error_fields[2] = {
    PB_FIELD(  1, STRING  , SINGULAR, CALLBACK, FIRST, fk_app_Error, message, message, 0),
    PB_LAST_FIELD
};

const pb_field_t fk_app_WireMessageReply_fields[13] = {
    PB_FIELD(  1, UENUM   , SINGULAR, STATIC  , FIRST, fk_app_WireMessageReply, type, type, 0),
    PB_FIELD(  2, MESSAGE , REPEATED, CALLBACK, OTHER, fk_app_WireMessageReply, errors, type, &fk_app_Error_fields),
    PB_FIELD(  3, MESSAGE , SINGULAR, STATIC  , OTHER, fk_app_WireMessageReply, capabilities, errors, &fk_app_Capabilities_fields),
    PB_FIELD(  4, MESSAGE , SINGULAR, STATIC  , OTHER, fk_app_WireMessageReply, dataSets, capabilities, &fk_app_DataSets_fields),
    PB_FIELD(  5, MESSAGE , SINGULAR, STATIC  , OTHER, fk_app_WireMessageReply, dataSetData, dataSets, &fk_app_DataSetData_fields),
    PB_FIELD(  6, MESSAGE , SINGULAR, STATIC  , OTHER, fk_app_WireMessageReply, liveData, dataSetData, &fk_app_LiveData_fields),
    PB_FIELD(  7, MESSAGE , SINGULAR, STATIC  , OTHER, fk_app_WireMessageReply, schedules, liveData, &fk_app_Schedules_fields),
    PB_FIELD(  8, MESSAGE , SINGULAR, STATIC  , OTHER, fk_app_WireMessageReply, files, schedules, &fk_app_Files_fields),
    PB_FIELD(  9, MESSAGE , SINGULAR, STATIC  , OTHER, fk_app_WireMessageReply, fileData, files, &fk_app_FileData_fields),
    PB_FIELD( 10, MESSAGE , SINGULAR, STATIC  , OTHER, fk_app_WireMessageReply, networkSettings, fileData, &fk_app_NetworkSettings_fields),
    PB_FIELD( 11, MESSAGE , SINGULAR, STATIC  , OTHER, fk_app_WireMessageReply, identity, networkSettings, &fk_app_Identity_fields),
    PB_FIELD( 12, MESSAGE , SINGULAR, STATIC  , OTHER, fk_app_WireMessageReply, status, identity, &fk_app_DeviceStatus_fields),
    PB_LAST_FIELD
};




/* Check that field information fits in pb_field_t */
#if !defined(PB_FIELD_32BIT)
/* If you get an error here, it means that you need to define PB_FIELD_32BIT
 * compile-time option. You can do that in pb.h or on compiler command line.
 * 
 * The reason you need to do this is that some of your messages contain tag
 * numbers or field sizes that are larger than what can fit in 8 or 16 bit
 * field descriptors.
 */
PB_STATIC_ASSERT((pb_membersize(fk_app_Schedule, second) < 65536 && pb_membersize(fk_app_Schedule, minute) < 65536 && pb_membersize(fk_app_Schedule, hour) < 65536 && pb_membersize(fk_app_Schedule, day) < 65536 && pb_membersize(fk_app_Schedules, readings) < 65536 && pb_membersize(fk_app_Schedules, transmission) < 65536 && pb_membersize(fk_app_Schedules, status) < 65536 && pb_membersize(fk_app_Schedules, location) < 65536 && pb_membersize(fk_app_WireMessageQuery, queryCapabilities) < 65536 && pb_membersize(fk_app_WireMessageQuery, configureSensor) < 65536 && pb_membersize(fk_app_WireMessageQuery, queryDataSets) < 65536 && pb_membersize(fk_app_WireMessageQuery, queryDataSet) < 65536 && pb_membersize(fk_app_WireMessageQuery, downloadDataSet) < 65536 && pb_membersize(fk_app_WireMessageQuery, eraseDataSet) < 65536 && pb_membersize(fk_app_WireMessageQuery, liveDataPoll) < 65536 && pb_membersize(fk_app_WireMessageQuery, newSchedules) < 65536 && pb_membersize(fk_app_WireMessageQuery, downloadFile) < 65536 && pb_membersize(fk_app_WireMessageQuery, eraseFile) < 65536 && pb_membersize(fk_app_WireMessageQuery, networkSettings) < 65536 && pb_membersize(fk_app_WireMessageQuery, identity) < 65536 && pb_membersize(fk_app_WireMessageReply, capabilities) < 65536 && pb_membersize(fk_app_WireMessageReply, dataSets) < 65536 && pb_membersize(fk_app_WireMessageReply, dataSetData) < 65536 && pb_membersize(fk_app_WireMessageReply, liveData) < 65536 && pb_membersize(fk_app_WireMessageReply, schedules) < 65536 && pb_membersize(fk_app_WireMessageReply, files) < 65536 && pb_membersize(fk_app_WireMessageReply, fileData) < 65536 && pb_membersize(fk_app_WireMessageReply, networkSettings) < 65536 && pb_membersize(fk_app_WireMessageReply, identity) < 65536 && pb_membersize(fk_app_WireMessageReply, status) < 65536), YOU_MUST_DEFINE_PB_FIELD_32BIT_FOR_MESSAGES_fk_app_QueryCapabilities_fk_app_ModuleCapabilities_fk_app_SensorCapabilities_fk_app_Capabilities_fk_app_TimeSpec_fk_app_Schedule_fk_app_Schedules_fk_app_NetworkInfo_fk_app_NetworkSettings_fk_app_Identity_fk_app_ConfigureSensorQuery_fk_app_QueryDataSets_fk_app_DataSet_fk_app_DataSets_fk_app_DownloadDataSet_fk_app_Sample_fk_app_DataSetData_fk_app_LiveDataPoll_fk_app_LiveDataSample_fk_app_LiveData_fk_app_QueryDataSet_fk_app_EraseDataSet_fk_app_File_fk_app_Files_fk_app_DownloadFile_fk_app_EraseFile_fk_app_FileData_fk_app_DeviceStatus_fk_app_WireMessageQuery_fk_app_Error_fk_app_WireMessageReply)
#endif

#if !defined(PB_FIELD_16BIT) && !defined(PB_FIELD_32BIT)
/* If you get an error here, it means that you need to define PB_FIELD_16BIT
 * compile-time option. You can do that in pb.h or on compiler command line.
 * 
 * The reason you need to do this is that some of your messages contain tag
 * numbers or field sizes that are larger than what can fit in the default
 * 8 bit descriptors.
 */
PB_STATIC_ASSERT((pb_membersize(fk_app_Schedule, second) < 256 && pb_membersize(fk_app_Schedule, minute) < 256 && pb_membersize(fk_app_Schedule, hour) < 256 && pb_membersize(fk_app_Schedule, day) < 256 && pb_membersize(fk_app_Schedules, readings) < 256 && pb_membersize(fk_app_Schedules, transmission) < 256 && pb_membersize(fk_app_Schedules, status) < 256 && pb_membersize(fk_app_Schedules, location) < 256 && pb_membersize(fk_app_WireMessageQuery, queryCapabilities) < 256 && pb_membersize(fk_app_WireMessageQuery, configureSensor) < 256 && pb_membersize(fk_app_WireMessageQuery, queryDataSets) < 256 && pb_membersize(fk_app_WireMessageQuery, queryDataSet) < 256 && pb_membersize(fk_app_WireMessageQuery, downloadDataSet) < 256 && pb_membersize(fk_app_WireMessageQuery, eraseDataSet) < 256 && pb_membersize(fk_app_WireMessageQuery, liveDataPoll) < 256 && pb_membersize(fk_app_WireMessageQuery, newSchedules) < 256 && pb_membersize(fk_app_WireMessageQuery, downloadFile) < 256 && pb_membersize(fk_app_WireMessageQuery, eraseFile) < 256 && pb_membersize(fk_app_WireMessageQuery, networkSettings) < 256 && pb_membersize(fk_app_WireMessageQuery, identity) < 256 && pb_membersize(fk_app_WireMessageReply, capabilities) < 256 && pb_membersize(fk_app_WireMessageReply, dataSets) < 256 && pb_membersize(fk_app_WireMessageReply, dataSetData) < 256 && pb_membersize(fk_app_WireMessageReply, liveData) < 256 && pb_membersize(fk_app_WireMessageReply, schedules) < 256 && pb_membersize(fk_app_WireMessageReply, files) < 256 && pb_membersize(fk_app_WireMessageReply, fileData) < 256 && pb_membersize(fk_app_WireMessageReply, networkSettings) < 256 && pb_membersize(fk_app_WireMessageReply, identity) < 256 && pb_membersize(fk_app_WireMessageReply, status) < 256), YOU_MUST_DEFINE_PB_FIELD_16BIT_FOR_MESSAGES_fk_app_QueryCapabilities_fk_app_ModuleCapabilities_fk_app_SensorCapabilities_fk_app_Capabilities_fk_app_TimeSpec_fk_app_Schedule_fk_app_Schedules_fk_app_NetworkInfo_fk_app_NetworkSettings_fk_app_Identity_fk_app_ConfigureSensorQuery_fk_app_QueryDataSets_fk_app_DataSet_fk_app_DataSets_fk_app_DownloadDataSet_fk_app_Sample_fk_app_DataSetData_fk_app_LiveDataPoll_fk_app_LiveDataSample_fk_app_LiveData_fk_app_QueryDataSet_fk_app_EraseDataSet_fk_app_File_fk_app_Files_fk_app_DownloadFile_fk_app_EraseFile_fk_app_FileData_fk_app_DeviceStatus_fk_app_WireMessageQuery_fk_app_Error_fk_app_WireMessageReply)
#endif


/* @@protoc_insertion_point(eof) */
