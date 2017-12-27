package fkdevice

import (
	"bytes"
	"encoding/json"
	"fmt"
	pb "github.com/fieldkit/app-protocol"
	"github.com/golang/protobuf/proto"
	"io"
	"log"
	"net"
	"os"
	"time"
)

type DeviceClient struct {
	Address string
	Port    int
}

func (d *DeviceClient) QueryCapabilities() (*pb.WireMessageReply, error) {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_CAPABILITIES,
	}
	reply, err := d.queryDevice(query, true)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) QueryStatus() (*pb.WireMessageReply, error) {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_STATUS,
	}
	reply, err := d.queryDevice(query, true)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) QuerySchedules() (*pb.WireMessageReply, error) {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_SCHEDULES,
	}
	reply, err := d.queryDevice(query, true)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) QueryDataSets() (*pb.WireMessageReply, error) {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_DATA_SETS,
	}
	reply, err := d.queryDevice(query, true)
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
	reply, err := d.queryDevice(query, true)
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
	reply, err := d.queryDevice(query, true)
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
		reply, err := d.queryDevice(query, false)
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
	reply, err := d.queryDevice(query, true)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) QueryFiles() (*pb.WireMessageReply, error) {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_FILES,
	}
	reply, err := d.queryDevice(query, true)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) DownloadFile(id int) (*pb.WireMessageReply, error) {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_FILES,
	}
	files, err := d.queryDevice(query, true)
	if err != nil {
		return nil, err
	}

	time.Sleep(500 * time.Millisecond)

	for _, file := range files.Files.Files { // Files, files, files!
		if int(file.Id) == id {
			fileName := fmt.Sprintf("%s_%d", file.Name, file.Version)
			f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
			if err != nil {
				log.Fatalf("Unable to open %s (%v)", fileName, err)
			}

			defer f.Close()

			retry := true
			token := []byte{}
			page := 0
			downloaded := 0
			for finished := false; !finished; {
				query := &pb.WireMessageQuery{
					Type: pb.QueryType_QUERY_DOWNLOAD_FILE,
					DownloadFile: &pb.DownloadFile{
						Id:    uint32(id),
						Page:  uint32(page),
						Token: token,
					},
				}
				replies, err := d.queryDeviceMultiple(query, false)
				if err != nil {
					if !retry {
						return nil, err
					}
					log.Printf("%s", err)
					time.Sleep(6 * time.Second)
					continue
				}

				for _, reply := range replies {
					log.Printf("Page#%d: (%d bytes) (%.2f%%)\n", page, len(reply.FileData.Data), 100.0*(float32(downloaded)/float32(file.Size)))

					if len(reply.FileData.Data) > 0 {
						wrote, err := f.Write(reply.FileData.Data)
						if err != nil {
							log.Fatalf("Error writing: %v", err)
						}
						downloaded += wrote
					}

					if bytes.Equal(token, reply.FileData.Token) {
						fmt.Printf("Page#%d: %+v (got same token back)\n", page, reply.FileData.Token)
						finished = true
						break
					}

					token = reply.FileData.Token
					page += 1
				}

				time.Sleep(100 * time.Millisecond)

				if page == 4 {
					break
				}
			}
		}
	}

	return files, nil
}

func (d *DeviceClient) QueryLiveData(liveDataInterval int) (*pb.WireMessageReply, error) {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_LIVE_DATA_POLL,
		LiveDataPoll: &pb.LiveDataPoll{
			Interval: uint32(liveDataInterval),
		},
	}
	reply, err := d.queryDevice(query, true)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) QueryNetworkSettings() (*pb.WireMessageReply, error) {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_NETWORK_SETTINGS,
	}
	reply, err := d.queryDevice(query, true)
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
	reply, err := d.queryDevice(query, true)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) QueryIdentity() (*pb.WireMessageReply, error) {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_IDENTITY,
	}

	reply, err := d.queryDevice(query, true)
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
	reply, err := d.queryDevice(query, true)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) queryDevice(query *pb.WireMessageQuery, echoReplyJson bool) (reply *pb.WireMessageReply, err error) {
	replies, err := d.queryDeviceMultiple(query, echoReplyJson)
	if err != nil {
		return nil, err
	}
	if len(replies) > 0 {
		return replies[0], nil
	}
	return nil, nil
}

func (d *DeviceClient) queryDeviceMultiple(query *pb.WireMessageQuery, echoReplyJson bool) (replies []*pb.WireMessageReply, err error) {
	c, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", d.Address, d.Port), 5*time.Second)
	if err != nil {
		return nil, err
	}

	defer c.Close()

	replies = make([]*pb.WireMessageReply, 0)

	queryJson, err := json.MarshalIndent(query, "", "  ")
	if err == nil {
		log.Printf("Sending: %s", queryJson)
	}

	data, err := proto.Marshal(query)
	if err != nil {
		return nil, err
	}

	// EncodeRawBytes includes the varint length.
	buf := proto.NewBuffer(make([]byte, 0))
	buf.EncodeRawBytes(data)

	encoded := buf.Bytes()
	wrote, err := c.Write(encoded)
	if err != nil {
		return nil, err
	}

	data = make([]byte, 0)

	for {
		page := make([]byte, 8192)
		length, err := c.Read(page)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("Unable to read %v (%d)", err, wrote)
		} else {
			data = append(data, page[0:length]...)
		}
	}

	log.Printf("Read %v bytes", len(data))

	buf = proto.NewBuffer(data[:])

	for {
		messageBytes, err := buf.DecodeRawBytes(true)
		if err != nil {
			if err == io.ErrUnexpectedEOF {
				err = nil
				break
			}
			return nil, fmt.Errorf("Unable to read length %v", err)
		}

		messageBuffer := proto.NewBuffer(messageBytes[:])
		reply := new(pb.WireMessageReply)
		err = messageBuffer.Unmarshal(reply)
		if err != nil {
			log.Printf("Length: %v", len(messageBuffer.Bytes()))
			if echoReplyJson {
				log.Printf("Bytes: %v", messageBuffer.Bytes())
			}
			return nil, fmt.Errorf("Unable to Unmarshal %v", err)
		}

		replies = append(replies, reply)

		replyJson, err := json.MarshalIndent(reply, "", "  ")
		if err == nil {
			if echoReplyJson {
				log.Printf("Received: %s", replyJson)
			}
		}
	}

	return
}
