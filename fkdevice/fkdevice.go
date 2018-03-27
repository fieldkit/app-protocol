package fkdevice

import (
	"encoding/json"
	"fmt"
	pb "github.com/fieldkit/app-protocol"
	"github.com/golang/protobuf/proto"
	"github.com/robinpowered/go-proto/message"
	"github.com/robinpowered/go-proto/stream"
	_ "hash/crc32"
	"io"
	"log"
	"net"
	"time"
)

type DeviceClientLoggingCallbacks interface {
	Sent(query *pb.WireMessageQuery)
	Received(reply *pb.WireMessageReply)
}

type LogJsonCallbacks struct {
}

func (cb *LogJsonCallbacks) Sent(query *pb.WireMessageQuery) {
	queryJson, err := json.MarshalIndent(query, "", "  ")
	if err == nil {
		log.Printf("Sending: %s", queryJson)
	}
}

func (cb *LogJsonCallbacks) Received(reply *pb.WireMessageReply) {
	replyJson, err := json.MarshalIndent(reply, "", "  ")
	if err == nil {
		log.Printf("Received: %s", replyJson)
	}
}

type DeviceClient struct {
	Callbacks DeviceClientLoggingCallbacks
	Address   string
	Port      int
}

func (d *DeviceClient) QueryCapabilities() (*pb.WireMessageReply, error) {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_CAPABILITIES,
	}
	reply, err := d.queryDevice(query)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) QueryStatus() (*pb.WireMessageReply, error) {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_STATUS,
	}
	reply, err := d.queryDevice(query)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) QuerySchedules() (*pb.WireMessageReply, error) {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_SCHEDULES,
	}
	reply, err := d.queryDevice(query)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) QueryDataSets() (*pb.WireMessageReply, error) {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_DATA_SETS,
	}
	reply, err := d.queryDevice(query)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) QueryDataSet(id uint32) (*pb.WireMessageReply, error) {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_DATA_SET,
		QueryDataSet: &pb.QueryDataSet{
			Id: id,
		},
	}
	reply, err := d.queryDevice(query)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) EraseDataSet(id uint32) (*pb.WireMessageReply, error) {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_ERASE_DATA_SET,
		EraseDataSet: &pb.EraseDataSet{
			Id: id,
		},
	}
	reply, err := d.queryDevice(query)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) DownloadDataSet(id uint32, pages uint32) (*pb.WireMessageReply, error) {
	for page := 0; uint32(page) < pages; page += 1 {
		query := &pb.WireMessageQuery{
			Type: pb.QueryType_QUERY_DOWNLOAD_DATA_SET,
			DownloadDataSet: &pb.DownloadDataSet{
				Id:   uint32(id),
				Page: uint32(page),
			},
		}
		reply, err := d.queryDevice(query)
		if err != nil {
			return nil, err
		}
		fmt.Printf("Page#%d: %+v", page, *reply)
	}
	return nil, nil
}

func (d *DeviceClient) Reset() (*pb.WireMessageReply, error) {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_RESET,
	}
	err := d.queryDeviceNoReply(query)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (d *DeviceClient) QueryFiles() (*pb.WireMessageReply, error) {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_FILES,
	}
	reply, err := d.queryDevice(query)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) EraseFile(id uint32) (*pb.WireMessageReply, error) {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_ERASE_FILE,
		EraseFile: &pb.EraseFile{
			Id: uint32(id),
		},
	}
	reply, err := d.queryDevice(query)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (d *DeviceClient) QueryLiveData(liveDataInterval int) (*pb.WireMessageReply, error) {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_LIVE_DATA_POLL,
		LiveDataPoll: &pb.LiveDataPoll{
			Interval: uint32(liveDataInterval),
		},
	}
	reply, err := d.queryDevice(query)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) QueryNetworkSettings() (*pb.WireMessageReply, error) {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_NETWORK_SETTINGS,
	}
	reply, err := d.queryDevice(query)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) ConfigureNetworkSettings(ns *pb.NetworkSettings) (*pb.WireMessageReply, error) {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_CONFIGURE_NETWORK_SETTINGS,
		NetworkSettings: &pb.NetworkSettings{
			CreateAccessPoint: ns.CreateAccessPoint,
			Networks:          ns.Networks,
		},
	}
	reply, err := d.queryDevice(query)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) QueryIdentity() (*pb.WireMessageReply, error) {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_IDENTITY,
	}

	reply, err := d.queryDevice(query)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) ConfigureIdentity(device, stream string) (*pb.WireMessageReply, error) {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_CONFIGURE_IDENTITY,
		Identity: &pb.Identity{
			Device: device,
			Stream: stream,
		},
	}
	reply, err := d.queryDevice(query)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) QueryFileInformation(id uint32) (*pb.File, error) {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_FILES,
	}
	reply, err := d.queryDevice(query)
	if err != nil {
		return nil, err
	}
	for _, file := range reply.Files.Files { // Files, files, files!
		if file.Id == id {
			return file, nil
		}
	}
	return nil, nil
}

func (d *DeviceClient) QueryModule(id uint32, message []byte) (*pb.WireMessageReply, error) {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_MODULE,
		Module: &pb.QueryModule{
			Id:      id,
			Address: id,
			Message: message,
		},
	}
	reply, err := d.queryDevice(query)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

var DeviceBusyErr = fmt.Errorf("Busy")

func (d *DeviceClient) DownloadFileToWriter(id, pageSize uint32, token []byte, f io.Writer) error {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_DOWNLOAD_FILE,
		DownloadFile: &pb.DownloadFile{
			Id:       uint32(id),
			Page:     uint32(0),
			PageSize: pageSize,
			Token:    token,
		},
	}

	_, err := d.queryDeviceCallback(query, CallbackFunc(func(reply *pb.WireMessageReply) error {
		if reply.Type != pb.ReplyType_REPLY_BUSY {
			_, err := f.Write(reply.FileData.Data)
			if err != nil {
				return fmt.Errorf("Unable to write to file: %v", err)
			}
		} else {
			return DeviceBusyErr
		}
		return nil
	}))

	return err
}

type CallbackFunc func(*pb.WireMessageReply) error

func (d *DeviceClient) openAndSendQuery(query *pb.WireMessageQuery) (net.Conn, error) {
	c, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", d.Address, d.Port), 5*time.Second)
	if err != nil {
		return nil, err
	}

	data, err := proto.Marshal(query)
	if err != nil {
		return nil, err
	}

	// EncodeRawBytes includes the length prefix.
	buf := proto.NewBuffer(make([]byte, 0))
	buf.EncodeRawBytes(data)
	encoded := buf.Bytes()
	c.SetDeadline(time.Now().Add(5 * time.Second))
	_, err = c.Write(encoded)
	if err != nil {
		return nil, err
	}

	return c, nil
}

type DebugReader struct {
	Target io.Reader
}

func (dr *DebugReader) Read(p []byte) (n int, err error) {
	n, err = dr.Target.Read(p)
	return
}

func (d *DeviceClient) queryDeviceCallback(query *pb.WireMessageQuery, callback CallbackFunc) ([]*pb.WireMessageReply, error) {
	c, err := d.openAndSendQuery(query)
	if err != nil {
		return nil, err
	}

	defer c.Close()

	c.SetDeadline(time.Now().Add(10 * time.Second))
	unmarshalFunc := message.UnmarshalFunc(func(b []byte) (proto.Message, error) {
		var reply pb.WireMessageReply
		err := proto.Unmarshal(b, &reply)
		if err != nil {
			return nil, err
		}

		err = callback(&reply)
		if err != nil {
			return nil, err
		}

		c.SetDeadline(time.Now().Add(10 * time.Second))
		return &reply, err
	})

	dr := &DebugReader{
		Target: c,
	}

	collection, err := stream.ReadLengthPrefixedCollection(dr, unmarshalFunc)
	if err != nil {
		return nil, err
	}

	if len(collection) == 0 {
		return nil, fmt.Errorf("No reply.")
	}

	replies := make([]*pb.WireMessageReply, 0)

	for _, m := range collection {
		replies = append(replies, m.(*pb.WireMessageReply))
	}

	return replies, err
}

func (d *DeviceClient) queryDeviceNoReply(query *pb.WireMessageQuery) (err error) {
	if d.Callbacks != nil {
		d.Callbacks.Sent(query)
	}

	c, err := d.openAndSendQuery(query)
	if err != nil {
		return err
	}

	defer c.Close()

	return err
}

func (d *DeviceClient) queryDevice(query *pb.WireMessageQuery) (reply *pb.WireMessageReply, err error) {
	if d.Callbacks != nil {
		d.Callbacks.Sent(query)
	}
	replies, err := d.queryDeviceMultiple(query)
	if err != nil {
		return nil, err
	}
	if len(replies) > 0 {
		if d.Callbacks != nil {
			d.Callbacks.Received(replies[0])
		}
		return replies[0], nil
	}
	return nil, nil
}

func (d *DeviceClient) queryDeviceMultiple(query *pb.WireMessageQuery) (replies []*pb.WireMessageReply, err error) {
	return d.queryDeviceCallback(query, CallbackFunc(func(reply *pb.WireMessageReply) error {
		return nil
	}))
}
