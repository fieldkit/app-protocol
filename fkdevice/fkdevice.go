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

type LimitedReader struct {
	Target           io.Reader
	MessagesExpected int
}

func (dr *LimitedReader) Read(p []byte) (n int, err error) {
	if dr.MessagesExpected == 0 {
		return 0, io.EOF
	}
	return dr.Target.Read(p)
}

func (d *DeviceClient) DownloadFileToWriter(id uint32, f io.Writer) error {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_DOWNLOAD_FILE,
		DownloadFile: &pb.DownloadFile{
			Id: uint32(id),
		},
	}

	headerCallback := CallbackFunc(func(reply *pb.WireMessageReply) error {
		if reply.Type != pb.ReplyType_REPLY_BUSY {
			if reply.FileData != nil {
				_, err := f.Write(reply.FileData.Data)
				if err != nil {
					return fmt.Errorf("Unable to write to file: %v", err)
				}
			}
		} else {
			return DeviceBusyErr
		}
		return nil
	})

	bodyCallback := BodyCallbackFunc(func(r io.Reader) error {
		_, err := io.Copy(f, r)
		return err
	})

	_, err := d.queryDeviceDownload(query, headerCallback, bodyCallback)

	return err
}

type CallbackFunc func(*pb.WireMessageReply) error
type BodyCallbackFunc func(r io.Reader) error

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

func (d *DeviceClient) queryDeviceDownload(query *pb.WireMessageQuery, callback CallbackFunc, body BodyCallbackFunc) ([]*pb.WireMessageReply, error) {
	c, err := d.openAndSendQuery(query)
	if err != nil {
		return nil, err
	}

	defer c.Close()

	dr := &DebugReader{
		Target: c,
	}

	lr := &LimitedReader{
		Target:           dr,
		MessagesExpected: 1,
	}

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

		lr.MessagesExpected -= 1

		c.SetDeadline(time.Now().Add(10 * time.Second))
		return &reply, err
	})

	collection, err := stream.ReadLengthPrefixedCollection(lr, unmarshalFunc)
	if err != nil {
		return nil, err
	}

	if len(collection) == 0 {
		return nil, fmt.Errorf("No reply.")
	}

	err = body(dr)
	if err != nil {
		return nil, fmt.Errorf("Error reading body: %v", err)
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
