package main

import (
	"encoding/json"
	"flag"
	"fmt"
	pb "github.com/fieldkit/app-protocol"
	"github.com/golang/protobuf/proto"
	"log"
	"net"
	"os"
	"time"
)

type options struct {
	LiveDataPoll     bool
	LiveDataInterval int
	Address          string
	Port             int
	DataSetId        int
	Scan             bool
	Files            bool
	DownloadFile     int
	DownloadData     bool
	EraseData        bool
	Schedules        bool
	// Continuous       bool
	// Queries          int
	// All              bool
}

type DeviceClient struct {
	Address string
	Port    int
}

func (d *DeviceClient) queryDevice(query *pb.WireMessageQuery, echoReplyJson bool) (reply *pb.WireMessageReply, err error) {
	c, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", d.Address, d.Port), 5*time.Second)
	if err != nil {
		return nil, err
	}

	defer c.Close()

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
	_, err = c.Write(encoded)
	if err != nil {
		return nil, err
	}

	data = make([]byte, 4096)
	length, err := c.Read(data)
	if err != nil {
		return nil, fmt.Errorf("Unable to read %v", err)
	}

	sliced := data[0:length]

	buf = proto.NewBuffer(sliced)
	_, err = buf.DecodeVarint()
	if err != nil {
		return nil, fmt.Errorf("Unable to read length %v", err)
	}

	reply = new(pb.WireMessageReply)
	err = buf.Unmarshal(reply)
	if err != nil {
		log.Printf("Length: %v", len(buf.Bytes()))
		log.Printf("Bytes: %v", buf.Bytes())
		return nil, fmt.Errorf("Unable to Unmarshal %v", err)
	}

	replyJson, err := json.MarshalIndent(reply, "", "  ")
	if err == nil {
		if echoReplyJson {
			log.Printf("Received: %s", replyJson)
		}
	}

	if reply.Type == pb.ReplyType_REPLY_ERROR {
		log.Fatalf("Error")
	}

	return
}

func (d *DeviceClient) queryCapabilities() (*pb.WireMessageReply, error) {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_CAPABILITIES,
	}
	reply, err := d.queryDevice(query, true)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) querySchedules() (*pb.WireMessageReply, error) {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_SCHEDULES,
	}
	reply, err := d.queryDevice(query, true)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) queryDataSets() (*pb.WireMessageReply, error) {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_DATA_SETS,
	}
	reply, err := d.queryDevice(query, true)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) queryFiles() (*pb.WireMessageReply, error) {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_FILES,
	}
	reply, err := d.queryDevice(query, true)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (d *DeviceClient) downloadFile(id int) (*pb.WireMessageReply, error) {
	query := &pb.WireMessageQuery{
		Type: pb.QueryType_QUERY_FILES,
	}
	files, err := d.queryDevice(query, true)
	if err != nil {
		return nil, err
	}

	for _, file := range files.Files.Files { // Files, files, files!
		if int(file.Id) == id {
			token := []byte{}
			for page := 0; uint32(page) < file.Pages; page += 1 {
				query := &pb.WireMessageQuery{
					Type: pb.QueryType_QUERY_DOWNLOAD_FILE,
					DownloadFile: &pb.DownloadFile{
						Id:    uint32(id),
						Page:  uint32(page),
						Token: token,
					},
				}
				reply, err := d.queryDevice(query, false)
				if err != nil {
					return nil, err
				}

				fmt.Printf("Page#%d: %+v (%d bytes)\n", page, reply.FileData.Token, len(reply.FileData.Data))

				token = reply.FileData.Token

				if len(reply.FileData.Data) == 0 {
					break
				}
			}
		}
	}

	return files, nil
}

func (d *DeviceClient) queryDataSet(id uint32) (*pb.WireMessageReply, error) {
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

func (d *DeviceClient) queryLiveData(liveDataInterval int) (*pb.WireMessageReply, error) {
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

func (d *DeviceClient) eraseDataSet(id uint32) (*pb.WireMessageReply, error) {
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

func (d *DeviceClient) downloadDataSet(id uint32, pages uint32) (*pb.WireMessageReply, error) {
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

func main() {
	o := options{}

	flag.StringVar(&o.Address, "address", "", "ip address of the device")
	flag.IntVar(&o.Port, "port", 12345, "port number")
	flag.BoolVar(&o.Scan, "scan", false, "scan the device's capabilities and data sets")
	flag.BoolVar(&o.Files, "files", false, "scan the device's files")
	flag.IntVar(&o.DownloadFile, "download-file", -1, "download file")
	flag.BoolVar(&o.Schedules, "schedules", false, "query for schedules")
	flag.BoolVar(&o.DownloadData, "download-data", false, "download data")
	flag.BoolVar(&o.EraseData, "erase-data", false, "erase data")
	flag.IntVar(&o.DataSetId, "data-set", 0, "data set id to download or erase")
	flag.IntVar(&o.LiveDataInterval, "live-data-interval", 1000, "interval to poll (0 disables)")
	flag.BoolVar(&o.LiveDataPoll, "live-data-poll", false, "send live data poll")
	/*
		flag.IntVar(&o.Queries, "queries", 1, "number of times to query")
		flag.BoolVar(&o.Continuous, "continuous", false, "keep trying")
		flag.BoolVar(&o.All, "all", false, "try a bunch of queries")
	*/
	flag.Parse()

	if o.Address == "" {
		flag.Usage()
		os.Exit(2)
	}

	device := &DeviceClient{
		Address: o.Address,
		Port:    o.Port,
	}

	if o.Scan {
		_, err := device.queryCapabilities()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		dataSets, err := device.queryDataSets()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		for _, ds := range dataSets.DataSets.DataSets {
			_, err := device.queryDataSet(ds.Id)
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
		}
	}

	if o.Files {
		_, err := device.queryFiles()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	if o.DownloadFile >= 0 {
		_, err := device.downloadFile(o.DownloadFile)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	if o.Schedules {
		_, err := device.querySchedules()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	if o.LiveDataPoll {
		for {
			_, err := device.queryLiveData(o.LiveDataInterval)
			if err != nil {
				log.Fatalf("Error: %v", err)
			}

			if o.LiveDataInterval > 0 {
				time.Sleep(time.Duration(o.LiveDataInterval) * time.Millisecond)
			} else {
				break
			}
		}
	}

	if o.DownloadData {
		ds, err := device.queryDataSet(uint32(o.DataSetId))
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		device.downloadDataSet(uint32(o.DataSetId), ds.DataSets.DataSets[0].Pages)
	}

	if o.EraseData {
		_, err := device.queryDataSet(uint32(o.DataSetId))
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		_, err = device.eraseDataSet(uint32(o.DataSetId))
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}
}
