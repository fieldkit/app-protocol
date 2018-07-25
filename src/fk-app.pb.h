/* Automatically generated nanopb header */
/* Generated by nanopb-0.3.9-dev at Tue Jul 24 19:28:07 2018. */

#ifndef PB_FK_APP_FK_APP_PB_H_INCLUDED
#define PB_FK_APP_FK_APP_PB_H_INCLUDED
#include <pb.h>

/* @@protoc_insertion_point(includes) */
#if PB_PROTO_HEADER_VERSION != 30
#error Regenerate this file with the current version of nanopb generator.
#endif

#ifdef __cplusplus
extern "C" {
#endif

/* Enum definitions */
typedef enum _fk_app_QueryType {
    fk_app_QueryType_QUERY_NONE = 0,
    fk_app_QueryType_QUERY_CAPABILITIES = 1,
    fk_app_QueryType_QUERY_CONFIGURE_SENSOR = 2,
    fk_app_QueryType_QUERY_LIVE_DATA_POLL = 7,
    fk_app_QueryType_QUERY_SCHEDULES = 8,
    fk_app_QueryType_QUERY_CONFIGUE_SCHEDULES = 9,
    fk_app_QueryType_QUERY_FILES = 10,
    fk_app_QueryType_QUERY_DOWNLOAD_FILE = 11,
    fk_app_QueryType_QUERY_ERASE_FILE = 12,
    fk_app_QueryType_QUERY_RESET = 13,
    fk_app_QueryType_QUERY_NETWORK_SETTINGS = 14,
    fk_app_QueryType_QUERY_CONFIGURE_NETWORK_SETTINGS = 15,
    fk_app_QueryType_QUERY_IDENTITY = 16,
    fk_app_QueryType_QUERY_CONFIGURE_IDENTITY = 17,
    fk_app_QueryType_QUERY_STATUS = 18,
    fk_app_QueryType_QUERY_MODULE = 19,
    fk_app_QueryType_QUERY_METADATA = 20,
    fk_app_QueryType_QUERY_FORMAT = 21
} fk_app_QueryType;
#define _fk_app_QueryType_MIN fk_app_QueryType_QUERY_NONE
#define _fk_app_QueryType_MAX fk_app_QueryType_QUERY_FORMAT
#define _fk_app_QueryType_ARRAYSIZE ((fk_app_QueryType)(fk_app_QueryType_QUERY_FORMAT+1))

typedef enum _fk_app_ReplyType {
    fk_app_ReplyType_REPLY_NONE = 0,
    fk_app_ReplyType_REPLY_SUCCESS = 1,
    fk_app_ReplyType_REPLY_BUSY = 2,
    fk_app_ReplyType_REPLY_ERROR = 3,
    fk_app_ReplyType_REPLY_CAPABILITIES = 4,
    fk_app_ReplyType_REPLY_LIVE_DATA_POLL = 8,
    fk_app_ReplyType_REPLY_SCHEDULES = 9,
    fk_app_ReplyType_REPLY_FILES = 10,
    fk_app_ReplyType_REPLY_DOWNLOAD_FILE = 11,
    fk_app_ReplyType_REPLY_RESET = 12,
    fk_app_ReplyType_REPLY_NETWORK_SETTINGS = 13,
    fk_app_ReplyType_REPLY_IDENTITY = 14,
    fk_app_ReplyType_REPLY_STATUS = 15,
    fk_app_ReplyType_REPLY_MODULE = 16,
    fk_app_ReplyType_REPLY_METADATA = 17
} fk_app_ReplyType;
#define _fk_app_ReplyType_MIN fk_app_ReplyType_REPLY_NONE
#define _fk_app_ReplyType_MAX fk_app_ReplyType_REPLY_METADATA
#define _fk_app_ReplyType_ARRAYSIZE ((fk_app_ReplyType)(fk_app_ReplyType_REPLY_METADATA+1))

typedef enum _fk_app_DownloadFlags {
    fk_app_DownloadFlags_DOWNLOAD_FLAG_NONE = 0,
    fk_app_DownloadFlags_DOWNLOAD_FLAG_METADATA_PREPEND = 1,
    fk_app_DownloadFlags_DOWNLOAD_FLAG_METADATA_ONLY = 2
} fk_app_DownloadFlags;
#define _fk_app_DownloadFlags_MIN fk_app_DownloadFlags_DOWNLOAD_FLAG_NONE
#define _fk_app_DownloadFlags_MAX fk_app_DownloadFlags_DOWNLOAD_FLAG_METADATA_ONLY
#define _fk_app_DownloadFlags_ARRAYSIZE ((fk_app_DownloadFlags)(fk_app_DownloadFlags_DOWNLOAD_FLAG_METADATA_ONLY+1))

/* Struct definitions */
typedef struct _fk_app_Error {
    pb_callback_t message;
/* @@protoc_insertion_point(struct:fk_app_Error) */
} fk_app_Error;

typedef struct _fk_app_Files {
    pb_callback_t files;
/* @@protoc_insertion_point(struct:fk_app_Files) */
} fk_app_Files;

typedef struct _fk_app_Identity {
    pb_callback_t device;
    pb_callback_t stream;
    pb_callback_t deviceId;
/* @@protoc_insertion_point(struct:fk_app_Identity) */
} fk_app_Identity;

typedef struct _fk_app_LiveData {
    pb_callback_t samples;
/* @@protoc_insertion_point(struct:fk_app_LiveData) */
} fk_app_LiveData;

typedef struct _fk_app_NetworkInfo {
    pb_callback_t ssid;
    pb_callback_t password;
/* @@protoc_insertion_point(struct:fk_app_NetworkInfo) */
} fk_app_NetworkInfo;

typedef struct _fk_app_Schedule {
    char dummy_field;
/* @@protoc_insertion_point(struct:fk_app_Schedule) */
} fk_app_Schedule;

typedef struct _fk_app_Capabilities {
    uint32_t version;
    pb_callback_t deviceId;
    pb_callback_t name;
    pb_callback_t modules;
    pb_callback_t sensors;
/* @@protoc_insertion_point(struct:fk_app_Capabilities) */
} fk_app_Capabilities;

typedef struct _fk_app_ConfigureSensorQuery {
    uint32_t id;
    uint32_t frequency;
/* @@protoc_insertion_point(struct:fk_app_ConfigureSensorQuery) */
} fk_app_ConfigureSensorQuery;

typedef struct _fk_app_DeviceStatus {
    uint32_t uptime;
    float batteryPercentage;
    float batteryVoltage;
    uint32_t gpsHasFix;
    uint32_t gpsSatellites;
/* @@protoc_insertion_point(struct:fk_app_DeviceStatus) */
} fk_app_DeviceStatus;

typedef struct _fk_app_DownloadFile {
    uint32_t id;
    uint32_t offset;
    uint32_t length;
    uint32_t flags;
/* @@protoc_insertion_point(struct:fk_app_DownloadFile) */
} fk_app_DownloadFile;

typedef struct _fk_app_EraseFile {
    uint32_t id;
/* @@protoc_insertion_point(struct:fk_app_EraseFile) */
} fk_app_EraseFile;

typedef struct _fk_app_File {
    uint32_t id;
    uint64_t time;
    uint64_t size;
    uint32_t version;
    pb_callback_t name;
    uint64_t maximum;
/* @@protoc_insertion_point(struct:fk_app_File) */
} fk_app_File;

typedef struct _fk_app_FileData {
    uint32_t offset;
    pb_callback_t data;
    uint32_t size;
    uint32_t hash;
    uint32_t version;
    uint32_t id;
/* @@protoc_insertion_point(struct:fk_app_FileData) */
} fk_app_FileData;

typedef struct _fk_app_LiveDataPoll {
    uint32_t interval;
/* @@protoc_insertion_point(struct:fk_app_LiveDataPoll) */
} fk_app_LiveDataPoll;

typedef struct _fk_app_LiveDataSample {
    uint32_t sensor;
    uint64_t time;
    float value;
/* @@protoc_insertion_point(struct:fk_app_LiveDataSample) */
} fk_app_LiveDataSample;

typedef struct _fk_app_ModuleCapabilities {
    uint32_t id;
    pb_callback_t name;
/* @@protoc_insertion_point(struct:fk_app_ModuleCapabilities) */
} fk_app_ModuleCapabilities;

typedef struct _fk_app_ModuleReply {
    uint32_t id;
    uint32_t address;
    pb_callback_t message;
/* @@protoc_insertion_point(struct:fk_app_ModuleReply) */
} fk_app_ModuleReply;

typedef struct _fk_app_NetworkSettings {
    int32_t createAccessPoint;
    pb_callback_t networks;
/* @@protoc_insertion_point(struct:fk_app_NetworkSettings) */
} fk_app_NetworkSettings;

typedef struct _fk_app_QueryCapabilities {
    uint32_t version;
    uint32_t callerTime;
/* @@protoc_insertion_point(struct:fk_app_QueryCapabilities) */
} fk_app_QueryCapabilities;

typedef struct _fk_app_QueryModule {
    uint32_t id;
    uint32_t address;
    pb_callback_t message;
/* @@protoc_insertion_point(struct:fk_app_QueryModule) */
} fk_app_QueryModule;

typedef struct _fk_app_Schedules {
    fk_app_Schedule readings;
    fk_app_Schedule transmission;
    fk_app_Schedule status;
    fk_app_Schedule location;
/* @@protoc_insertion_point(struct:fk_app_Schedules) */
} fk_app_Schedules;

typedef struct _fk_app_SensorCapabilities {
    uint32_t id;
    uint32_t module;
    pb_callback_t name;
    uint32_t frequency;
    pb_callback_t unitOfMeasure;
/* @@protoc_insertion_point(struct:fk_app_SensorCapabilities) */
} fk_app_SensorCapabilities;

typedef struct _fk_app_TimeSpec {
    int32_t fixed;
    int32_t interval;
    int32_t offset;
/* @@protoc_insertion_point(struct:fk_app_TimeSpec) */
} fk_app_TimeSpec;

typedef struct _fk_app_WireMessageQuery {
    fk_app_QueryType type;
    fk_app_QueryCapabilities queryCapabilities;
    fk_app_ConfigureSensorQuery configureSensor;
    fk_app_LiveDataPoll liveDataPoll;
    fk_app_Schedules newSchedules;
    fk_app_DownloadFile downloadFile;
    fk_app_EraseFile eraseFile;
    fk_app_NetworkSettings networkSettings;
    fk_app_Identity identity;
    fk_app_QueryModule module;
/* @@protoc_insertion_point(struct:fk_app_WireMessageQuery) */
} fk_app_WireMessageQuery;

typedef struct _fk_app_WireMessageReply {
    fk_app_ReplyType type;
    pb_callback_t errors;
    fk_app_Capabilities capabilities;
    fk_app_LiveData liveData;
    fk_app_Schedules schedules;
    fk_app_Files files;
    fk_app_FileData fileData;
    fk_app_NetworkSettings networkSettings;
    fk_app_Identity identity;
    fk_app_DeviceStatus status;
    fk_app_ModuleReply module;
/* @@protoc_insertion_point(struct:fk_app_WireMessageReply) */
} fk_app_WireMessageReply;

/* Default values for struct fields */

/* Initializer values for message structs */
#define fk_app_QueryCapabilities_init_default    {0, 0}
#define fk_app_ModuleCapabilities_init_default   {0, {{NULL}, NULL}}
#define fk_app_SensorCapabilities_init_default   {0, 0, {{NULL}, NULL}, 0, {{NULL}, NULL}}
#define fk_app_Capabilities_init_default         {0, {{NULL}, NULL}, {{NULL}, NULL}, {{NULL}, NULL}, {{NULL}, NULL}}
#define fk_app_TimeSpec_init_default             {0, 0, 0}
#define fk_app_Schedule_init_default             {0}
#define fk_app_Schedules_init_default            {fk_app_Schedule_init_default, fk_app_Schedule_init_default, fk_app_Schedule_init_default, fk_app_Schedule_init_default}
#define fk_app_NetworkInfo_init_default          {{{NULL}, NULL}, {{NULL}, NULL}}
#define fk_app_NetworkSettings_init_default      {0, {{NULL}, NULL}}
#define fk_app_Identity_init_default             {{{NULL}, NULL}, {{NULL}, NULL}, {{NULL}, NULL}}
#define fk_app_ConfigureSensorQuery_init_default {0, 0}
#define fk_app_LiveDataPoll_init_default         {0}
#define fk_app_LiveDataSample_init_default       {0, 0, 0}
#define fk_app_LiveData_init_default             {{{NULL}, NULL}}
#define fk_app_File_init_default                 {0, 0, 0, 0, {{NULL}, NULL}, 0}
#define fk_app_Files_init_default                {{{NULL}, NULL}}
#define fk_app_DownloadFile_init_default         {0, 0, 0, 0}
#define fk_app_EraseFile_init_default            {0}
#define fk_app_FileData_init_default             {0, {{NULL}, NULL}, 0, 0, 0, 0}
#define fk_app_DeviceStatus_init_default         {0, 0, 0, 0, 0}
#define fk_app_QueryModule_init_default          {0, 0, {{NULL}, NULL}}
#define fk_app_ModuleReply_init_default          {0, 0, {{NULL}, NULL}}
#define fk_app_WireMessageQuery_init_default     {(fk_app_QueryType)0, fk_app_QueryCapabilities_init_default, fk_app_ConfigureSensorQuery_init_default, fk_app_LiveDataPoll_init_default, fk_app_Schedules_init_default, fk_app_DownloadFile_init_default, fk_app_EraseFile_init_default, fk_app_NetworkSettings_init_default, fk_app_Identity_init_default, fk_app_QueryModule_init_default}
#define fk_app_Error_init_default                {{{NULL}, NULL}}
#define fk_app_WireMessageReply_init_default     {(fk_app_ReplyType)0, {{NULL}, NULL}, fk_app_Capabilities_init_default, fk_app_LiveData_init_default, fk_app_Schedules_init_default, fk_app_Files_init_default, fk_app_FileData_init_default, fk_app_NetworkSettings_init_default, fk_app_Identity_init_default, fk_app_DeviceStatus_init_default, fk_app_ModuleReply_init_default}
#define fk_app_QueryCapabilities_init_zero       {0, 0}
#define fk_app_ModuleCapabilities_init_zero      {0, {{NULL}, NULL}}
#define fk_app_SensorCapabilities_init_zero      {0, 0, {{NULL}, NULL}, 0, {{NULL}, NULL}}
#define fk_app_Capabilities_init_zero            {0, {{NULL}, NULL}, {{NULL}, NULL}, {{NULL}, NULL}, {{NULL}, NULL}}
#define fk_app_TimeSpec_init_zero                {0, 0, 0}
#define fk_app_Schedule_init_zero                {0}
#define fk_app_Schedules_init_zero               {fk_app_Schedule_init_zero, fk_app_Schedule_init_zero, fk_app_Schedule_init_zero, fk_app_Schedule_init_zero}
#define fk_app_NetworkInfo_init_zero             {{{NULL}, NULL}, {{NULL}, NULL}}
#define fk_app_NetworkSettings_init_zero         {0, {{NULL}, NULL}}
#define fk_app_Identity_init_zero                {{{NULL}, NULL}, {{NULL}, NULL}, {{NULL}, NULL}}
#define fk_app_ConfigureSensorQuery_init_zero    {0, 0}
#define fk_app_LiveDataPoll_init_zero            {0}
#define fk_app_LiveDataSample_init_zero          {0, 0, 0}
#define fk_app_LiveData_init_zero                {{{NULL}, NULL}}
#define fk_app_File_init_zero                    {0, 0, 0, 0, {{NULL}, NULL}, 0}
#define fk_app_Files_init_zero                   {{{NULL}, NULL}}
#define fk_app_DownloadFile_init_zero            {0, 0, 0, 0}
#define fk_app_EraseFile_init_zero               {0}
#define fk_app_FileData_init_zero                {0, {{NULL}, NULL}, 0, 0, 0, 0}
#define fk_app_DeviceStatus_init_zero            {0, 0, 0, 0, 0}
#define fk_app_QueryModule_init_zero             {0, 0, {{NULL}, NULL}}
#define fk_app_ModuleReply_init_zero             {0, 0, {{NULL}, NULL}}
#define fk_app_WireMessageQuery_init_zero        {(fk_app_QueryType)0, fk_app_QueryCapabilities_init_zero, fk_app_ConfigureSensorQuery_init_zero, fk_app_LiveDataPoll_init_zero, fk_app_Schedules_init_zero, fk_app_DownloadFile_init_zero, fk_app_EraseFile_init_zero, fk_app_NetworkSettings_init_zero, fk_app_Identity_init_zero, fk_app_QueryModule_init_zero}
#define fk_app_Error_init_zero                   {{{NULL}, NULL}}
#define fk_app_WireMessageReply_init_zero        {(fk_app_ReplyType)0, {{NULL}, NULL}, fk_app_Capabilities_init_zero, fk_app_LiveData_init_zero, fk_app_Schedules_init_zero, fk_app_Files_init_zero, fk_app_FileData_init_zero, fk_app_NetworkSettings_init_zero, fk_app_Identity_init_zero, fk_app_DeviceStatus_init_zero, fk_app_ModuleReply_init_zero}

/* Field tags (for use in manual encoding/decoding) */
#define fk_app_Error_message_tag                 1
#define fk_app_Files_files_tag                   1
#define fk_app_Identity_device_tag               1
#define fk_app_Identity_stream_tag               2
#define fk_app_Identity_deviceId_tag             3
#define fk_app_LiveData_samples_tag              1
#define fk_app_NetworkInfo_ssid_tag              1
#define fk_app_NetworkInfo_password_tag          2
#define fk_app_Capabilities_version_tag          1
#define fk_app_Capabilities_deviceId_tag         2
#define fk_app_Capabilities_name_tag             3
#define fk_app_Capabilities_modules_tag          4
#define fk_app_Capabilities_sensors_tag          5
#define fk_app_ConfigureSensorQuery_id_tag       1
#define fk_app_ConfigureSensorQuery_frequency_tag 2
#define fk_app_DeviceStatus_uptime_tag           1
#define fk_app_DeviceStatus_batteryPercentage_tag 2
#define fk_app_DeviceStatus_batteryVoltage_tag   3
#define fk_app_DeviceStatus_gpsHasFix_tag        4
#define fk_app_DeviceStatus_gpsSatellites_tag    5
#define fk_app_DownloadFile_id_tag               1
#define fk_app_DownloadFile_offset_tag           2
#define fk_app_DownloadFile_length_tag           3
#define fk_app_DownloadFile_flags_tag            4
#define fk_app_EraseFile_id_tag                  1
#define fk_app_File_id_tag                       1
#define fk_app_File_time_tag                     2
#define fk_app_File_size_tag                     3
#define fk_app_File_version_tag                  4
#define fk_app_File_name_tag                     5
#define fk_app_File_maximum_tag                  6
#define fk_app_FileData_offset_tag               1
#define fk_app_FileData_data_tag                 2
#define fk_app_FileData_size_tag                 3
#define fk_app_FileData_hash_tag                 4
#define fk_app_FileData_version_tag              5
#define fk_app_FileData_id_tag                   6
#define fk_app_LiveDataPoll_interval_tag         1
#define fk_app_LiveDataSample_sensor_tag         1
#define fk_app_LiveDataSample_time_tag           2
#define fk_app_LiveDataSample_value_tag          3
#define fk_app_ModuleCapabilities_id_tag         1
#define fk_app_ModuleCapabilities_name_tag       2
#define fk_app_ModuleReply_id_tag                1
#define fk_app_ModuleReply_address_tag           2
#define fk_app_ModuleReply_message_tag           3
#define fk_app_NetworkSettings_createAccessPoint_tag 1
#define fk_app_NetworkSettings_networks_tag      2
#define fk_app_QueryCapabilities_version_tag     1
#define fk_app_QueryCapabilities_callerTime_tag  2
#define fk_app_QueryModule_id_tag                1
#define fk_app_QueryModule_address_tag           2
#define fk_app_QueryModule_message_tag           3
#define fk_app_Schedules_readings_tag            1
#define fk_app_Schedules_transmission_tag        2
#define fk_app_Schedules_status_tag              3
#define fk_app_Schedules_location_tag            4
#define fk_app_SensorCapabilities_id_tag         1
#define fk_app_SensorCapabilities_module_tag     2
#define fk_app_SensorCapabilities_name_tag       3
#define fk_app_SensorCapabilities_frequency_tag  4
#define fk_app_SensorCapabilities_unitOfMeasure_tag 5
#define fk_app_TimeSpec_fixed_tag                1
#define fk_app_TimeSpec_interval_tag             2
#define fk_app_TimeSpec_offset_tag               3
#define fk_app_WireMessageQuery_type_tag         1
#define fk_app_WireMessageQuery_queryCapabilities_tag 2
#define fk_app_WireMessageQuery_configureSensor_tag 3
#define fk_app_WireMessageQuery_liveDataPoll_tag 8
#define fk_app_WireMessageQuery_newSchedules_tag 9
#define fk_app_WireMessageQuery_downloadFile_tag 10
#define fk_app_WireMessageQuery_eraseFile_tag    11
#define fk_app_WireMessageQuery_networkSettings_tag 12
#define fk_app_WireMessageQuery_identity_tag     13
#define fk_app_WireMessageQuery_module_tag       14
#define fk_app_WireMessageReply_type_tag         1
#define fk_app_WireMessageReply_errors_tag       2
#define fk_app_WireMessageReply_capabilities_tag 3
#define fk_app_WireMessageReply_liveData_tag     6
#define fk_app_WireMessageReply_schedules_tag    7
#define fk_app_WireMessageReply_files_tag        8
#define fk_app_WireMessageReply_fileData_tag     9
#define fk_app_WireMessageReply_networkSettings_tag 10
#define fk_app_WireMessageReply_identity_tag     11
#define fk_app_WireMessageReply_status_tag       12
#define fk_app_WireMessageReply_module_tag       13

/* Struct field encoding specification for nanopb */
extern const pb_field_t fk_app_QueryCapabilities_fields[3];
extern const pb_field_t fk_app_ModuleCapabilities_fields[3];
extern const pb_field_t fk_app_SensorCapabilities_fields[6];
extern const pb_field_t fk_app_Capabilities_fields[6];
extern const pb_field_t fk_app_TimeSpec_fields[4];
extern const pb_field_t fk_app_Schedule_fields[1];
extern const pb_field_t fk_app_Schedules_fields[5];
extern const pb_field_t fk_app_NetworkInfo_fields[3];
extern const pb_field_t fk_app_NetworkSettings_fields[3];
extern const pb_field_t fk_app_Identity_fields[4];
extern const pb_field_t fk_app_ConfigureSensorQuery_fields[3];
extern const pb_field_t fk_app_LiveDataPoll_fields[2];
extern const pb_field_t fk_app_LiveDataSample_fields[4];
extern const pb_field_t fk_app_LiveData_fields[2];
extern const pb_field_t fk_app_File_fields[7];
extern const pb_field_t fk_app_Files_fields[2];
extern const pb_field_t fk_app_DownloadFile_fields[5];
extern const pb_field_t fk_app_EraseFile_fields[2];
extern const pb_field_t fk_app_FileData_fields[7];
extern const pb_field_t fk_app_DeviceStatus_fields[6];
extern const pb_field_t fk_app_QueryModule_fields[4];
extern const pb_field_t fk_app_ModuleReply_fields[4];
extern const pb_field_t fk_app_WireMessageQuery_fields[11];
extern const pb_field_t fk_app_Error_fields[2];
extern const pb_field_t fk_app_WireMessageReply_fields[12];

/* Maximum encoded size of messages (where known) */
#define fk_app_QueryCapabilities_size            12
/* fk_app_ModuleCapabilities_size depends on runtime parameters */
/* fk_app_SensorCapabilities_size depends on runtime parameters */
/* fk_app_Capabilities_size depends on runtime parameters */
#define fk_app_TimeSpec_size                     33
#define fk_app_Schedule_size                     0
#define fk_app_Schedules_size                    8
/* fk_app_NetworkInfo_size depends on runtime parameters */
/* fk_app_NetworkSettings_size depends on runtime parameters */
/* fk_app_Identity_size depends on runtime parameters */
#define fk_app_ConfigureSensorQuery_size         12
#define fk_app_LiveDataPoll_size                 6
#define fk_app_LiveDataSample_size               22
/* fk_app_LiveData_size depends on runtime parameters */
/* fk_app_File_size depends on runtime parameters */
/* fk_app_Files_size depends on runtime parameters */
#define fk_app_DownloadFile_size                 24
#define fk_app_EraseFile_size                    6
/* fk_app_FileData_size depends on runtime parameters */
#define fk_app_DeviceStatus_size                 28
/* fk_app_QueryModule_size depends on runtime parameters */
/* fk_app_ModuleReply_size depends on runtime parameters */
#define fk_app_WireMessageQuery_size             (100 + fk_app_NetworkSettings_size + fk_app_Identity_size + fk_app_QueryModule_size)
/* fk_app_Error_size depends on runtime parameters */
/* fk_app_WireMessageReply_size depends on runtime parameters */

/* Message IDs (where set with "msgid" option) */
#ifdef PB_MSGID

#define FK_APP_MESSAGES \


#endif

#ifdef __cplusplus
} /* extern "C" */
#endif
/* @@protoc_insertion_point(eof) */

#endif
