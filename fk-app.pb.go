// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fk-app.proto

/*
Package fk_app is a generated protocol buffer package.

It is generated from these files:
	fk-app.proto

It has these top-level messages:
	QueryCapabilities
	ModuleCapabilities
	SensorCapabilities
	Capabilities
	TimeSpec
	Schedule
	Schedules
	ConfigureSensorQuery
	QueryDataSets
	DataSet
	DataSets
	DownloadDataSet
	Sample
	DataSetData
	LiveDataPoll
	LiveDataSample
	LiveData
	QueryDataSet
	EraseDataSet
	File
	Files
	DownloadFile
	FileData
	WireMessageQuery
	Error
	WireMessageReply
*/
package fk_app

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type QueryType int32

const (
	QueryType_QUERY_NONE               QueryType = 0
	QueryType_QUERY_CAPABILITIES       QueryType = 1
	QueryType_QUERY_CONFIGURE_SENSOR   QueryType = 2
	QueryType_QUERY_DATA_SETS          QueryType = 3
	QueryType_QUERY_DATA_SET           QueryType = 4
	QueryType_QUERY_DOWNLOAD_DATA_SET  QueryType = 5
	QueryType_QUERY_ERASE_DATA_SET     QueryType = 6
	QueryType_QUERY_LIVE_DATA_POLL     QueryType = 7
	QueryType_QUERY_SCHEDULES          QueryType = 8
	QueryType_QUERY_CONFIGUE_SCHEDULES QueryType = 9
	QueryType_QUERY_FILES              QueryType = 10
	QueryType_QUERY_DOWNLOAD_FILE      QueryType = 11
)

var QueryType_name = map[int32]string{
	0:  "QUERY_NONE",
	1:  "QUERY_CAPABILITIES",
	2:  "QUERY_CONFIGURE_SENSOR",
	3:  "QUERY_DATA_SETS",
	4:  "QUERY_DATA_SET",
	5:  "QUERY_DOWNLOAD_DATA_SET",
	6:  "QUERY_ERASE_DATA_SET",
	7:  "QUERY_LIVE_DATA_POLL",
	8:  "QUERY_SCHEDULES",
	9:  "QUERY_CONFIGUE_SCHEDULES",
	10: "QUERY_FILES",
	11: "QUERY_DOWNLOAD_FILE",
}
var QueryType_value = map[string]int32{
	"QUERY_NONE":               0,
	"QUERY_CAPABILITIES":       1,
	"QUERY_CONFIGURE_SENSOR":   2,
	"QUERY_DATA_SETS":          3,
	"QUERY_DATA_SET":           4,
	"QUERY_DOWNLOAD_DATA_SET":  5,
	"QUERY_ERASE_DATA_SET":     6,
	"QUERY_LIVE_DATA_POLL":     7,
	"QUERY_SCHEDULES":          8,
	"QUERY_CONFIGUE_SCHEDULES": 9,
	"QUERY_FILES":              10,
	"QUERY_DOWNLOAD_FILE":      11,
}

func (x QueryType) String() string {
	return proto.EnumName(QueryType_name, int32(x))
}
func (QueryType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ReplyType int32

const (
	ReplyType_REPLY_NONE              ReplyType = 0
	ReplyType_REPLY_SUCCESS           ReplyType = 1
	ReplyType_REPLY_ERROR             ReplyType = 2
	ReplyType_REPLY_CAPABILITIES      ReplyType = 3
	ReplyType_REPLY_DATA_SETS         ReplyType = 4
	ReplyType_REPLY_DATA_SET          ReplyType = 5
	ReplyType_REPLY_DOWNLOAD_DATA_SET ReplyType = 6
	ReplyType_REPLY_LIVE_DATA_POLL    ReplyType = 7
	ReplyType_REPLY_SCHEDULES         ReplyType = 8
	ReplyType_REPLY_FILES             ReplyType = 9
	ReplyType_REPLY_DOWNLOAD_FILE     ReplyType = 10
)

var ReplyType_name = map[int32]string{
	0:  "REPLY_NONE",
	1:  "REPLY_SUCCESS",
	2:  "REPLY_ERROR",
	3:  "REPLY_CAPABILITIES",
	4:  "REPLY_DATA_SETS",
	5:  "REPLY_DATA_SET",
	6:  "REPLY_DOWNLOAD_DATA_SET",
	7:  "REPLY_LIVE_DATA_POLL",
	8:  "REPLY_SCHEDULES",
	9:  "REPLY_FILES",
	10: "REPLY_DOWNLOAD_FILE",
}
var ReplyType_value = map[string]int32{
	"REPLY_NONE":              0,
	"REPLY_SUCCESS":           1,
	"REPLY_ERROR":             2,
	"REPLY_CAPABILITIES":      3,
	"REPLY_DATA_SETS":         4,
	"REPLY_DATA_SET":          5,
	"REPLY_DOWNLOAD_DATA_SET": 6,
	"REPLY_LIVE_DATA_POLL":    7,
	"REPLY_SCHEDULES":         8,
	"REPLY_FILES":             9,
	"REPLY_DOWNLOAD_FILE":     10,
}

func (x ReplyType) String() string {
	return proto.EnumName(ReplyType_name, int32(x))
}
func (ReplyType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type QueryCapabilities struct {
	Version    uint32 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	CallerTime uint32 `protobuf:"varint,2,opt,name=callerTime" json:"callerTime,omitempty"`
}

func (m *QueryCapabilities) Reset()                    { *m = QueryCapabilities{} }
func (m *QueryCapabilities) String() string            { return proto.CompactTextString(m) }
func (*QueryCapabilities) ProtoMessage()               {}
func (*QueryCapabilities) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *QueryCapabilities) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *QueryCapabilities) GetCallerTime() uint32 {
	if m != nil {
		return m.CallerTime
	}
	return 0
}

type ModuleCapabilities struct {
	Id   uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *ModuleCapabilities) Reset()                    { *m = ModuleCapabilities{} }
func (m *ModuleCapabilities) String() string            { return proto.CompactTextString(m) }
func (*ModuleCapabilities) ProtoMessage()               {}
func (*ModuleCapabilities) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ModuleCapabilities) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ModuleCapabilities) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SensorCapabilities struct {
	Id            uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Module        uint32 `protobuf:"varint,2,opt,name=module" json:"module,omitempty"`
	Name          string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Frequency     uint32 `protobuf:"varint,4,opt,name=frequency" json:"frequency,omitempty"`
	UnitOfMeasure string `protobuf:"bytes,5,opt,name=unitOfMeasure" json:"unitOfMeasure,omitempty"`
}

func (m *SensorCapabilities) Reset()                    { *m = SensorCapabilities{} }
func (m *SensorCapabilities) String() string            { return proto.CompactTextString(m) }
func (*SensorCapabilities) ProtoMessage()               {}
func (*SensorCapabilities) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SensorCapabilities) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SensorCapabilities) GetModule() uint32 {
	if m != nil {
		return m.Module
	}
	return 0
}

func (m *SensorCapabilities) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SensorCapabilities) GetFrequency() uint32 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

func (m *SensorCapabilities) GetUnitOfMeasure() string {
	if m != nil {
		return m.UnitOfMeasure
	}
	return ""
}

type Capabilities struct {
	Version uint32                `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Name    string                `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Modules []*ModuleCapabilities `protobuf:"bytes,3,rep,name=modules" json:"modules,omitempty"`
	Sensors []*SensorCapabilities `protobuf:"bytes,4,rep,name=sensors" json:"sensors,omitempty"`
}

func (m *Capabilities) Reset()                    { *m = Capabilities{} }
func (m *Capabilities) String() string            { return proto.CompactTextString(m) }
func (*Capabilities) ProtoMessage()               {}
func (*Capabilities) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Capabilities) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Capabilities) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Capabilities) GetModules() []*ModuleCapabilities {
	if m != nil {
		return m.Modules
	}
	return nil
}

func (m *Capabilities) GetSensors() []*SensorCapabilities {
	if m != nil {
		return m.Sensors
	}
	return nil
}

type TimeSpec struct {
	Fixed    int32 `protobuf:"varint,1,opt,name=fixed" json:"fixed,omitempty"`
	Interval int32 `protobuf:"varint,2,opt,name=interval" json:"interval,omitempty"`
	Offset   int32 `protobuf:"varint,3,opt,name=offset" json:"offset,omitempty"`
}

func (m *TimeSpec) Reset()                    { *m = TimeSpec{} }
func (m *TimeSpec) String() string            { return proto.CompactTextString(m) }
func (*TimeSpec) ProtoMessage()               {}
func (*TimeSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *TimeSpec) GetFixed() int32 {
	if m != nil {
		return m.Fixed
	}
	return 0
}

func (m *TimeSpec) GetInterval() int32 {
	if m != nil {
		return m.Interval
	}
	return 0
}

func (m *TimeSpec) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type Schedule struct {
	Second *TimeSpec `protobuf:"bytes,1,opt,name=second" json:"second,omitempty"`
	Minute *TimeSpec `protobuf:"bytes,2,opt,name=minute" json:"minute,omitempty"`
	Hour   *TimeSpec `protobuf:"bytes,3,opt,name=hour" json:"hour,omitempty"`
	Day    *TimeSpec `protobuf:"bytes,4,opt,name=day" json:"day,omitempty"`
}

func (m *Schedule) Reset()                    { *m = Schedule{} }
func (m *Schedule) String() string            { return proto.CompactTextString(m) }
func (*Schedule) ProtoMessage()               {}
func (*Schedule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Schedule) GetSecond() *TimeSpec {
	if m != nil {
		return m.Second
	}
	return nil
}

func (m *Schedule) GetMinute() *TimeSpec {
	if m != nil {
		return m.Minute
	}
	return nil
}

func (m *Schedule) GetHour() *TimeSpec {
	if m != nil {
		return m.Hour
	}
	return nil
}

func (m *Schedule) GetDay() *TimeSpec {
	if m != nil {
		return m.Day
	}
	return nil
}

type Schedules struct {
	Readings     *Schedule `protobuf:"bytes,1,opt,name=readings" json:"readings,omitempty"`
	Transmission *Schedule `protobuf:"bytes,2,opt,name=transmission" json:"transmission,omitempty"`
	Status       *Schedule `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	Location     *Schedule `protobuf:"bytes,4,opt,name=location" json:"location,omitempty"`
}

func (m *Schedules) Reset()                    { *m = Schedules{} }
func (m *Schedules) String() string            { return proto.CompactTextString(m) }
func (*Schedules) ProtoMessage()               {}
func (*Schedules) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Schedules) GetReadings() *Schedule {
	if m != nil {
		return m.Readings
	}
	return nil
}

func (m *Schedules) GetTransmission() *Schedule {
	if m != nil {
		return m.Transmission
	}
	return nil
}

func (m *Schedules) GetStatus() *Schedule {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Schedules) GetLocation() *Schedule {
	if m != nil {
		return m.Location
	}
	return nil
}

type ConfigureSensorQuery struct {
	Id        uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Frequency uint32 `protobuf:"varint,2,opt,name=frequency" json:"frequency,omitempty"`
}

func (m *ConfigureSensorQuery) Reset()                    { *m = ConfigureSensorQuery{} }
func (m *ConfigureSensorQuery) String() string            { return proto.CompactTextString(m) }
func (*ConfigureSensorQuery) ProtoMessage()               {}
func (*ConfigureSensorQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ConfigureSensorQuery) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ConfigureSensorQuery) GetFrequency() uint32 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

type QueryDataSets struct {
}

func (m *QueryDataSets) Reset()                    { *m = QueryDataSets{} }
func (m *QueryDataSets) String() string            { return proto.CompactTextString(m) }
func (*QueryDataSets) ProtoMessage()               {}
func (*QueryDataSets) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type DataSet struct {
	Id     uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Sensor uint32 `protobuf:"varint,2,opt,name=sensor" json:"sensor,omitempty"`
	Time   uint64 `protobuf:"varint,3,opt,name=time" json:"time,omitempty"`
	Size   uint32 `protobuf:"varint,4,opt,name=size" json:"size,omitempty"`
	Pages  uint32 `protobuf:"varint,5,opt,name=pages" json:"pages,omitempty"`
	Hash   uint32 `protobuf:"varint,6,opt,name=hash" json:"hash,omitempty"`
	Name   string `protobuf:"bytes,7,opt,name=name" json:"name,omitempty"`
}

func (m *DataSet) Reset()                    { *m = DataSet{} }
func (m *DataSet) String() string            { return proto.CompactTextString(m) }
func (*DataSet) ProtoMessage()               {}
func (*DataSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DataSet) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DataSet) GetSensor() uint32 {
	if m != nil {
		return m.Sensor
	}
	return 0
}

func (m *DataSet) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *DataSet) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *DataSet) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *DataSet) GetHash() uint32 {
	if m != nil {
		return m.Hash
	}
	return 0
}

func (m *DataSet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// TODO: Paging?
type DataSets struct {
	DataSets []*DataSet `protobuf:"bytes,1,rep,name=dataSets" json:"dataSets,omitempty"`
}

func (m *DataSets) Reset()                    { *m = DataSets{} }
func (m *DataSets) String() string            { return proto.CompactTextString(m) }
func (*DataSets) ProtoMessage()               {}
func (*DataSets) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *DataSets) GetDataSets() []*DataSet {
	if m != nil {
		return m.DataSets
	}
	return nil
}

type DownloadDataSet struct {
	Id   uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Page uint32 `protobuf:"varint,2,opt,name=page" json:"page,omitempty"`
}

func (m *DownloadDataSet) Reset()                    { *m = DownloadDataSet{} }
func (m *DownloadDataSet) String() string            { return proto.CompactTextString(m) }
func (*DownloadDataSet) ProtoMessage()               {}
func (*DownloadDataSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *DownloadDataSet) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DownloadDataSet) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

type Sample struct {
	Time  uint64  `protobuf:"varint,1,opt,name=time" json:"time,omitempty"`
	Value float32 `protobuf:"fixed32,2,opt,name=value" json:"value,omitempty"`
}

func (m *Sample) Reset()                    { *m = Sample{} }
func (m *Sample) String() string            { return proto.CompactTextString(m) }
func (*Sample) ProtoMessage()               {}
func (*Sample) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *Sample) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Sample) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type DataSetData struct {
	Time    uint64    `protobuf:"varint,1,opt,name=time" json:"time,omitempty"`
	Page    uint32    `protobuf:"varint,2,opt,name=page" json:"page,omitempty"`
	Sensor  uint32    `protobuf:"varint,3,opt,name=sensor" json:"sensor,omitempty"`
	Samples []*Sample `protobuf:"bytes,4,rep,name=samples" json:"samples,omitempty"`
	Floats  []float32 `protobuf:"fixed32,5,rep,packed,name=floats" json:"floats,omitempty"`
	Data    []byte    `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	Hash    uint32    `protobuf:"varint,7,opt,name=hash" json:"hash,omitempty"`
}

func (m *DataSetData) Reset()                    { *m = DataSetData{} }
func (m *DataSetData) String() string            { return proto.CompactTextString(m) }
func (*DataSetData) ProtoMessage()               {}
func (*DataSetData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *DataSetData) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *DataSetData) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *DataSetData) GetSensor() uint32 {
	if m != nil {
		return m.Sensor
	}
	return 0
}

func (m *DataSetData) GetSamples() []*Sample {
	if m != nil {
		return m.Samples
	}
	return nil
}

func (m *DataSetData) GetFloats() []float32 {
	if m != nil {
		return m.Floats
	}
	return nil
}

func (m *DataSetData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *DataSetData) GetHash() uint32 {
	if m != nil {
		return m.Hash
	}
	return 0
}

type LiveDataPoll struct {
	Interval uint32 `protobuf:"varint,1,opt,name=interval" json:"interval,omitempty"`
}

func (m *LiveDataPoll) Reset()                    { *m = LiveDataPoll{} }
func (m *LiveDataPoll) String() string            { return proto.CompactTextString(m) }
func (*LiveDataPoll) ProtoMessage()               {}
func (*LiveDataPoll) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *LiveDataPoll) GetInterval() uint32 {
	if m != nil {
		return m.Interval
	}
	return 0
}

type LiveDataSample struct {
	Sensor uint32  `protobuf:"varint,1,opt,name=sensor" json:"sensor,omitempty"`
	Time   uint64  `protobuf:"varint,2,opt,name=time" json:"time,omitempty"`
	Value  float32 `protobuf:"fixed32,3,opt,name=value" json:"value,omitempty"`
}

func (m *LiveDataSample) Reset()                    { *m = LiveDataSample{} }
func (m *LiveDataSample) String() string            { return proto.CompactTextString(m) }
func (*LiveDataSample) ProtoMessage()               {}
func (*LiveDataSample) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *LiveDataSample) GetSensor() uint32 {
	if m != nil {
		return m.Sensor
	}
	return 0
}

func (m *LiveDataSample) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *LiveDataSample) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type LiveData struct {
	Samples []*LiveDataSample `protobuf:"bytes,1,rep,name=samples" json:"samples,omitempty"`
}

func (m *LiveData) Reset()                    { *m = LiveData{} }
func (m *LiveData) String() string            { return proto.CompactTextString(m) }
func (*LiveData) ProtoMessage()               {}
func (*LiveData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *LiveData) GetSamples() []*LiveDataSample {
	if m != nil {
		return m.Samples
	}
	return nil
}

type QueryDataSet struct {
	Id uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *QueryDataSet) Reset()                    { *m = QueryDataSet{} }
func (m *QueryDataSet) String() string            { return proto.CompactTextString(m) }
func (*QueryDataSet) ProtoMessage()               {}
func (*QueryDataSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *QueryDataSet) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type EraseDataSet struct {
	Id uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *EraseDataSet) Reset()                    { *m = EraseDataSet{} }
func (m *EraseDataSet) String() string            { return proto.CompactTextString(m) }
func (*EraseDataSet) ProtoMessage()               {}
func (*EraseDataSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *EraseDataSet) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type File struct {
	Id    uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Time  uint64 `protobuf:"varint,2,opt,name=time" json:"time,omitempty"`
	Size  uint32 `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
	Pages uint32 `protobuf:"varint,4,opt,name=pages" json:"pages,omitempty"`
	Name  string `protobuf:"bytes,5,opt,name=name" json:"name,omitempty"`
}

func (m *File) Reset()                    { *m = File{} }
func (m *File) String() string            { return proto.CompactTextString(m) }
func (*File) ProtoMessage()               {}
func (*File) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *File) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *File) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *File) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *File) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *File) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Files struct {
	Files []*File `protobuf:"bytes,1,rep,name=files" json:"files,omitempty"`
}

func (m *Files) Reset()                    { *m = Files{} }
func (m *Files) String() string            { return proto.CompactTextString(m) }
func (*Files) ProtoMessage()               {}
func (*Files) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *Files) GetFiles() []*File {
	if m != nil {
		return m.Files
	}
	return nil
}

type DownloadFile struct {
	Id   uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Page uint32 `protobuf:"varint,2,opt,name=page" json:"page,omitempty"`
}

func (m *DownloadFile) Reset()                    { *m = DownloadFile{} }
func (m *DownloadFile) String() string            { return proto.CompactTextString(m) }
func (*DownloadFile) ProtoMessage()               {}
func (*DownloadFile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func (m *DownloadFile) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DownloadFile) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

type FileData struct {
	Page uint32 `protobuf:"varint,1,opt,name=page" json:"page,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Hash uint32 `protobuf:"varint,3,opt,name=hash" json:"hash,omitempty"`
}

func (m *FileData) Reset()                    { *m = FileData{} }
func (m *FileData) String() string            { return proto.CompactTextString(m) }
func (*FileData) ProtoMessage()               {}
func (*FileData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{22} }

func (m *FileData) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *FileData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *FileData) GetHash() uint32 {
	if m != nil {
		return m.Hash
	}
	return 0
}

type WireMessageQuery struct {
	Type              QueryType             `protobuf:"varint,1,opt,name=type,enum=fk_app.QueryType" json:"type,omitempty"`
	QueryCapabilities *QueryCapabilities    `protobuf:"bytes,2,opt,name=queryCapabilities" json:"queryCapabilities,omitempty"`
	ConfigureSensor   *ConfigureSensorQuery `protobuf:"bytes,3,opt,name=configureSensor" json:"configureSensor,omitempty"`
	QueryDataSets     *QueryDataSets        `protobuf:"bytes,4,opt,name=queryDataSets" json:"queryDataSets,omitempty"`
	QueryDataSet      *QueryDataSet         `protobuf:"bytes,5,opt,name=queryDataSet" json:"queryDataSet,omitempty"`
	DownloadDataSet   *DownloadDataSet      `protobuf:"bytes,6,opt,name=downloadDataSet" json:"downloadDataSet,omitempty"`
	EraseDataSet      *EraseDataSet         `protobuf:"bytes,7,opt,name=eraseDataSet" json:"eraseDataSet,omitempty"`
	LiveDataPoll      *LiveDataPoll         `protobuf:"bytes,8,opt,name=liveDataPoll" json:"liveDataPoll,omitempty"`
	NewSchedules      *Schedules            `protobuf:"bytes,9,opt,name=newSchedules" json:"newSchedules,omitempty"`
	DownloadFile      *DownloadFile         `protobuf:"bytes,10,opt,name=downloadFile" json:"downloadFile,omitempty"`
}

func (m *WireMessageQuery) Reset()                    { *m = WireMessageQuery{} }
func (m *WireMessageQuery) String() string            { return proto.CompactTextString(m) }
func (*WireMessageQuery) ProtoMessage()               {}
func (*WireMessageQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{23} }

func (m *WireMessageQuery) GetType() QueryType {
	if m != nil {
		return m.Type
	}
	return QueryType_QUERY_NONE
}

func (m *WireMessageQuery) GetQueryCapabilities() *QueryCapabilities {
	if m != nil {
		return m.QueryCapabilities
	}
	return nil
}

func (m *WireMessageQuery) GetConfigureSensor() *ConfigureSensorQuery {
	if m != nil {
		return m.ConfigureSensor
	}
	return nil
}

func (m *WireMessageQuery) GetQueryDataSets() *QueryDataSets {
	if m != nil {
		return m.QueryDataSets
	}
	return nil
}

func (m *WireMessageQuery) GetQueryDataSet() *QueryDataSet {
	if m != nil {
		return m.QueryDataSet
	}
	return nil
}

func (m *WireMessageQuery) GetDownloadDataSet() *DownloadDataSet {
	if m != nil {
		return m.DownloadDataSet
	}
	return nil
}

func (m *WireMessageQuery) GetEraseDataSet() *EraseDataSet {
	if m != nil {
		return m.EraseDataSet
	}
	return nil
}

func (m *WireMessageQuery) GetLiveDataPoll() *LiveDataPoll {
	if m != nil {
		return m.LiveDataPoll
	}
	return nil
}

func (m *WireMessageQuery) GetNewSchedules() *Schedules {
	if m != nil {
		return m.NewSchedules
	}
	return nil
}

func (m *WireMessageQuery) GetDownloadFile() *DownloadFile {
	if m != nil {
		return m.DownloadFile
	}
	return nil
}

type Error struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{24} }

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type WireMessageReply struct {
	Type         ReplyType     `protobuf:"varint,1,opt,name=type,enum=fk_app.ReplyType" json:"type,omitempty"`
	Errors       []*Error      `protobuf:"bytes,2,rep,name=errors" json:"errors,omitempty"`
	Capabilities *Capabilities `protobuf:"bytes,3,opt,name=capabilities" json:"capabilities,omitempty"`
	DataSets     *DataSets     `protobuf:"bytes,4,opt,name=dataSets" json:"dataSets,omitempty"`
	DataSetData  *DataSetData  `protobuf:"bytes,5,opt,name=dataSetData" json:"dataSetData,omitempty"`
	LiveData     *LiveData     `protobuf:"bytes,6,opt,name=liveData" json:"liveData,omitempty"`
	Schedules    *Schedules    `protobuf:"bytes,7,opt,name=schedules" json:"schedules,omitempty"`
	Files        *Files        `protobuf:"bytes,8,opt,name=files" json:"files,omitempty"`
	FileData     *FileData     `protobuf:"bytes,9,opt,name=fileData" json:"fileData,omitempty"`
}

func (m *WireMessageReply) Reset()                    { *m = WireMessageReply{} }
func (m *WireMessageReply) String() string            { return proto.CompactTextString(m) }
func (*WireMessageReply) ProtoMessage()               {}
func (*WireMessageReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{25} }

func (m *WireMessageReply) GetType() ReplyType {
	if m != nil {
		return m.Type
	}
	return ReplyType_REPLY_NONE
}

func (m *WireMessageReply) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *WireMessageReply) GetCapabilities() *Capabilities {
	if m != nil {
		return m.Capabilities
	}
	return nil
}

func (m *WireMessageReply) GetDataSets() *DataSets {
	if m != nil {
		return m.DataSets
	}
	return nil
}

func (m *WireMessageReply) GetDataSetData() *DataSetData {
	if m != nil {
		return m.DataSetData
	}
	return nil
}

func (m *WireMessageReply) GetLiveData() *LiveData {
	if m != nil {
		return m.LiveData
	}
	return nil
}

func (m *WireMessageReply) GetSchedules() *Schedules {
	if m != nil {
		return m.Schedules
	}
	return nil
}

func (m *WireMessageReply) GetFiles() *Files {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *WireMessageReply) GetFileData() *FileData {
	if m != nil {
		return m.FileData
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryCapabilities)(nil), "fk_app.QueryCapabilities")
	proto.RegisterType((*ModuleCapabilities)(nil), "fk_app.ModuleCapabilities")
	proto.RegisterType((*SensorCapabilities)(nil), "fk_app.SensorCapabilities")
	proto.RegisterType((*Capabilities)(nil), "fk_app.Capabilities")
	proto.RegisterType((*TimeSpec)(nil), "fk_app.TimeSpec")
	proto.RegisterType((*Schedule)(nil), "fk_app.Schedule")
	proto.RegisterType((*Schedules)(nil), "fk_app.Schedules")
	proto.RegisterType((*ConfigureSensorQuery)(nil), "fk_app.ConfigureSensorQuery")
	proto.RegisterType((*QueryDataSets)(nil), "fk_app.QueryDataSets")
	proto.RegisterType((*DataSet)(nil), "fk_app.DataSet")
	proto.RegisterType((*DataSets)(nil), "fk_app.DataSets")
	proto.RegisterType((*DownloadDataSet)(nil), "fk_app.DownloadDataSet")
	proto.RegisterType((*Sample)(nil), "fk_app.Sample")
	proto.RegisterType((*DataSetData)(nil), "fk_app.DataSetData")
	proto.RegisterType((*LiveDataPoll)(nil), "fk_app.LiveDataPoll")
	proto.RegisterType((*LiveDataSample)(nil), "fk_app.LiveDataSample")
	proto.RegisterType((*LiveData)(nil), "fk_app.LiveData")
	proto.RegisterType((*QueryDataSet)(nil), "fk_app.QueryDataSet")
	proto.RegisterType((*EraseDataSet)(nil), "fk_app.EraseDataSet")
	proto.RegisterType((*File)(nil), "fk_app.File")
	proto.RegisterType((*Files)(nil), "fk_app.Files")
	proto.RegisterType((*DownloadFile)(nil), "fk_app.DownloadFile")
	proto.RegisterType((*FileData)(nil), "fk_app.FileData")
	proto.RegisterType((*WireMessageQuery)(nil), "fk_app.WireMessageQuery")
	proto.RegisterType((*Error)(nil), "fk_app.Error")
	proto.RegisterType((*WireMessageReply)(nil), "fk_app.WireMessageReply")
	proto.RegisterEnum("fk_app.QueryType", QueryType_name, QueryType_value)
	proto.RegisterEnum("fk_app.ReplyType", ReplyType_name, ReplyType_value)
}

func init() { proto.RegisterFile("fk-app.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x57, 0xdd, 0x6e, 0xdb, 0xc6,
	0x12, 0x3e, 0x14, 0xa9, 0xbf, 0x91, 0x64, 0xd1, 0x1b, 0x1f, 0x87, 0x27, 0x27, 0x08, 0x52, 0x36,
	0x01, 0x8c, 0xa4, 0x4d, 0x0b, 0x35, 0x41, 0x03, 0xb4, 0x37, 0xaa, 0x4d, 0xa7, 0x06, 0x64, 0xcb,
	0x59, 0xda, 0x0d, 0x7a, 0x65, 0x30, 0xd2, 0xca, 0x26, 0x42, 0x93, 0x0a, 0x97, 0x72, 0xea, 0x3e,
	0x47, 0x81, 0xbe, 0x41, 0xd1, 0xdb, 0x5e, 0xf5, 0x11, 0xfa, 0x2c, 0x7d, 0x81, 0x5e, 0x17, 0xfb,
	0x47, 0x72, 0x29, 0xaa, 0xe8, 0x95, 0xb9, 0x33, 0xdf, 0xec, 0xcc, 0x7e, 0xfb, 0xcd, 0xac, 0x0c,
	0xfd, 0xc5, 0xbb, 0x4f, 0x83, 0xe5, 0xf2, 0xd9, 0x32, 0x4d, 0xb2, 0x04, 0xb5, 0x16, 0xef, 0x2e,
	0x82, 0xe5, 0xd2, 0x3d, 0x86, 0xed, 0xd7, 0x2b, 0x92, 0xde, 0xee, 0x07, 0xcb, 0xe0, 0x6d, 0x18,
	0x85, 0x59, 0x48, 0x28, 0x72, 0xa0, 0x7d, 0x43, 0x52, 0x1a, 0x26, 0xb1, 0x63, 0x3c, 0x34, 0xf6,
	0x06, 0x58, 0x2d, 0xd1, 0x03, 0x80, 0x59, 0x10, 0x45, 0x24, 0x3d, 0x0b, 0xaf, 0x89, 0xd3, 0xe0,
	0xce, 0x92, 0xc5, 0x7d, 0x09, 0xe8, 0x38, 0x99, 0xaf, 0x22, 0xa2, 0xed, 0xb7, 0x05, 0x8d, 0x70,
	0x2e, 0xb7, 0x6a, 0x84, 0x73, 0x84, 0xc0, 0x8a, 0x03, 0x19, 0xdf, 0xc5, 0xfc, 0xdb, 0xfd, 0xc9,
	0x00, 0xe4, 0x93, 0x98, 0x26, 0xe9, 0x3f, 0x86, 0xee, 0x42, 0xeb, 0x9a, 0x27, 0x90, 0xc9, 0xe5,
	0x2a, 0xdf, 0xd2, 0x2c, 0xb6, 0x44, 0xf7, 0xa1, 0xbb, 0x48, 0xc9, 0xfb, 0x15, 0x89, 0x67, 0xb7,
	0x8e, 0xc5, 0xe1, 0x85, 0x01, 0x3d, 0x82, 0xc1, 0x2a, 0x0e, 0xb3, 0xe9, 0xe2, 0x98, 0x04, 0x74,
	0x95, 0x12, 0xa7, 0xc9, 0x43, 0x75, 0xa3, 0xfb, 0xab, 0x01, 0xfd, 0x7f, 0xc9, 0x4d, 0xcd, 0xa9,
	0xd0, 0x73, 0x68, 0x8b, 0x02, 0xa9, 0x63, 0x3e, 0x34, 0xf7, 0x7a, 0xa3, 0x7b, 0xcf, 0x04, 0xf1,
	0xcf, 0xd6, 0x69, 0xc2, 0x0a, 0xca, 0xa2, 0x28, 0xa7, 0x82, 0x3a, 0x96, 0x1e, 0xb5, 0xce, 0x10,
	0x56, 0x50, 0xf7, 0x0c, 0x3a, 0xec, 0x0e, 0xfc, 0x25, 0x99, 0xa1, 0x1d, 0x68, 0x2e, 0xc2, 0x1f,
	0x88, 0x60, 0xae, 0x89, 0xc5, 0x02, 0xdd, 0x83, 0x4e, 0x18, 0x67, 0x24, 0xbd, 0x09, 0x22, 0x5e,
	0x65, 0x13, 0xe7, 0x6b, 0x46, 0x6c, 0xb2, 0x58, 0x50, 0x92, 0x71, 0x0a, 0x9b, 0x58, 0xae, 0x18,
	0x01, 0x1d, 0x7f, 0x76, 0x45, 0x38, 0xcb, 0x7b, 0xd0, 0xa2, 0x64, 0x96, 0xc4, 0x62, 0xdf, 0xde,
	0xc8, 0x56, 0x75, 0xa9, 0xc4, 0x58, 0xfa, 0x19, 0xf2, 0x3a, 0x8c, 0x57, 0x99, 0xa0, 0xa3, 0x16,
	0x29, 0xfc, 0xe8, 0x11, 0x58, 0x57, 0xc9, 0x2a, 0xe5, 0x69, 0xeb, 0x70, 0xdc, 0x8b, 0x5c, 0x30,
	0xe7, 0x81, 0xb8, 0xc5, 0x3a, 0x10, 0x73, 0xba, 0x7f, 0x18, 0xd0, 0x55, 0xa5, 0x52, 0xf4, 0x09,
	0x74, 0x52, 0x12, 0xcc, 0xc3, 0xf8, 0x92, 0x56, 0xab, 0x55, 0x20, 0x9c, 0x23, 0xd0, 0x73, 0xe8,
	0x67, 0x69, 0x10, 0xd3, 0xeb, 0x90, 0xf2, 0xbb, 0x6d, 0x6c, 0x88, 0xd0, 0x50, 0x9c, 0x8f, 0x2c,
	0xc8, 0x56, 0xb4, 0x5a, 0x7d, 0x8e, 0x97, 0x7e, 0x56, 0x4d, 0x94, 0xcc, 0x82, 0x8c, 0xed, 0x6d,
	0x6d, 0xaa, 0x46, 0x21, 0xdc, 0x03, 0xd8, 0xd9, 0x4f, 0xe2, 0x45, 0x78, 0xb9, 0x4a, 0x89, 0xb8,
	0x72, 0xde, 0xa4, 0x6b, 0xdd, 0xa0, 0x29, 0xbc, 0x51, 0x51, 0xb8, 0x3b, 0x84, 0x01, 0x0f, 0x3b,
	0x08, 0xb2, 0xc0, 0x27, 0x19, 0x75, 0x7f, 0x36, 0xa0, 0x2d, 0x17, 0x75, 0x8d, 0x25, 0x84, 0xa4,
	0x1a, 0x4b, 0xac, 0x98, 0xaa, 0xb3, 0x50, 0x36, 0x96, 0x85, 0xf9, 0x37, 0xb3, 0xd1, 0xf0, 0x47,
	0x22, 0x7b, 0x8a, 0x7f, 0x33, 0xc5, 0x2d, 0x83, 0x4b, 0x42, 0x79, 0x1b, 0x0d, 0xb0, 0x58, 0x30,
	0xe4, 0x55, 0x40, 0xaf, 0x9c, 0x96, 0x40, 0xb2, 0xef, 0xbc, 0x4f, 0xda, 0xa5, 0xee, 0xff, 0x12,
	0x3a, 0xaa, 0x4a, 0xf4, 0x14, 0x3a, 0x73, 0xf9, 0xed, 0x18, 0x5c, 0xfe, 0x43, 0x45, 0x95, 0xc4,
	0xe0, 0x1c, 0xe0, 0xbe, 0x80, 0xe1, 0x41, 0xf2, 0x21, 0x8e, 0x92, 0x60, 0xbe, 0xe9, 0x64, 0x08,
	0x2c, 0x56, 0x8c, 0x3c, 0x17, 0xff, 0x76, 0x47, 0xd0, 0xf2, 0x83, 0xeb, 0xa5, 0x18, 0x1c, 0xfc,
	0x7c, 0x46, 0xe9, 0x7c, 0x3b, 0xd0, 0xbc, 0x09, 0xa2, 0x95, 0x08, 0x69, 0x60, 0xb1, 0x70, 0x7f,
	0x37, 0xa0, 0x27, 0x73, 0xb0, 0x3f, 0xb5, 0x91, 0x35, 0xb9, 0x4a, 0xcc, 0x9a, 0x1a, 0xb3, 0x7b,
	0xd0, 0xa6, 0xbc, 0x06, 0xd5, 0xe5, 0x5b, 0xb9, 0x22, 0xb8, 0x19, 0x2b, 0x37, 0xdb, 0x61, 0x11,
	0x25, 0x41, 0xc6, 0xc8, 0x35, 0xf7, 0x1a, 0x58, 0xae, 0x58, 0x36, 0x46, 0x04, 0x67, 0xb7, 0x8f,
	0xf9, 0x77, 0xce, 0x78, 0xbb, 0x60, 0xdc, 0x7d, 0x02, 0xfd, 0x49, 0x78, 0x43, 0x58, 0xd5, 0xa7,
	0x49, 0x14, 0x69, 0x73, 0x40, 0xf0, 0x94, 0xaf, 0x5d, 0x0c, 0x5b, 0x0a, 0x2b, 0x19, 0x2a, 0xea,
	0x37, 0x6a, 0x95, 0xd1, 0xa8, 0x63, 0xce, 0x2c, 0x33, 0xf7, 0x35, 0x74, 0xd4, 0x9e, 0xe8, 0xf3,
	0xe2, 0xd4, 0xe2, 0x72, 0x77, 0xd5, 0xa9, 0xf5, 0xb4, 0xf9, 0xe9, 0xdd, 0x07, 0xd0, 0x2f, 0xcb,
	0xb8, 0x7a, 0xbf, 0xcc, 0xef, 0xa5, 0x01, 0x25, 0x9b, 0xfc, 0x57, 0x60, 0x1d, 0x86, 0x11, 0xa9,
	0xd3, 0xc5, 0x5a, 0xfd, 0x4a, 0xd9, 0x66, 0x9d, 0xb2, 0xad, 0x8a, 0xb2, 0xb9, 0x8a, 0x9b, 0x25,
	0x15, 0x3f, 0x85, 0x26, 0xcb, 0x44, 0x91, 0xcb, 0xc6, 0x6f, 0x71, 0xc4, 0xbe, 0x3a, 0x22, 0xf3,
	0x62, 0xe1, 0x72, 0x47, 0xd0, 0x57, 0xca, 0xdd, 0x54, 0xde, 0x9a, 0x6c, 0x0f, 0xa1, 0xc3, 0xb0,
	0x4a, 0x7e, 0xdc, 0x6f, 0x94, 0xa4, 0xa6, 0x04, 0xd1, 0xa8, 0x11, 0x84, 0x59, 0x12, 0xc4, 0x9f,
	0x16, 0xd8, 0x6f, 0xc2, 0x94, 0x1c, 0x13, 0x4a, 0x83, 0x4b, 0x22, 0x86, 0xcb, 0x63, 0xb0, 0xb2,
	0xdb, 0xa5, 0xd8, 0x70, 0x6b, 0xb4, 0xad, 0x6a, 0xe6, 0xce, 0xb3, 0xdb, 0x25, 0xc1, 0xdc, 0x8d,
	0x5e, 0xc1, 0xf6, 0xfb, 0xea, 0x2f, 0x06, 0x39, 0x2e, 0xff, 0xa7, 0xc5, 0x68, 0xaf, 0xd4, 0x7a,
	0x0c, 0x3a, 0x84, 0xe1, 0x4c, 0x1f, 0x72, 0x72, 0x8a, 0xde, 0x57, 0xdb, 0xd4, 0xcd, 0x40, 0x5c,
	0x0d, 0x42, 0x5f, 0xc1, 0xe0, 0x7d, 0x79, 0xcc, 0xc9, 0xf9, 0xfa, 0x5f, 0xad, 0x18, 0xe5, 0xc4,
	0x3a, 0x16, 0xbd, 0x84, 0x7e, 0xd9, 0xc0, 0xaf, 0xb3, 0x37, 0xda, 0xa9, 0x8b, 0xc5, 0x1a, 0x12,
	0x8d, 0x61, 0x38, 0xd7, 0x27, 0x0f, 0xef, 0xc3, 0xde, 0xe8, 0x6e, 0x3e, 0xad, 0x74, 0x37, 0xae,
	0xe2, 0x59, 0x72, 0x52, 0x52, 0x2e, 0xef, 0xd9, 0x52, 0xf2, 0xb2, 0xaa, 0xb1, 0x86, 0x64, 0x91,
	0x51, 0xa9, 0xa3, 0x9d, 0x8e, 0x1e, 0x59, 0xee, 0x76, 0xac, 0x21, 0xd1, 0x0b, 0xe8, 0xc7, 0xe4,
	0x43, 0xfe, 0x4c, 0x3a, 0x5d, 0x1e, 0xb9, 0x5d, 0x7d, 0x8c, 0x28, 0xd6, 0x60, 0x2c, 0xe1, 0xbc,
	0xa4, 0x56, 0x07, 0xf4, 0x84, 0x65, 0x25, 0x63, 0x0d, 0xe9, 0x7e, 0x04, 0x4d, 0x2f, 0x4d, 0x93,
	0x94, 0xfd, 0x72, 0xba, 0x16, 0x7a, 0xe3, 0x12, 0xeb, 0x62, 0xb5, 0x74, 0x7f, 0x33, 0x35, 0x39,
	0x62, 0xb2, 0x8c, 0x36, 0xca, 0x91, 0x3b, 0x4b, 0x72, 0x7c, 0x0c, 0x2d, 0xc2, 0xb6, 0x67, 0x1a,
	0x64, 0xbd, 0x36, 0x28, 0xd8, 0x4b, 0x93, 0x14, 0x4b, 0x27, 0xab, 0x7f, 0x56, 0x16, 0xac, 0xa9,
	0xd7, 0xaf, 0x69, 0x55, 0x43, 0xb2, 0x97, 0x7b, 0xae, 0x2b, 0xcb, 0xae, 0x3c, 0x47, 0xb4, 0x78,
	0x8f, 0xd0, 0x0b, 0xe8, 0xcd, 0x8b, 0x37, 0x42, 0xca, 0xe9, 0x4e, 0x25, 0x80, 0xfd, 0xc1, 0x65,
	0x1c, 0xff, 0x79, 0x20, 0x6f, 0x49, 0xaa, 0xc8, 0xae, 0xde, 0x25, 0xce, 0x11, 0xe8, 0x33, 0xe8,
	0xd2, 0xfc, 0x02, 0xdb, 0x9b, 0x2e, 0xb0, 0xc0, 0xa0, 0x8f, 0xd5, 0x3c, 0x12, 0x3a, 0x19, 0x94,
	0xe7, 0x11, 0x95, 0x03, 0x89, 0xd5, 0xb0, 0x90, 0xc3, 0x45, 0xaa, 0xc2, 0x2e, 0xe3, 0x44, 0x0d,
	0x0a, 0xf1, 0xe4, 0x97, 0x06, 0x74, 0xf3, 0xd1, 0x80, 0xb6, 0x00, 0x5e, 0x9f, 0x7b, 0xf8, 0xfb,
	0x8b, 0x93, 0xe9, 0x89, 0x67, 0xff, 0x07, 0xed, 0x02, 0x12, 0xeb, 0xfd, 0xf1, 0xe9, 0xf8, 0x9b,
	0xa3, 0xc9, 0xd1, 0xd9, 0x91, 0xe7, 0xdb, 0x06, 0xba, 0x07, 0xbb, 0xd2, 0x3e, 0x3d, 0x39, 0x3c,
	0x7a, 0x75, 0x8e, 0xbd, 0x0b, 0xdf, 0x3b, 0xf1, 0xa7, 0xd8, 0x6e, 0xa0, 0x3b, 0x30, 0x14, 0xbe,
	0x83, 0xf1, 0xd9, 0xf8, 0xc2, 0xf7, 0xce, 0x7c, 0xdb, 0x44, 0x08, 0xb6, 0x74, 0xa3, 0x6d, 0xa1,
	0xff, 0xc3, 0x5d, 0x69, 0x9b, 0xbe, 0x39, 0x99, 0x4c, 0xc7, 0x07, 0x85, 0xb3, 0x89, 0x1c, 0xd8,
	0x11, 0x4e, 0x0f, 0x8f, 0x7d, 0xaf, 0xf0, 0xb4, 0x0a, 0xcf, 0xe4, 0xe8, 0x3b, 0xe9, 0x38, 0x9d,
	0x4e, 0x26, 0x76, 0xbb, 0xc8, 0xec, 0xef, 0x7f, 0xeb, 0x1d, 0x9c, 0x4f, 0x3c, 0xdf, 0xee, 0xa0,
	0xfb, 0xe0, 0x68, 0xa5, 0x7a, 0x25, 0x6f, 0x17, 0x0d, 0xa1, 0x27, 0xbc, 0x87, 0x47, 0xcc, 0x00,
	0xe8, 0x2e, 0xdc, 0xa9, 0x14, 0xc5, 0x3c, 0x76, 0xef, 0xc9, 0x5f, 0x06, 0x74, 0x73, 0xd1, 0x32,
	0xa2, 0xb0, 0x77, 0x3a, 0xc9, 0x89, 0xda, 0x86, 0x81, 0x58, 0xfb, 0xe7, 0xfb, 0xfb, 0x9e, 0xcf,
	0x38, 0x1a, 0x42, 0x4f, 0x98, 0x3c, 0x8c, 0x39, 0x31, 0xbb, 0x80, 0x84, 0x41, 0x23, 0xd3, 0x64,
	0x65, 0x0b, 0x7b, 0x41, 0x98, 0xc5, 0x08, 0xd3, 0x8d, 0x76, 0x93, 0x11, 0x26, 0x6d, 0x6b, 0x84,
	0x71, 0x5a, 0x84, 0xb3, 0x8e, 0x16, 0x59, 0x5b, 0x89, 0x96, 0xbc, 0x3a, 0x71, 0xf0, 0x2e, 0x3b,
	0x78, 0x65, 0x73, 0x7e, 0x70, 0x78, 0xdb, 0xe2, 0xff, 0x69, 0x7e, 0xf1, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xa0, 0x3a, 0x26, 0x7a, 0x79, 0x0e, 0x00, 0x00,
}
