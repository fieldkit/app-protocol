/* Automatically generated nanopb header */
/* Generated by nanopb-0.3.9-dev at Fri Dec  8 16:48:47 2017. */

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
    fk_app_QueryType_QUERY_DATA_SETS = 3,
    fk_app_QueryType_QUERY_DATA_SET = 4,
    fk_app_QueryType_QUERY_DOWNLOAD_DATA_SET = 5,
    fk_app_QueryType_QUERY_ERASE_DATA_SET = 6,
    fk_app_QueryType_QUERY_LIVE_DATA_POLL = 7,
    fk_app_QueryType_QUERY_SCHEDULES = 8,
    fk_app_QueryType_QUERY_CONFIGUE_SCHEDULES = 9,
    fk_app_QueryType_QUERY_FILES = 10,
    fk_app_QueryType_QUERY_DOWNLOAD_FILE = 11,
    fk_app_QueryType_QUERY_RESET = 12
} fk_app_QueryType;
#define _fk_app_QueryType_MIN fk_app_QueryType_QUERY_NONE
#define _fk_app_QueryType_MAX fk_app_QueryType_QUERY_RESET
#define _fk_app_QueryType_ARRAYSIZE ((fk_app_QueryType)(fk_app_QueryType_QUERY_RESET+1))

typedef enum _fk_app_ReplyType {
    fk_app_ReplyType_REPLY_NONE = 0,
    fk_app_ReplyType_REPLY_SUCCESS = 1,
    fk_app_ReplyType_REPLY_ERROR = 2,
    fk_app_ReplyType_REPLY_CAPABILITIES = 3,
    fk_app_ReplyType_REPLY_DATA_SETS = 4,
    fk_app_ReplyType_REPLY_DATA_SET = 5,
    fk_app_ReplyType_REPLY_DOWNLOAD_DATA_SET = 6,
    fk_app_ReplyType_REPLY_LIVE_DATA_POLL = 7,
    fk_app_ReplyType_REPLY_SCHEDULES = 8,
    fk_app_ReplyType_REPLY_FILES = 9,
    fk_app_ReplyType_REPLY_DOWNLOAD_FILE = 10,
    fk_app_ReplyType_REPLY_RESET = 11
} fk_app_ReplyType;
#define _fk_app_ReplyType_MIN fk_app_ReplyType_REPLY_NONE
#define _fk_app_ReplyType_MAX fk_app_ReplyType_REPLY_RESET
#define _fk_app_ReplyType_ARRAYSIZE ((fk_app_ReplyType)(fk_app_ReplyType_REPLY_RESET+1))

/* Struct definitions */
typedef struct _fk_app_DataSets {
    pb_callback_t dataSets;
/* @@protoc_insertion_point(struct:fk_app_DataSets) */
} fk_app_DataSets;

typedef struct _fk_app_Error {
    pb_callback_t message;
/* @@protoc_insertion_point(struct:fk_app_Error) */
} fk_app_Error;

typedef struct _fk_app_Files {
    pb_callback_t files;
/* @@protoc_insertion_point(struct:fk_app_Files) */
} fk_app_Files;

typedef struct _fk_app_LiveData {
    pb_callback_t samples;
/* @@protoc_insertion_point(struct:fk_app_LiveData) */
} fk_app_LiveData;

typedef struct _fk_app_QueryDataSets {
    char dummy_field;
/* @@protoc_insertion_point(struct:fk_app_QueryDataSets) */
} fk_app_QueryDataSets;

typedef struct _fk_app_Capabilities {
    uint32_t version;
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

typedef struct _fk_app_DataSet {
    uint32_t id;
    uint32_t sensor;
    uint64_t time;
    uint32_t size;
    uint32_t pages;
    uint32_t hash;
    pb_callback_t name;
/* @@protoc_insertion_point(struct:fk_app_DataSet) */
} fk_app_DataSet;

typedef struct _fk_app_DataSetData {
    uint64_t time;
    uint32_t page;
    uint32_t sensor;
    pb_callback_t samples;
    pb_callback_t floats;
    pb_callback_t data;
    uint32_t hash;
/* @@protoc_insertion_point(struct:fk_app_DataSetData) */
} fk_app_DataSetData;

typedef struct _fk_app_DownloadDataSet {
    uint32_t id;
    uint32_t page;
/* @@protoc_insertion_point(struct:fk_app_DownloadDataSet) */
} fk_app_DownloadDataSet;

typedef struct _fk_app_DownloadFile {
    uint32_t id;
    uint32_t page;
    pb_callback_t token;
/* @@protoc_insertion_point(struct:fk_app_DownloadFile) */
} fk_app_DownloadFile;

typedef struct _fk_app_EraseDataSet {
    uint32_t id;
/* @@protoc_insertion_point(struct:fk_app_EraseDataSet) */
} fk_app_EraseDataSet;

typedef struct _fk_app_File {
    uint32_t id;
    uint64_t time;
    uint32_t size;
    uint32_t pages;
    pb_callback_t name;
/* @@protoc_insertion_point(struct:fk_app_File) */
} fk_app_File;

typedef struct _fk_app_FileData {
    uint32_t page;
    pb_callback_t data;
    uint32_t hash;
    pb_callback_t token;
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

typedef struct _fk_app_QueryCapabilities {
    uint32_t version;
    uint32_t callerTime;
/* @@protoc_insertion_point(struct:fk_app_QueryCapabilities) */
} fk_app_QueryCapabilities;

typedef struct _fk_app_QueryDataSet {
    uint32_t id;
/* @@protoc_insertion_point(struct:fk_app_QueryDataSet) */
} fk_app_QueryDataSet;

typedef struct _fk_app_Sample {
    uint64_t time;
    float value;
/* @@protoc_insertion_point(struct:fk_app_Sample) */
} fk_app_Sample;

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

typedef struct _fk_app_Schedule {
    fk_app_TimeSpec second;
    fk_app_TimeSpec minute;
    fk_app_TimeSpec hour;
    fk_app_TimeSpec day;
/* @@protoc_insertion_point(struct:fk_app_Schedule) */
} fk_app_Schedule;

typedef struct _fk_app_Schedules {
    fk_app_Schedule readings;
    fk_app_Schedule transmission;
    fk_app_Schedule status;
    fk_app_Schedule location;
/* @@protoc_insertion_point(struct:fk_app_Schedules) */
} fk_app_Schedules;

typedef struct _fk_app_WireMessageQuery {
    fk_app_QueryType type;
    fk_app_QueryCapabilities queryCapabilities;
    fk_app_ConfigureSensorQuery configureSensor;
    fk_app_QueryDataSets queryDataSets;
    fk_app_QueryDataSet queryDataSet;
    fk_app_DownloadDataSet downloadDataSet;
    fk_app_EraseDataSet eraseDataSet;
    fk_app_LiveDataPoll liveDataPoll;
    fk_app_Schedules newSchedules;
    fk_app_DownloadFile downloadFile;
/* @@protoc_insertion_point(struct:fk_app_WireMessageQuery) */
} fk_app_WireMessageQuery;

typedef struct _fk_app_WireMessageReply {
    fk_app_ReplyType type;
    pb_callback_t errors;
    fk_app_Capabilities capabilities;
    fk_app_DataSets dataSets;
    fk_app_DataSetData dataSetData;
    fk_app_LiveData liveData;
    fk_app_Schedules schedules;
    fk_app_Files files;
    fk_app_FileData fileData;
/* @@protoc_insertion_point(struct:fk_app_WireMessageReply) */
} fk_app_WireMessageReply;

/* Default values for struct fields */

/* Initializer values for message structs */
#define fk_app_QueryCapabilities_init_default    {0, 0}
#define fk_app_ModuleCapabilities_init_default   {0, {{NULL}, NULL}}
#define fk_app_SensorCapabilities_init_default   {0, 0, {{NULL}, NULL}, 0, {{NULL}, NULL}}
#define fk_app_Capabilities_init_default         {0, {{NULL}, NULL}, {{NULL}, NULL}, {{NULL}, NULL}}
#define fk_app_TimeSpec_init_default             {0, 0, 0}
#define fk_app_Schedule_init_default             {fk_app_TimeSpec_init_default, fk_app_TimeSpec_init_default, fk_app_TimeSpec_init_default, fk_app_TimeSpec_init_default}
#define fk_app_Schedules_init_default            {fk_app_Schedule_init_default, fk_app_Schedule_init_default, fk_app_Schedule_init_default, fk_app_Schedule_init_default}
#define fk_app_ConfigureSensorQuery_init_default {0, 0}
#define fk_app_QueryDataSets_init_default        {0}
#define fk_app_DataSet_init_default              {0, 0, 0, 0, 0, 0, {{NULL}, NULL}}
#define fk_app_DataSets_init_default             {{{NULL}, NULL}}
#define fk_app_DownloadDataSet_init_default      {0, 0}
#define fk_app_Sample_init_default               {0, 0}
#define fk_app_DataSetData_init_default          {0, 0, 0, {{NULL}, NULL}, {{NULL}, NULL}, {{NULL}, NULL}, 0}
#define fk_app_LiveDataPoll_init_default         {0}
#define fk_app_LiveDataSample_init_default       {0, 0, 0}
#define fk_app_LiveData_init_default             {{{NULL}, NULL}}
#define fk_app_QueryDataSet_init_default         {0}
#define fk_app_EraseDataSet_init_default         {0}
#define fk_app_File_init_default                 {0, 0, 0, 0, {{NULL}, NULL}}
#define fk_app_Files_init_default                {{{NULL}, NULL}}
#define fk_app_DownloadFile_init_default         {0, 0, {{NULL}, NULL}}
#define fk_app_FileData_init_default             {0, {{NULL}, NULL}, 0, {{NULL}, NULL}}
#define fk_app_WireMessageQuery_init_default     {(fk_app_QueryType)0, fk_app_QueryCapabilities_init_default, fk_app_ConfigureSensorQuery_init_default, fk_app_QueryDataSets_init_default, fk_app_QueryDataSet_init_default, fk_app_DownloadDataSet_init_default, fk_app_EraseDataSet_init_default, fk_app_LiveDataPoll_init_default, fk_app_Schedules_init_default, fk_app_DownloadFile_init_default}
#define fk_app_Error_init_default                {{{NULL}, NULL}}
#define fk_app_WireMessageReply_init_default     {(fk_app_ReplyType)0, {{NULL}, NULL}, fk_app_Capabilities_init_default, fk_app_DataSets_init_default, fk_app_DataSetData_init_default, fk_app_LiveData_init_default, fk_app_Schedules_init_default, fk_app_Files_init_default, fk_app_FileData_init_default}
#define fk_app_QueryCapabilities_init_zero       {0, 0}
#define fk_app_ModuleCapabilities_init_zero      {0, {{NULL}, NULL}}
#define fk_app_SensorCapabilities_init_zero      {0, 0, {{NULL}, NULL}, 0, {{NULL}, NULL}}
#define fk_app_Capabilities_init_zero            {0, {{NULL}, NULL}, {{NULL}, NULL}, {{NULL}, NULL}}
#define fk_app_TimeSpec_init_zero                {0, 0, 0}
#define fk_app_Schedule_init_zero                {fk_app_TimeSpec_init_zero, fk_app_TimeSpec_init_zero, fk_app_TimeSpec_init_zero, fk_app_TimeSpec_init_zero}
#define fk_app_Schedules_init_zero               {fk_app_Schedule_init_zero, fk_app_Schedule_init_zero, fk_app_Schedule_init_zero, fk_app_Schedule_init_zero}
#define fk_app_ConfigureSensorQuery_init_zero    {0, 0}
#define fk_app_QueryDataSets_init_zero           {0}
#define fk_app_DataSet_init_zero                 {0, 0, 0, 0, 0, 0, {{NULL}, NULL}}
#define fk_app_DataSets_init_zero                {{{NULL}, NULL}}
#define fk_app_DownloadDataSet_init_zero         {0, 0}
#define fk_app_Sample_init_zero                  {0, 0}
#define fk_app_DataSetData_init_zero             {0, 0, 0, {{NULL}, NULL}, {{NULL}, NULL}, {{NULL}, NULL}, 0}
#define fk_app_LiveDataPoll_init_zero            {0}
#define fk_app_LiveDataSample_init_zero          {0, 0, 0}
#define fk_app_LiveData_init_zero                {{{NULL}, NULL}}
#define fk_app_QueryDataSet_init_zero            {0}
#define fk_app_EraseDataSet_init_zero            {0}
#define fk_app_File_init_zero                    {0, 0, 0, 0, {{NULL}, NULL}}
#define fk_app_Files_init_zero                   {{{NULL}, NULL}}
#define fk_app_DownloadFile_init_zero            {0, 0, {{NULL}, NULL}}
#define fk_app_FileData_init_zero                {0, {{NULL}, NULL}, 0, {{NULL}, NULL}}
#define fk_app_WireMessageQuery_init_zero        {(fk_app_QueryType)0, fk_app_QueryCapabilities_init_zero, fk_app_ConfigureSensorQuery_init_zero, fk_app_QueryDataSets_init_zero, fk_app_QueryDataSet_init_zero, fk_app_DownloadDataSet_init_zero, fk_app_EraseDataSet_init_zero, fk_app_LiveDataPoll_init_zero, fk_app_Schedules_init_zero, fk_app_DownloadFile_init_zero}
#define fk_app_Error_init_zero                   {{{NULL}, NULL}}
#define fk_app_WireMessageReply_init_zero        {(fk_app_ReplyType)0, {{NULL}, NULL}, fk_app_Capabilities_init_zero, fk_app_DataSets_init_zero, fk_app_DataSetData_init_zero, fk_app_LiveData_init_zero, fk_app_Schedules_init_zero, fk_app_Files_init_zero, fk_app_FileData_init_zero}

/* Field tags (for use in manual encoding/decoding) */
#define fk_app_DataSets_dataSets_tag             1
#define fk_app_Error_message_tag                 1
#define fk_app_Files_files_tag                   1
#define fk_app_LiveData_samples_tag              1
#define fk_app_Capabilities_version_tag          1
#define fk_app_Capabilities_name_tag             2
#define fk_app_Capabilities_modules_tag          3
#define fk_app_Capabilities_sensors_tag          4
#define fk_app_ConfigureSensorQuery_id_tag       1
#define fk_app_ConfigureSensorQuery_frequency_tag 2
#define fk_app_DataSet_id_tag                    1
#define fk_app_DataSet_sensor_tag                2
#define fk_app_DataSet_time_tag                  3
#define fk_app_DataSet_size_tag                  4
#define fk_app_DataSet_pages_tag                 5
#define fk_app_DataSet_hash_tag                  6
#define fk_app_DataSet_name_tag                  7
#define fk_app_DataSetData_time_tag              1
#define fk_app_DataSetData_page_tag              2
#define fk_app_DataSetData_sensor_tag            3
#define fk_app_DataSetData_samples_tag           4
#define fk_app_DataSetData_floats_tag            5
#define fk_app_DataSetData_data_tag              6
#define fk_app_DataSetData_hash_tag              7
#define fk_app_DownloadDataSet_id_tag            1
#define fk_app_DownloadDataSet_page_tag          2
#define fk_app_DownloadFile_id_tag               1
#define fk_app_DownloadFile_page_tag             2
#define fk_app_DownloadFile_token_tag            3
#define fk_app_EraseDataSet_id_tag               1
#define fk_app_File_id_tag                       1
#define fk_app_File_time_tag                     2
#define fk_app_File_size_tag                     3
#define fk_app_File_pages_tag                    4
#define fk_app_File_name_tag                     5
#define fk_app_FileData_page_tag                 1
#define fk_app_FileData_data_tag                 2
#define fk_app_FileData_hash_tag                 3
#define fk_app_FileData_token_tag                4
#define fk_app_LiveDataPoll_interval_tag         1
#define fk_app_LiveDataSample_sensor_tag         1
#define fk_app_LiveDataSample_time_tag           2
#define fk_app_LiveDataSample_value_tag          3
#define fk_app_ModuleCapabilities_id_tag         1
#define fk_app_ModuleCapabilities_name_tag       2
#define fk_app_QueryCapabilities_version_tag     1
#define fk_app_QueryCapabilities_callerTime_tag  2
#define fk_app_QueryDataSet_id_tag               1
#define fk_app_Sample_time_tag                   1
#define fk_app_Sample_value_tag                  2
#define fk_app_SensorCapabilities_id_tag         1
#define fk_app_SensorCapabilities_module_tag     2
#define fk_app_SensorCapabilities_name_tag       3
#define fk_app_SensorCapabilities_frequency_tag  4
#define fk_app_SensorCapabilities_unitOfMeasure_tag 5
#define fk_app_TimeSpec_fixed_tag                1
#define fk_app_TimeSpec_interval_tag             2
#define fk_app_TimeSpec_offset_tag               3
#define fk_app_Schedule_second_tag               1
#define fk_app_Schedule_minute_tag               2
#define fk_app_Schedule_hour_tag                 3
#define fk_app_Schedule_day_tag                  4
#define fk_app_Schedules_readings_tag            1
#define fk_app_Schedules_transmission_tag        2
#define fk_app_Schedules_status_tag              3
#define fk_app_Schedules_location_tag            4
#define fk_app_WireMessageQuery_type_tag         1
#define fk_app_WireMessageQuery_queryCapabilities_tag 2
#define fk_app_WireMessageQuery_configureSensor_tag 3
#define fk_app_WireMessageQuery_queryDataSets_tag 4
#define fk_app_WireMessageQuery_queryDataSet_tag 5
#define fk_app_WireMessageQuery_downloadDataSet_tag 6
#define fk_app_WireMessageQuery_eraseDataSet_tag 7
#define fk_app_WireMessageQuery_liveDataPoll_tag 8
#define fk_app_WireMessageQuery_newSchedules_tag 9
#define fk_app_WireMessageQuery_downloadFile_tag 10
#define fk_app_WireMessageReply_type_tag         1
#define fk_app_WireMessageReply_errors_tag       2
#define fk_app_WireMessageReply_capabilities_tag 3
#define fk_app_WireMessageReply_dataSets_tag     4
#define fk_app_WireMessageReply_dataSetData_tag  5
#define fk_app_WireMessageReply_liveData_tag     6
#define fk_app_WireMessageReply_schedules_tag    7
#define fk_app_WireMessageReply_files_tag        8
#define fk_app_WireMessageReply_fileData_tag     9

/* Struct field encoding specification for nanopb */
extern const pb_field_t fk_app_QueryCapabilities_fields[3];
extern const pb_field_t fk_app_ModuleCapabilities_fields[3];
extern const pb_field_t fk_app_SensorCapabilities_fields[6];
extern const pb_field_t fk_app_Capabilities_fields[5];
extern const pb_field_t fk_app_TimeSpec_fields[4];
extern const pb_field_t fk_app_Schedule_fields[5];
extern const pb_field_t fk_app_Schedules_fields[5];
extern const pb_field_t fk_app_ConfigureSensorQuery_fields[3];
extern const pb_field_t fk_app_QueryDataSets_fields[1];
extern const pb_field_t fk_app_DataSet_fields[8];
extern const pb_field_t fk_app_DataSets_fields[2];
extern const pb_field_t fk_app_DownloadDataSet_fields[3];
extern const pb_field_t fk_app_Sample_fields[3];
extern const pb_field_t fk_app_DataSetData_fields[8];
extern const pb_field_t fk_app_LiveDataPoll_fields[2];
extern const pb_field_t fk_app_LiveDataSample_fields[4];
extern const pb_field_t fk_app_LiveData_fields[2];
extern const pb_field_t fk_app_QueryDataSet_fields[2];
extern const pb_field_t fk_app_EraseDataSet_fields[2];
extern const pb_field_t fk_app_File_fields[6];
extern const pb_field_t fk_app_Files_fields[2];
extern const pb_field_t fk_app_DownloadFile_fields[4];
extern const pb_field_t fk_app_FileData_fields[5];
extern const pb_field_t fk_app_WireMessageQuery_fields[11];
extern const pb_field_t fk_app_Error_fields[2];
extern const pb_field_t fk_app_WireMessageReply_fields[10];

/* Maximum encoded size of messages (where known) */
#define fk_app_QueryCapabilities_size            12
/* fk_app_ModuleCapabilities_size depends on runtime parameters */
/* fk_app_SensorCapabilities_size depends on runtime parameters */
/* fk_app_Capabilities_size depends on runtime parameters */
#define fk_app_TimeSpec_size                     33
#define fk_app_Schedule_size                     140
#define fk_app_Schedules_size                    572
#define fk_app_ConfigureSensorQuery_size         12
#define fk_app_QueryDataSets_size                0
/* fk_app_DataSet_size depends on runtime parameters */
/* fk_app_DataSets_size depends on runtime parameters */
#define fk_app_DownloadDataSet_size              12
#define fk_app_Sample_size                       16
/* fk_app_DataSetData_size depends on runtime parameters */
#define fk_app_LiveDataPoll_size                 6
#define fk_app_LiveDataSample_size               22
/* fk_app_LiveData_size depends on runtime parameters */
#define fk_app_QueryDataSet_size                 6
#define fk_app_EraseDataSet_size                 6
/* fk_app_File_size depends on runtime parameters */
/* fk_app_Files_size depends on runtime parameters */
/* fk_app_DownloadFile_size depends on runtime parameters */
/* fk_app_FileData_size depends on runtime parameters */
#define fk_app_WireMessageQuery_size             (651 + fk_app_DownloadFile_size)
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
