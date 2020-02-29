package fkdevice

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
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

type DeviceClient struct {
	Callbacks DeviceClientLoggingCallbacks
	Address   string
	HexEncode bool
	Port      int
}

func (d *DeviceClient) QueryStatus() (*pb.HttpReply, error) {
	reply, err := d.queryDeviceSimple(pb.QueryType_QUERY_STATUS)
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
	reply, err := d.queryDevice(query)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) QueryStartRecording() (*pb.HttpReply, error) {
	query := &pb.HttpQuery{
		Type: pb.QueryType_QUERY_RECORDING_CONTROL,
		Recording: &pb.Recording{
			Modifying: true,
			Enabled:   true,
		},
	}
	reply, err := d.queryDevice(query)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) QueryStopRecording() (*pb.HttpReply, error) {
	query := &pb.HttpQuery{
		Type: pb.QueryType_QUERY_RECORDING_CONTROL,
		Recording: &pb.Recording{
			Modifying: true,
			Enabled:   false,
		},
	}
	reply, err := d.queryDevice(query)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) ConfigureSchedule(readings, network, gps, lora uint32) (*pb.HttpReply, error) {
	query := &pb.HttpQuery{
		Type: pb.QueryType_QUERY_CONFIGURE,
		Schedules: &pb.Schedules{
			Modifying: true,
			Readings: &pb.Schedule{
				Interval: readings,
			},
			Network: &pb.Schedule{
				Interval: network,
			},
			Gps: &pb.Schedule{
				Interval: gps,
			},
			Lora: &pb.Schedule{
				Interval: lora,
			},
		},
	}
	reply, err := d.queryDevice(query)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) ConfigureWifiNetworks(networks []*pb.NetworkInfo) (*pb.HttpReply, error) {
	query := &pb.HttpQuery{
		Type: pb.QueryType_QUERY_CONFIGURE,
		NetworkSettings: &pb.NetworkSettings{
			Networks: networks,
		},
	}
	reply, err := d.queryDevice(query)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) ConfigureTransmission(url, token string) (*pb.HttpReply, error) {
	query := &pb.HttpQuery{
		Type: pb.QueryType_QUERY_CONFIGURE,
		Transmission: &pb.Transmission{
			Wifi: &pb.WifiTransmission{
				Modifying: true,
				Url:       url,
				Token:     token,
			},
		},
	}
	reply, err := d.queryDevice(query)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) ConfigureLoraAbp(appSessionKey, networkSessionKey, deviceAddress string, uplinkCounter, downlinkCounter uint32) (*pb.HttpReply, error) {
	appSessionKeyBytes, err := hex.DecodeString(appSessionKey)
	if err != nil {
		return nil, err
	}

	networkSessionKeyBytes, err := hex.DecodeString(networkSessionKey)
	if err != nil {
		return nil, err
	}

	deviceAddressBytes, err := hex.DecodeString(deviceAddress)
	if err != nil {
		return nil, err
	}

	log.Printf("AppSessionKey %+v", appSessionKeyBytes)
	log.Printf("NetworkSessionKey %+v", networkSessionKeyBytes)
	log.Printf("DeviceAddress %+v", deviceAddressBytes)

	query := &pb.HttpQuery{
		Type: pb.QueryType_QUERY_CONFIGURE,
		LoraSettings: &pb.LoraSettings{
			Modifying:         true,
			AppSessionKey:     appSessionKeyBytes,
			NetworkSessionKey: networkSessionKeyBytes,
			DeviceAddress:     deviceAddressBytes,
			UplinkCounter:     uplinkCounter,
			DownlinkCounter:   downlinkCounter,
		},
	}
	reply, err := d.queryDevice(query)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) ConfigureLoraOtaa(appKey, appEui string) (*pb.HttpReply, error) {
	appKeyBytes, err := hex.DecodeString(appKey)
	if err != nil {
		return nil, err
	}

	appEuiBytes, err := hex.DecodeString(appEui)
	if err != nil {
		return nil, err
	}

	log.Printf("EUI %+v", appEuiBytes)
	log.Printf("KEY %+v", appKeyBytes)

	query := &pb.HttpQuery{
		Type: pb.QueryType_QUERY_CONFIGURE,
		LoraSettings: &pb.LoraSettings{
			Modifying: true,
			AppKey:    appKeyBytes,
			AppEui:    appEuiBytes,
		},
	}
	reply, err := d.queryDevice(query)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) QueryGetReadings() (*pb.HttpReply, error) {
	reply, err := d.queryDeviceSimple(pb.QueryType_QUERY_GET_READINGS)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) QueryTakeReadings() (*pb.HttpReply, error) {
	reply, err := d.queryDeviceSimple(pb.QueryType_QUERY_TAKE_READINGS)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) FactoryReset() (*pb.HttpReply, error) {
	reply, err := d.queryDeviceSimple(pb.QueryType_QUERY_RESET)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) ModuleQuery(bay uint32, query []byte) ([]byte, error) {
	return d.queryModule(bay, query)
}

var DeviceBusyErr = fmt.Errorf("Busy")

func (d *DeviceClient) openAndSendQuery(opts *DeviceQueryOpts) (*QueryResponse, error) {
	// EncodeRawBytes includes the length prefix.
	buf := proto.NewBuffer(make([]byte, 0))
	buf.EncodeRawBytes(opts.query)

	body := bytes.NewReader(buf.Bytes())
	contentType := "application/fkhttp"

	if d.HexEncode {
		hexEncoded := hex.EncodeToString(buf.Bytes())
		body = bytes.NewReader([]byte(hexEncoded))
		contentType = "text/plain"
	}

	// We only write once, so this single call is fine here. If we need to
	// write more in the future we'll need another one of these.
	url := fmt.Sprintf("http://%s:%d/%s", d.Address, d.Port, opts.path)
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

type DeviceQueryOpts struct {
	path    string
	query   []byte
	timeout int
}

func (d *DeviceClient) queryDeviceSimple(queryType pb.QueryType) (reply *pb.HttpReply, err error) {
	query := &pb.HttpQuery{
		Type: queryType,
	}
	return d.queryDevice(query)
}

func (d *DeviceClient) queryDevice(query *pb.HttpQuery) (reply *pb.HttpReply, err error) {
	query.Time = uint64(time.Now().Unix())

	data, err := proto.Marshal(query)
	if err != nil {
		return nil, err
	}

	if d.Callbacks != nil {
		d.Callbacks.Sent(query)
	}

	return d.queryDeviceOpts(&DeviceQueryOpts{"fk/v1", data, DefaultTimeout})
}

func (d *DeviceClient) queryDeviceCallback(opts *DeviceQueryOpts) (collection.MessageCollection, error) {
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
	replies, err := d.queryDeviceCallback(opts)

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
		return replies[0].(*pb.HttpReply), nil
	}
	return nil, nil
}

func (d *DeviceClient) queryModuleCallback(opts *DeviceQueryOpts) ([]byte, error) {
	c, err := d.openAndSendQuery(opts)
	if err != nil {
		return nil, err
	}

	defer c.Close()

	dr := c.Reader(opts.timeout)

	return ioutil.ReadAll(dr)
}

func (d *DeviceClient) queryModule(module uint32, query []byte) (rawReply []byte, err error) {
	return d.queryModuleCallback(&DeviceQueryOpts{fmt.Sprintf("fk/v1/module/%d", module), query, DefaultTimeout})
}
