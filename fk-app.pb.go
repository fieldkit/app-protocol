// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fk-app.proto

/*
Package fk_app is a generated protocol buffer package.

It is generated from these files:
	fk-app.proto

It has these top-level messages:
	QueryCapabilities
	SensorCapabilities
	Capabilities
	ConfigureSensorQuery
	QueryDataSets
	DataSet
	DataSets
	DownloadDataSet
	Sample
	DataSetData
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
	QueryType_QUERY_NONE              QueryType = 0
	QueryType_QUERY_CAPABILITIES      QueryType = 1
	QueryType_QUEYR_CONFIGURE_SENSOR  QueryType = 2
	QueryType_QUERY_DATA_SETS         QueryType = 3
	QueryType_QUERY_DATA_SET          QueryType = 4
	QueryType_QUERY_DOWNLOAD_DATA_SET QueryType = 5
	QueryType_QUERY_ERASE_DATA_SET    QueryType = 6
	QueryType_QUERY_LIVE_DATA_POLL    QueryType = 7
)

var QueryType_name = map[int32]string{
	0: "QUERY_NONE",
	1: "QUERY_CAPABILITIES",
	2: "QUEYR_CONFIGURE_SENSOR",
	3: "QUERY_DATA_SETS",
	4: "QUERY_DATA_SET",
	5: "QUERY_DOWNLOAD_DATA_SET",
	6: "QUERY_ERASE_DATA_SET",
	7: "QUERY_LIVE_DATA_POLL",
}
var QueryType_value = map[string]int32{
	"QUERY_NONE":              0,
	"QUERY_CAPABILITIES":      1,
	"QUEYR_CONFIGURE_SENSOR":  2,
	"QUERY_DATA_SETS":         3,
	"QUERY_DATA_SET":          4,
	"QUERY_DOWNLOAD_DATA_SET": 5,
	"QUERY_ERASE_DATA_SET":    6,
	"QUERY_LIVE_DATA_POLL":    7,
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
}

func (x ReplyType) String() string {
	return proto.EnumName(ReplyType_name, int32(x))
}
func (ReplyType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type QueryCapabilities struct {
	Version uint32 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
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

type SensorCapabilities struct {
	Id        uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Frequency uint32 `protobuf:"varint,3,opt,name=frequency" json:"frequency,omitempty"`
}

func (m *SensorCapabilities) Reset()                    { *m = SensorCapabilities{} }
func (m *SensorCapabilities) String() string            { return proto.CompactTextString(m) }
func (*SensorCapabilities) ProtoMessage()               {}
func (*SensorCapabilities) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SensorCapabilities) GetId() uint32 {
	if m != nil {
		return m.Id
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

type Capabilities struct {
	Version uint32                `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Name    string                `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Sensors []*SensorCapabilities `protobuf:"bytes,4,rep,name=sensors" json:"sensors,omitempty"`
}

func (m *Capabilities) Reset()                    { *m = Capabilities{} }
func (m *Capabilities) String() string            { return proto.CompactTextString(m) }
func (*Capabilities) ProtoMessage()               {}
func (*Capabilities) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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

func (m *Capabilities) GetSensors() []*SensorCapabilities {
	if m != nil {
		return m.Sensors
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
func (*ConfigureSensorQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

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
func (*QueryDataSets) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

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
func (*DataSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

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
func (*DataSets) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

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
func (*DownloadDataSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

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
func (*Sample) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

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
func (*DataSetData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

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

type LiveData struct {
	DataSetDatas []*DataSetData `protobuf:"bytes,1,rep,name=dataSetDatas" json:"dataSetDatas,omitempty"`
}

func (m *LiveData) Reset()                    { *m = LiveData{} }
func (m *LiveData) String() string            { return proto.CompactTextString(m) }
func (*LiveData) ProtoMessage()               {}
func (*LiveData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *LiveData) GetDataSetDatas() []*DataSetData {
	if m != nil {
		return m.DataSetDatas
	}
	return nil
}

type QueryDataSet struct {
	Id uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *QueryDataSet) Reset()                    { *m = QueryDataSet{} }
func (m *QueryDataSet) String() string            { return proto.CompactTextString(m) }
func (*QueryDataSet) ProtoMessage()               {}
func (*QueryDataSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

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
func (*EraseDataSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

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
}

func (m *WireMessageQuery) Reset()                    { *m = WireMessageQuery{} }
func (m *WireMessageQuery) String() string            { return proto.CompactTextString(m) }
func (*WireMessageQuery) ProtoMessage()               {}
func (*WireMessageQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

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

type Error struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

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
}

func (m *WireMessageReply) Reset()                    { *m = WireMessageReply{} }
func (m *WireMessageReply) String() string            { return proto.CompactTextString(m) }
func (*WireMessageReply) ProtoMessage()               {}
func (*WireMessageReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

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

func init() {
	proto.RegisterType((*QueryCapabilities)(nil), "fk_app.QueryCapabilities")
	proto.RegisterType((*SensorCapabilities)(nil), "fk_app.SensorCapabilities")
	proto.RegisterType((*Capabilities)(nil), "fk_app.Capabilities")
	proto.RegisterType((*ConfigureSensorQuery)(nil), "fk_app.ConfigureSensorQuery")
	proto.RegisterType((*QueryDataSets)(nil), "fk_app.QueryDataSets")
	proto.RegisterType((*DataSet)(nil), "fk_app.DataSet")
	proto.RegisterType((*DataSets)(nil), "fk_app.DataSets")
	proto.RegisterType((*DownloadDataSet)(nil), "fk_app.DownloadDataSet")
	proto.RegisterType((*Sample)(nil), "fk_app.Sample")
	proto.RegisterType((*DataSetData)(nil), "fk_app.DataSetData")
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
	// 882 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0x5b, 0x8f, 0xda, 0x46,
	0x14, 0xae, 0x2f, 0xc0, 0x72, 0xcc, 0xc5, 0x4c, 0x28, 0x71, 0xd3, 0xa8, 0xda, 0x5a, 0x8a, 0x84,
	0xd2, 0x66, 0x1f, 0xdc, 0x46, 0xa9, 0xd4, 0x27, 0x0a, 0x4e, 0x84, 0x44, 0x61, 0x77, 0x0c, 0x89,
	0xf2, 0x84, 0x9c, 0x65, 0xd8, 0x58, 0x61, 0xb1, 0xd7, 0x86, 0xad, 0xe8, 0x1f, 0xe9, 0xaf, 0xe8,
	0xf3, 0xbe, 0xf5, 0xbd, 0xff, 0xaa, 0x9a, 0x8b, 0x2f, 0x63, 0x58, 0xa9, 0x4f, 0xcc, 0x9c, 0xf3,
	0xcd, 0xcc, 0xf7, 0x7d, 0xe7, 0x1c, 0x0b, 0x68, 0xac, 0xbf, 0xbc, 0xf2, 0xa3, 0xe8, 0x22, 0x8a,
	0xc3, 0x5d, 0x88, 0xaa, 0xeb, 0x2f, 0x4b, 0x3f, 0x8a, 0xec, 0x57, 0xd0, 0xb9, 0xda, 0x93, 0xf8,
	0x30, 0xf4, 0x23, 0xff, 0x53, 0xb0, 0x09, 0x76, 0x01, 0x49, 0x90, 0x05, 0xb5, 0x7b, 0x12, 0x27,
	0x41, 0xb8, 0xb5, 0x94, 0x73, 0xa5, 0xdf, 0xc4, 0xe9, 0xd6, 0x7e, 0x0f, 0xc8, 0x23, 0xdb, 0x24,
	0x8c, 0x25, 0x7c, 0x0b, 0xd4, 0x60, 0x25, 0xa0, 0x6a, 0xb0, 0x42, 0x08, 0xf4, 0xad, 0x7f, 0x4b,
	0x2c, 0xf5, 0x5c, 0xe9, 0xd7, 0x31, 0x5b, 0xa3, 0xe7, 0x50, 0x5f, 0xc7, 0xe4, 0x6e, 0x4f, 0xb6,
	0xd7, 0x07, 0x4b, 0x63, 0xd0, 0x3c, 0x60, 0xc7, 0xd0, 0xf8, 0x7f, 0x0c, 0xb2, 0xbb, 0xb5, 0xc2,
	0xdd, 0x3f, 0x43, 0x2d, 0x61, 0xac, 0x12, 0x4b, 0x3f, 0xd7, 0xfa, 0x86, 0xf3, 0xec, 0x82, 0xcb,
	0xbb, 0x38, 0x26, 0x8b, 0x53, 0xa8, 0x3d, 0x82, 0xee, 0x30, 0xdc, 0xae, 0x83, 0x9b, 0x7d, 0x4c,
	0x38, 0x8e, 0x39, 0x71, 0xa4, 0x46, 0x62, 0xae, 0x96, 0x99, 0xb7, 0xa1, 0xc9, 0x8e, 0x8d, 0xfc,
	0x9d, 0xef, 0x91, 0x5d, 0x62, 0xff, 0xa5, 0x40, 0x4d, 0x6c, 0x8e, 0xae, 0xea, 0x41, 0x95, 0xbf,
	0x2e, 0xee, 0x11, 0x3b, 0x2a, 0x6a, 0x17, 0x08, 0x51, 0x3a, 0x66, 0x6b, 0x1a, 0x4b, 0x82, 0x3f,
	0x89, 0xa5, 0x33, 0x24, 0x5b, 0xa3, 0x2e, 0x54, 0x22, 0xff, 0x86, 0x24, 0x56, 0x85, 0x05, 0xf9,
	0x86, 0x22, 0x3f, 0xfb, 0xc9, 0x67, 0xab, 0xca, 0x91, 0x74, 0x9d, 0xd9, 0x54, 0xcb, 0x6d, 0xb2,
	0xdf, 0xc0, 0x59, 0xca, 0x12, 0xfd, 0x00, 0x67, 0x2b, 0xb1, 0xb6, 0x14, 0xe6, 0x59, 0x3b, 0xf5,
	0x4c, 0x60, 0x70, 0x06, 0xb0, 0x5f, 0x43, 0x7b, 0x14, 0xfe, 0xb1, 0xdd, 0x84, 0xfe, 0xea, 0x31,
	0x65, 0x08, 0x74, 0x4a, 0x46, 0xe8, 0x62, 0x6b, 0xdb, 0x81, 0xaa, 0xe7, 0xdf, 0x46, 0x1b, 0x92,
	0xe9, 0x53, 0x0a, 0xfa, 0xba, 0x50, 0xb9, 0xf7, 0x37, 0x7b, 0x7e, 0x44, 0xc5, 0x7c, 0x63, 0x3f,
	0x28, 0x60, 0x88, 0x37, 0xe8, 0xcf, 0xc9, 0x93, 0x27, 0xde, 0x2a, 0x38, 0xab, 0x49, 0xce, 0xf6,
	0xa1, 0x96, 0x30, 0x0e, 0x69, 0x6b, 0xb4, 0xb2, 0xd6, 0x60, 0x61, 0x9c, 0xa6, 0xe9, 0x0d, 0xeb,
	0x4d, 0xe8, 0xef, 0xa8, 0xb9, 0x5a, 0x5f, 0xc5, 0x62, 0x47, 0x5f, 0xa3, 0x46, 0x30, 0x77, 0x1b,
	0x98, 0xad, 0x33, 0xc7, 0x6b, 0xb9, 0xe3, 0xf6, 0x10, 0xce, 0x26, 0xc1, 0x3d, 0x61, 0xac, 0xdf,
	0x40, 0x63, 0x95, 0x8b, 0x48, 0x1d, 0x7e, 0x52, 0x72, 0x98, 0xfe, 0x60, 0x09, 0x68, 0x7f, 0x07,
	0x8d, 0x62, 0x37, 0x95, 0x6d, 0xa6, 0x79, 0x37, 0xf6, 0x13, 0xf2, 0x58, 0xfe, 0x1f, 0x0d, 0xcc,
	0x0f, 0x41, 0x4c, 0x7e, 0x27, 0x49, 0xe2, 0xdf, 0x10, 0xde, 0xd0, 0x2f, 0x40, 0xdf, 0x1d, 0x22,
	0xee, 0x61, 0xcb, 0xe9, 0xa4, 0x2c, 0x58, 0x72, 0x7e, 0x88, 0x08, 0x66, 0x69, 0xf4, 0x0e, 0x3a,
	0x77, 0xe5, 0x4f, 0x01, 0xf3, 0xd8, 0x70, 0xbe, 0x91, 0xce, 0x48, 0xe3, 0x74, 0x7c, 0x06, 0xbd,
	0x85, 0xf6, 0xb5, 0x3c, 0x58, 0xac, 0x28, 0x86, 0xf3, 0x3c, 0xbd, 0xe6, 0xd4, 0xdc, 0xe1, 0xf2,
	0x21, 0xf4, 0x2b, 0x34, 0xef, 0x8a, 0xa3, 0xc5, 0x46, 0xc1, 0x70, 0xbe, 0x96, 0xc8, 0xa4, 0x49,
	0x2c, 0x63, 0xd1, 0x2f, 0xd0, 0x28, 0x06, 0xd8, 0xc4, 0x18, 0x4e, 0xf7, 0xd4, 0x59, 0x2c, 0x21,
	0xd1, 0x00, 0xda, 0x2b, 0xb9, 0xdb, 0x59, 0xed, 0x0d, 0xe7, 0x69, 0x56, 0x3f, 0x39, 0x8d, 0xcb,
	0x78, 0xfa, 0x38, 0x29, 0x94, 0x89, 0xf5, 0x49, 0xe1, 0xf1, 0x62, 0x09, 0xb1, 0x84, 0xb4, 0xbf,
	0x87, 0x8a, 0x1b, 0xc7, 0x61, 0x4c, 0xbf, 0x80, 0xb7, 0xbc, 0x88, 0xac, 0x6e, 0x75, 0x9c, 0x6e,
	0xed, 0xbf, 0x55, 0xa9, 0xc6, 0x98, 0x44, 0x9b, 0x47, 0x6b, 0xcc, 0x92, 0x85, 0x1a, 0xbf, 0x80,
	0x2a, 0xa1, 0xd7, 0xd3, 0xc2, 0xd2, 0x96, 0x6c, 0xe6, 0x94, 0xe2, 0x30, 0xc6, 0x22, 0x49, 0xf9,
	0x5f, 0x17, 0xbb, 0x40, 0x93, 0xf9, 0x4b, 0x0d, 0x20, 0x21, 0xd1, 0x8f, 0x85, 0xef, 0x0a, 0x2f,
	0x97, 0x59, 0xea, 0xfa, 0x24, 0xff, 0xb0, 0xa0, 0xd7, 0x60, 0x14, 0xda, 0x5f, 0xd4, 0xe8, 0xe4,
	0x98, 0x14, 0x71, 0xf4, 0x91, 0x8d, 0x18, 0x35, 0x51, 0x9a, 0xec, 0x91, 0x74, 0x04, 0x71, 0x86,
	0x78, 0xf9, 0xaf, 0x02, 0xf5, 0xac, 0xd7, 0x51, 0x0b, 0xe0, 0x6a, 0xe1, 0xe2, 0x8f, 0xcb, 0xe9,
	0x6c, 0xea, 0x9a, 0x5f, 0xa1, 0x1e, 0x20, 0xbe, 0x1f, 0x0e, 0x2e, 0x07, 0xbf, 0x8d, 0x27, 0xe3,
	0xf9, 0xd8, 0xf5, 0x4c, 0x05, 0x3d, 0x83, 0xde, 0xd5, 0xc2, 0xfd, 0x88, 0x97, 0xc3, 0xd9, 0xf4,
	0xed, 0xf8, 0xdd, 0x02, 0xbb, 0x4b, 0xcf, 0x9d, 0x7a, 0x33, 0x6c, 0xaa, 0xe8, 0x09, 0xb4, 0xf9,
	0x99, 0xd1, 0x60, 0x3e, 0x58, 0x7a, 0xee, 0xdc, 0x33, 0x35, 0x84, 0xa0, 0x25, 0x07, 0x4d, 0x1d,
	0x7d, 0x0b, 0x4f, 0x45, 0x6c, 0xf6, 0x61, 0x3a, 0x99, 0x0d, 0x46, 0x79, 0xb2, 0x82, 0x2c, 0xe8,
	0xf2, 0xa4, 0x8b, 0x07, 0x9e, 0x9b, 0x67, 0xaa, 0x79, 0x66, 0x32, 0x7e, 0x2f, 0x12, 0x97, 0xb3,
	0xc9, 0xc4, 0xac, 0xbd, 0x7c, 0x50, 0xa0, 0x9e, 0xd5, 0x94, 0x6a, 0xc1, 0xee, 0xe5, 0x24, 0xd3,
	0xd2, 0x81, 0x26, 0xdf, 0x7b, 0x8b, 0xe1, 0xd0, 0xf5, 0xa8, 0x8c, 0x36, 0x18, 0x3c, 0xe4, 0x62,
	0xcc, 0xb8, 0xf7, 0x00, 0xf1, 0x80, 0xa4, 0x57, 0xa3, 0x9a, 0x78, 0x3c, 0xd7, 0xa4, 0x53, 0x4d,
	0x72, 0xd0, 0xac, 0x50, 0x4d, 0x22, 0x76, 0xa4, 0x89, 0x31, 0xe7, 0xc9, 0x32, 0xf3, 0x4f, 0x55,
	0xf6, 0xbf, 0xe3, 0xa7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x1d, 0x68, 0x1e, 0x87, 0x08,
	0x00, 0x00,
}
