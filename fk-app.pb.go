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
)

var QueryType_name = map[int32]string{
	0: "QUERY_NONE",
	1: "QUERY_CAPABILITIES",
	2: "QUERY_CONFIGURE_SENSOR",
	3: "QUERY_DATA_SETS",
	4: "QUERY_DATA_SET",
	5: "QUERY_DOWNLOAD_DATA_SET",
	6: "QUERY_ERASE_DATA_SET",
	7: "QUERY_LIVE_DATA_POLL",
	8: "QUERY_SCHEDULES",
	9: "QUERY_CONFIGUE_SCHEDULES",
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
)

var ReplyType_name = map[int32]string{
	0: "REPLY_NONE",
	1: "REPLY_SUCCESS",
	2: "REPLY_ERROR",
	3: "REPLY_CAPABILITIES",
	4: "REPLY_DATA_SETS",
	5: "REPLY_DATA_SET",
	6: "REPLY_DOWNLOAD_DATA_SET",
	7: "REPLY_LIVE_DATA_POLL",
	8: "REPLY_SCHEDULES",
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
}

func (m *WireMessageQuery) Reset()                    { *m = WireMessageQuery{} }
func (m *WireMessageQuery) String() string            { return proto.CompactTextString(m) }
func (*WireMessageQuery) ProtoMessage()               {}
func (*WireMessageQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

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

type Error struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

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
}

func (m *WireMessageReply) Reset()                    { *m = WireMessageReply{} }
func (m *WireMessageReply) String() string            { return proto.CompactTextString(m) }
func (*WireMessageReply) ProtoMessage()               {}
func (*WireMessageReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

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
	proto.RegisterType((*WireMessageQuery)(nil), "fk_app.WireMessageQuery")
	proto.RegisterType((*Error)(nil), "fk_app.Error")
	proto.RegisterType((*WireMessageReply)(nil), "fk_app.WireMessageReply")
	proto.RegisterEnum("fk_app.QueryType", QueryType_name, QueryType_value)
	proto.RegisterEnum("fk_app.ReplyType", ReplyType_name, ReplyType_value)
}

func init() { proto.RegisterFile("fk-app.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x57, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0x66, 0xfd, 0xef, 0xe3, 0xbf, 0xcd, 0x34, 0xa4, 0x4b, 0xa9, 0xaa, 0xb2, 0x6a, 0xa5, 0xa8,
	0x40, 0x41, 0xa6, 0x15, 0x95, 0xe0, 0xc6, 0xd8, 0xdb, 0x12, 0xc9, 0x89, 0xd3, 0x59, 0x87, 0x8a,
	0xab, 0x68, 0x6a, 0x8f, 0x93, 0x55, 0xd7, 0xbb, 0x9b, 0x9d, 0x75, 0x4a, 0x78, 0x0e, 0x24, 0x5e,
	0x81, 0x3b, 0x2e, 0x79, 0x04, 0x6e, 0x78, 0x06, 0x9e, 0x84, 0x0b, 0x34, 0x3f, 0xfb, 0x33, 0xce,
	0x1a, 0x71, 0xe5, 0x39, 0xe7, 0x7c, 0x67, 0xe6, 0x9c, 0x6f, 0xbe, 0x33, 0x2b, 0x43, 0x77, 0xf5,
	0xee, 0x73, 0x12, 0x45, 0x4f, 0xa3, 0x38, 0x4c, 0x42, 0xd4, 0x58, 0xbd, 0x3b, 0x27, 0x51, 0x64,
	0x1f, 0xc3, 0xde, 0xeb, 0x0d, 0x8d, 0x6f, 0xc6, 0x24, 0x22, 0x6f, 0x3d, 0xdf, 0x4b, 0x3c, 0xca,
	0x90, 0x05, 0xcd, 0x6b, 0x1a, 0x33, 0x2f, 0x0c, 0x2c, 0xe3, 0xa1, 0x71, 0xd8, 0xc3, 0xa9, 0x89,
	0x1e, 0x00, 0x2c, 0x88, 0xef, 0xd3, 0x78, 0xee, 0xad, 0xa9, 0x55, 0x11, 0xc1, 0x82, 0xc7, 0x7e,
	0x01, 0xe8, 0x38, 0x5c, 0x6e, 0x7c, 0xaa, 0xed, 0xd7, 0x87, 0x8a, 0xb7, 0x54, 0x5b, 0x55, 0xbc,
	0x25, 0x42, 0x50, 0x0b, 0x88, 0xca, 0x6f, 0x63, 0xb1, 0xb6, 0x7f, 0x31, 0x00, 0xb9, 0x34, 0x60,
	0x61, 0xfc, 0x9f, 0xa9, 0x07, 0xd0, 0x58, 0x8b, 0x03, 0xd4, 0xe1, 0xca, 0xca, 0xb6, 0xac, 0xe6,
	0x5b, 0xa2, 0xfb, 0xd0, 0x5e, 0xc5, 0xf4, 0x6a, 0x43, 0x83, 0xc5, 0x8d, 0x55, 0x13, 0xf0, 0xdc,
	0x81, 0x1e, 0x41, 0x6f, 0x13, 0x78, 0xc9, 0x6c, 0x75, 0x4c, 0x09, 0xdb, 0xc4, 0xd4, 0xaa, 0x8b,
	0x54, 0xdd, 0x69, 0xff, 0x66, 0x40, 0xf7, 0x7f, 0x72, 0x53, 0xd2, 0x15, 0x7a, 0x06, 0x4d, 0x59,
	0x20, 0xb3, 0xaa, 0x0f, 0xab, 0x87, 0x9d, 0xe1, 0xbd, 0xa7, 0x92, 0xf8, 0xa7, 0xb7, 0x69, 0xc2,
	0x29, 0x94, 0x67, 0x31, 0x41, 0x05, 0xb3, 0x6a, 0x7a, 0xd6, 0x6d, 0x86, 0x70, 0x0a, 0xb5, 0xe7,
	0xd0, 0xe2, 0x77, 0xe0, 0x46, 0x74, 0x81, 0xf6, 0xa1, 0xbe, 0xf2, 0x7e, 0xa2, 0x92, 0xb9, 0x3a,
	0x96, 0x06, 0xba, 0x07, 0x2d, 0x2f, 0x48, 0x68, 0x7c, 0x4d, 0x7c, 0x51, 0x65, 0x1d, 0x67, 0x36,
	0x27, 0x36, 0x5c, 0xad, 0x18, 0x4d, 0x04, 0x85, 0x75, 0xac, 0x2c, 0x4e, 0x40, 0xcb, 0x5d, 0x5c,
	0x52, 0xc1, 0xf2, 0x21, 0x34, 0x18, 0x5d, 0x84, 0x81, 0xdc, 0xb7, 0x33, 0x34, 0xd3, 0xba, 0xd2,
	0x83, 0xb1, 0x8a, 0x73, 0xe4, 0xda, 0x0b, 0x36, 0x89, 0xa4, 0xa3, 0x14, 0x29, 0xe3, 0xe8, 0x11,
	0xd4, 0x2e, 0xc3, 0x4d, 0x2c, 0x8e, 0x2d, 0xc3, 0x89, 0x28, 0xb2, 0xa1, 0xba, 0x24, 0xf2, 0x16,
	0xcb, 0x40, 0x3c, 0x68, 0xff, 0x69, 0x40, 0x3b, 0x2d, 0x95, 0xa1, 0xcf, 0xa0, 0x15, 0x53, 0xb2,
	0xf4, 0x82, 0x0b, 0xb6, 0x5d, 0x6d, 0x0a, 0xc2, 0x19, 0x02, 0x3d, 0x83, 0x6e, 0x12, 0x93, 0x80,
	0xad, 0x3d, 0x26, 0xee, 0xb6, 0xb2, 0x23, 0x43, 0x43, 0x09, 0x3e, 0x12, 0x92, 0x6c, 0xd8, 0x76,
	0xf5, 0x19, 0x5e, 0xc5, 0x79, 0x35, 0x7e, 0xb8, 0x20, 0x09, 0xdf, 0xbb, 0xb6, 0xab, 0x9a, 0x14,
	0x61, 0x4f, 0x60, 0x7f, 0x1c, 0x06, 0x2b, 0xef, 0x62, 0x13, 0x53, 0x79, 0xe5, 0x62, 0x48, 0x6f,
	0x4d, 0x83, 0xa6, 0xf0, 0xca, 0x96, 0xc2, 0xed, 0x01, 0xf4, 0x44, 0xda, 0x84, 0x24, 0xc4, 0xa5,
	0x09, 0xb3, 0x7f, 0x35, 0xa0, 0xa9, 0x8c, 0xb2, 0xc1, 0x92, 0x42, 0x4a, 0x07, 0x4b, 0x5a, 0x5c,
	0xd5, 0x89, 0xa7, 0x06, 0xab, 0x86, 0xc5, 0x9a, 0xfb, 0x98, 0xf7, 0x33, 0x55, 0x33, 0x25, 0xd6,
	0x5c, 0x71, 0x11, 0xb9, 0xa0, 0x4c, 0x8c, 0x51, 0x0f, 0x4b, 0x83, 0x23, 0x2f, 0x09, 0xbb, 0xb4,
	0x1a, 0x12, 0xc9, 0xd7, 0xd9, 0x9c, 0x34, 0x0b, 0xd3, 0xff, 0x35, 0xb4, 0xd2, 0x2a, 0xd1, 0xa7,
	0xd0, 0x5a, 0xaa, 0xb5, 0x65, 0x08, 0xf9, 0x0f, 0x52, 0xaa, 0x14, 0x06, 0x67, 0x00, 0xfb, 0x39,
	0x0c, 0x26, 0xe1, 0xfb, 0xc0, 0x0f, 0xc9, 0x72, 0x57, 0x67, 0x08, 0x6a, 0xbc, 0x18, 0xd5, 0x97,
	0x58, 0xdb, 0x43, 0x68, 0xb8, 0x64, 0x1d, 0xc9, 0x87, 0x43, 0xf4, 0x67, 0x14, 0xfa, 0xdb, 0x87,
	0xfa, 0x35, 0xf1, 0x37, 0x32, 0xa5, 0x82, 0xa5, 0x61, 0xff, 0x61, 0x40, 0x47, 0x9d, 0xc1, 0x7f,
	0x4a, 0x33, 0x4b, 0xce, 0x2a, 0x30, 0x5b, 0xd5, 0x98, 0x3d, 0x84, 0x26, 0x13, 0x35, 0xa4, 0x53,
	0xde, 0xcf, 0x14, 0x21, 0xdc, 0x38, 0x0d, 0xf3, 0x1d, 0x56, 0x7e, 0x48, 0x12, 0x4e, 0x6e, 0xf5,
	0xb0, 0x82, 0x95, 0xc5, 0x4f, 0xe3, 0x44, 0x08, 0x76, 0xbb, 0x58, 0xac, 0x33, 0xc6, 0x9b, 0x39,
	0xe3, 0xf6, 0x13, 0xe8, 0x4e, 0xbd, 0x6b, 0xca, 0xab, 0x3e, 0x0d, 0x7d, 0x5f, 0x7b, 0x07, 0x24,
	0x4f, 0x99, 0x6d, 0x63, 0xe8, 0xa7, 0x58, 0xc5, 0x50, 0x5e, 0xbf, 0x51, 0xaa, 0x8c, 0x4a, 0x19,
	0x73, 0xd5, 0x22, 0x73, 0xdf, 0x42, 0x2b, 0xdd, 0x13, 0x7d, 0x99, 0x77, 0x2d, 0x2f, 0xf7, 0x20,
	0xed, 0x5a, 0x3f, 0x36, 0xeb, 0xde, 0x7e, 0x00, 0xdd, 0xa2, 0x8c, 0xb7, 0xef, 0x97, 0xc7, 0x9d,
	0x98, 0x30, 0xba, 0x2b, 0xfe, 0x7b, 0x0d, 0xcc, 0x37, 0x5e, 0x4c, 0x8f, 0x29, 0x63, 0xe4, 0x82,
	0xca, 0x49, 0x7a, 0x0c, 0xb5, 0xe4, 0x26, 0x92, 0x97, 0xd7, 0x1f, 0xee, 0xa5, 0x35, 0x88, 0xe0,
	0xfc, 0x26, 0xa2, 0x58, 0x84, 0xd1, 0x2b, 0xd8, 0xbb, 0xda, 0xfe, 0x3c, 0xaa, 0xb7, 0xe1, 0x23,
	0x2d, 0x47, 0x7b, 0x92, 0x6f, 0xe7, 0xa0, 0x97, 0x30, 0x58, 0xe8, 0x13, 0xad, 0x9e, 0x8c, 0xfb,
	0xe9, 0x36, 0x65, 0x03, 0x8f, 0xb7, 0x93, 0xd0, 0x37, 0xd0, 0xbb, 0x2a, 0xce, 0xb4, 0x7a, 0x4c,
	0x3e, 0xd4, 0x8a, 0x49, 0x83, 0x58, 0xc7, 0xa2, 0x17, 0xd0, 0x2d, 0x3a, 0xc4, 0xa8, 0x76, 0x86,
	0xfb, 0x65, 0xb9, 0x58, 0x43, 0xa2, 0x11, 0x0c, 0x96, 0xfa, 0x98, 0x09, 0xd1, 0x75, 0x86, 0x77,
	0xb3, 0xd1, 0xd4, 0xc3, 0x78, 0x1b, 0xcf, 0x0f, 0xa7, 0x85, 0x6b, 0x12, 0x02, 0x2d, 0x1c, 0x5e,
	0xbc, 0x42, 0xac, 0x21, 0x79, 0xa6, 0x5f, 0x90, 0xaf, 0xd5, 0xd2, 0x33, 0x8b, 0xd2, 0xc6, 0x1a,
	0x12, 0x3d, 0x87, 0x6e, 0x40, 0xdf, 0x67, 0xdf, 0x04, 0xab, 0x2d, 0x32, 0xf7, 0xb6, 0x5f, 0x5e,
	0x86, 0x35, 0x98, 0xfd, 0x09, 0xd4, 0x9d, 0x38, 0x0e, 0x63, 0xfe, 0xb1, 0x5f, 0x4b, 0xd5, 0x08,
	0xa1, 0xb4, 0x71, 0x6a, 0xda, 0x7f, 0x57, 0x34, 0x51, 0x61, 0x1a, 0xf9, 0x3b, 0x45, 0x25, 0x82,
	0x05, 0x51, 0x3d, 0x86, 0x06, 0xe5, 0xdb, 0x73, 0x25, 0xf1, 0x09, 0xe8, 0xe5, 0x1c, 0xc4, 0x61,
	0x8c, 0x55, 0x90, 0xb7, 0xbd, 0x28, 0xca, 0xae, 0xaa, 0xb7, 0xad, 0x29, 0x4e, 0x43, 0xf2, 0x8f,
	0xcd, 0x52, 0xd7, 0x87, 0xb9, 0xf5, 0x82, 0xb2, 0xfc, 0x09, 0x45, 0xcf, 0xa1, 0xb3, 0xcc, 0x9f,
	0x35, 0x25, 0x8a, 0x3b, 0x5b, 0x09, 0xfc, 0x07, 0x17, 0x71, 0xe2, 0x8b, 0xa6, 0xb8, 0x56, 0x5a,
	0x30, 0xb7, 0x6f, 0x04, 0x67, 0x08, 0xf4, 0x05, 0xb4, 0x59, 0x76, 0x0d, 0xcd, 0x5d, 0xd7, 0x90,
	0x63, 0x9e, 0xfc, 0x63, 0x40, 0x3b, 0x9b, 0x46, 0xd4, 0x07, 0x78, 0x7d, 0xe6, 0xe0, 0x1f, 0xcf,
	0x4f, 0x66, 0x27, 0x8e, 0xf9, 0x01, 0x3a, 0x00, 0x24, 0xed, 0xf1, 0xe8, 0x74, 0xf4, 0xdd, 0xd1,
	0xf4, 0x68, 0x7e, 0xe4, 0xb8, 0xa6, 0x81, 0xee, 0xc1, 0x81, 0xf2, 0xcf, 0x4e, 0x5e, 0x1e, 0xbd,
	0x3a, 0xc3, 0xce, 0xb9, 0xeb, 0x9c, 0xb8, 0x33, 0x6c, 0x56, 0xd0, 0x1d, 0x18, 0xc8, 0xd8, 0x64,
	0x34, 0x1f, 0x9d, 0xbb, 0xce, 0xdc, 0x35, 0xab, 0x08, 0x41, 0x5f, 0x77, 0x9a, 0x35, 0xf4, 0x31,
	0xdc, 0x55, 0xbe, 0xd9, 0x9b, 0x93, 0xe9, 0x6c, 0x34, 0xc9, 0x83, 0x75, 0x64, 0xc1, 0xbe, 0x0c,
	0x3a, 0x78, 0xe4, 0x3a, 0x79, 0xa4, 0x91, 0x47, 0xa6, 0x47, 0x3f, 0xa8, 0xc0, 0xe9, 0x6c, 0x3a,
	0x35, 0x9b, 0xf9, 0xc9, 0xee, 0xf8, 0x7b, 0x67, 0x72, 0x36, 0x75, 0x5c, 0xb3, 0x85, 0xee, 0x83,
	0xa5, 0x95, 0xea, 0x14, 0xa2, 0xed, 0x27, 0x7f, 0x19, 0xd0, 0xce, 0x74, 0xc3, 0xdb, 0xc7, 0xce,
	0xe9, 0x34, 0x6b, 0x7f, 0x0f, 0x7a, 0xd2, 0x76, 0xcf, 0xc6, 0x63, 0xc7, 0xe5, 0x9d, 0x0f, 0xa0,
	0x23, 0x5d, 0x0e, 0xc6, 0xa2, 0xdd, 0x03, 0x40, 0xd2, 0xa1, 0x51, 0x54, 0xe5, 0xc5, 0x48, 0x7f,
	0x4e, 0x43, 0x8d, 0xd3, 0xa0, 0x3b, 0xcd, 0x3a, 0xa7, 0x41, 0xf9, 0x6e, 0xd1, 0x20, 0x9a, 0x95,
	0xc1, 0xb2, 0x66, 0x55, 0x6d, 0x79, 0xb3, 0x6f, 0x1b, 0xe2, 0x5f, 0xc7, 0x57, 0xff, 0x06, 0x00,
	0x00, 0xff, 0xff, 0xfa, 0x0b, 0x25, 0xd8, 0x85, 0x0c, 0x00, 0x00,
}
