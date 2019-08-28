package fkdevice

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/robinpowered/go-proto/collection"
	"github.com/robinpowered/go-proto/message"
	"github.com/robinpowered/go-proto/stream"

	pb "github.com/fieldkit/app-protocol"
)

const (
	DefaultTimeout = 5
)

type DeviceClientLoggingCallbacks interface {
	Sent(query *pb.HttpQuery)
	Received(reply *pb.HttpReply)
}

type LogJsonCallbacks struct {
}

func (cb *LogJsonCallbacks) Sent(query *pb.HttpQuery) {
	queryJson, err := json.MarshalIndent(query, "", "  ")
	if err == nil {
		log.Printf("Sending: %s", queryJson)
	}
}

func (cb *LogJsonCallbacks) Received(reply *pb.HttpReply) {
	replyJson, err := json.MarshalIndent(reply, "", "  ")
	if err == nil {
		log.Printf("Received: %s", replyJson)
	}
}

type DeviceClient struct {
	Callbacks DeviceClientLoggingCallbacks
	Address   string
	Port      int
	HexEncode bool
}

func (d *DeviceClient) QueryStatus() (*pb.HttpReply, error) {
	reply, err := d.queryDevice(pb.QueryType_QUERY_STATUS)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) ConfigureName(name string) (*pb.HttpReply, error) {
	query := &pb.HttpQuery{
		Type: pb.QueryType_QUERY_CONFIGURE,
		Identity: &pb.Identity{
			Name: name,
		},
	}
	reply, err := d.queryDeviceQuery(query)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) QueryStartRecording() (*pb.HttpReply, error) {
	query := &pb.HttpQuery{
		Type: pb.QueryType_QUERY_RECORDING_CONTROL,
		Recording: &pb.Recording{
			Enabled: 1,
		},
	}
	reply, err := d.queryDeviceQuery(query)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) QueryStopRecording() (*pb.HttpReply, error) {
	query := &pb.HttpQuery{
		Type: pb.QueryType_QUERY_RECORDING_CONTROL,
		Recording: &pb.Recording{
			Enabled: 0,
		},
	}
	reply, err := d.queryDeviceQuery(query)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) QueryGetReadings() (*pb.HttpReply, error) {
	reply, err := d.queryDevice(pb.QueryType_QUERY_GET_READINGS)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) QueryTakeReadings() (*pb.HttpReply, error) {
	reply, err := d.queryDevice(pb.QueryType_QUERY_TAKE_READINGS)
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

type CallbackFunc func(*pb.WireMessageReply) error
type BodyCallbackFunc func(r io.Reader) error

type QueryResponse struct {
	tcp  net.Conn
	http *http.Response
}

func (qr *QueryResponse) Reader(to int) io.Reader {
	if qr.tcp != nil {
		return &DebugReader{Target: &DeadlineReader{qr.tcp, time.Duration(to) * time.Second}}
	}
	return qr.http.Body
}

func (qr *QueryResponse) Close() error {
	if qr.tcp != nil {
		qr.tcp.Close()
	}
	return nil
}

func (d *DeviceClient) openAndSendQueryHttp(query *pb.HttpQuery) (*QueryResponse, error) {
	data, err := proto.Marshal(query)
	if err != nil {
		return nil, err
	}

	// EncodeRawBytes includes the length prefix.
	buf := proto.NewBuffer(make([]byte, 0))
	buf.EncodeRawBytes(data)

	body := bytes.NewReader(buf.Bytes())
	contentType := "application/fkhttp"

	if d.HexEncode {
		hexEncoded := hex.EncodeToString(buf.Bytes())
		body = bytes.NewReader([]byte(hexEncoded))
		contentType = "text/plain"
	}

	// We only write once, so this single call is fine here. If we need to
	// write more in the future we'll need another one of these.
	url := fmt.Sprintf("http://%s:%d/fk/v1", d.Address, d.Port)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Content-Length", fmt.Sprintf("%d", len(buf.Bytes())))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return &QueryResponse{http: resp}, nil
}

func (d *DeviceClient) openAndSendQuery(opts *DeviceQueryOpts) (*QueryResponse, error) {
	return d.openAndSendQueryHttp(opts.httpQuery)
}

type DeadlineReader struct {
	Target  net.Conn
	Timeout time.Duration
}

func (dr *DeadlineReader) Read(p []byte) (n int, err error) {
	dr.Target.SetReadDeadline(time.Now().Add(dr.Timeout))
	n, err = dr.Target.Read(p)
	return
}

type DebugReader struct {
	Target io.Reader
}

func (dr *DebugReader) Read(p []byte) (n int, err error) {
	n, err = dr.Target.Read(p)
	return
}

type DeviceQueryOpts struct {
	httpQuery *pb.HttpQuery
	timeout   int
}

func (d *DeviceClient) queryDeviceCallback(opts *DeviceQueryOpts, callback CallbackFunc) (collection.MessageCollection, error) {
	c, err := d.openAndSendQuery(opts)
	if err != nil {
		return nil, err
	}

	defer c.Close()

	dr := c.Reader(opts.timeout)

	if d.HexEncode {
		dr = hex.NewDecoder(dr)
	}

	unmarshalFunc := message.UnmarshalFunc(func(b []byte) (proto.Message, error) {
		var reply pb.HttpReply
		err := proto.Unmarshal(b, &reply)
		if err != nil {
			return nil, err
		}

		return &reply, nil
	})

	collection, err := stream.ReadLengthPrefixedCollection(dr, unmarshalFunc)
	if err != nil {
		return nil, err
	}

	if len(collection) == 0 {
		return nil, fmt.Errorf("No reply.")
	}

	return collection, err
}

func (d *DeviceClient) queryDeviceOpts(opts *DeviceQueryOpts) (reply *pb.HttpReply, err error) {
	if d.Callbacks != nil {
		d.Callbacks.Sent(opts.httpQuery)
	}
	replies, err := d.queryDeviceMultiple(opts)
	if err != nil {
		return nil, err
	}
	if len(replies) > 0 {
		if d.Callbacks != nil {
			for _, m := range replies {
				replies = append(replies, m.(*pb.HttpReply))
				d.Callbacks.Received(m.(*pb.HttpReply))
			}
		}
		return nil, nil
	}
	return nil, nil
}

func (d *DeviceClient) queryDevice(queryType pb.QueryType) (reply *pb.HttpReply, err error) {
	query := &pb.HttpQuery{
		Type: queryType,
	}
	return d.queryDeviceOpts(&DeviceQueryOpts{query, DefaultTimeout})
}

func (d *DeviceClient) queryDeviceQuery(query *pb.HttpQuery) (reply *pb.HttpReply, err error) {
	return d.queryDeviceOpts(&DeviceQueryOpts{query, DefaultTimeout})
}

func (d *DeviceClient) queryDeviceMultiple(opts *DeviceQueryOpts) (replies collection.MessageCollection, err error) {
	return d.queryDeviceCallback(opts, CallbackFunc(func(reply *pb.WireMessageReply) error {
		return nil
	}))
}
