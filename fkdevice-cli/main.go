package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/fieldkit/app-protocol"
	fkc "github.com/fieldkit/app-protocol/fkdevice"
	fkatlaspb "github.com/fieldkit/atlas-protocol"
	testing "github.com/fieldkit/cloud/server/tools"
	"github.com/golang/protobuf/proto"
	progress "gopkg.in/cheggaaa/pb.v1"
	"io"
	"log"
	"os"
	"time"
)

type options struct {
	Address  string
	Port     int
	HttpMode bool

	Scan        bool
	Status      bool
	Network     bool
	ConfigureAp bool
	Schedules   bool

	LiveDataOnce     bool
	LiveDataPoll     bool
	LiveDataInterval int

	Files           bool
	DownloadFile    int
	EraseFile       int
	WriteTo         string
	MetadataOnly    bool
	PrependMetadata bool

	Offset int
	Length int

	Metadata bool

	Identity bool
	Device   string
	Stream   string

	Reset  bool
	Format bool

	DeviceName   string
	LinkIdentity bool

	AtlasQuery int

	Host     string
	Scheme   string
	Username string
	Password string
	Project  string
}

func main() {
	o := options{}

	flag.StringVar(&o.Address, "address", "", "ip address of the device")
	flag.IntVar(&o.Port, "port", 80, "port number")
	flag.BoolVar(&o.HttpMode, "http", true, "use http instead of tcp")

	flag.BoolVar(&o.Scan, "scan", false, "scan the device's capabilities and data sets")
	flag.BoolVar(&o.Schedules, "schedules", false, "query for schedules")
	flag.BoolVar(&o.Status, "status", false, "device status")

	flag.BoolVar(&o.Network, "network", false, "query the device's network settings")
	flag.BoolVar(&o.ConfigureAp, "network-configure-ap", false, "force network ap mode (empty configured networks)")

	flag.StringVar(&o.WriteTo, "write", "", "file to write to")

	flag.BoolVar(&o.Files, "files", false, "scan the device's files")
	flag.IntVar(&o.DownloadFile, "download-file", -1, "download file")
	flag.IntVar(&o.EraseFile, "erase-file", -1, "erase file")
	flag.BoolVar(&o.MetadataOnly, "metadata-only", false, "only get metadata")
	flag.BoolVar(&o.PrependMetadata, "prepend-metadata", false, "include metadata")

	flag.IntVar(&o.Offset, "offset", 0, "download offset")
	flag.IntVar(&o.Length, "length", 0, "download length")

	flag.BoolVar(&o.Metadata, "metadata", false, "query metadata")

	flag.IntVar(&o.LiveDataInterval, "live-data-interval", 1000, "interval to poll (0 disables)")
	flag.BoolVar(&o.LiveDataPoll, "live-data-poll", false, "send live data poll")
	flag.BoolVar(&o.LiveDataOnce, "live-data-once", false, "one live data query")

	flag.BoolVar(&o.Identity, "identity", false, "retrieve the device's identity")
	flag.StringVar(&o.Device, "device", "", "device identity")
	flag.StringVar(&o.Stream, "stream", "", "stream identity")

	flag.BoolVar(&o.Reset, "reset", false, "reset the device")
	flag.BoolVar(&o.Format, "format", false, "format the device")

	flag.StringVar(&o.DeviceName, "device-name", "test-device", "device name")
	flag.BoolVar(&o.LinkIdentity, "link", false, "link device identity to website")

	flag.StringVar(&o.Host, "host", "127.0.0.1:8080", "hostname to use")
	flag.StringVar(&o.Scheme, "scheme", "http", "scheme to use")
	flag.StringVar(&o.Username, "username", "demo-user", "username to use")
	flag.StringVar(&o.Password, "password", "asdfasdfasdf", "password to use")
	flag.StringVar(&o.Project, "project", "www", "project")

	flag.IntVar(&o.AtlasQuery, "atlas-query", -1, "test atlas module queries")

	flag.Parse()

	if o.Address == "" {
		flag.Usage()
		os.Exit(2)
	}

	device := &fkc.DeviceClient{
		Address:   o.Address,
		Port:      o.Port,
		Callbacks: &fkc.LogJsonCallbacks{},
		HttpMode:  o.HttpMode,
	}

	if o.Scan {
		_, err := device.QueryCapabilities()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	if o.Status {
		_, err := device.QueryStatus()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	if o.Metadata {
		_, err := device.QueryMetadata()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	if o.Schedules {
		_, err := device.QuerySchedules()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	if o.Network {
		_, err := device.QueryNetworkSettings()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	if o.ConfigureAp {
		_, err := device.ConfigureNetworkSettings(&pb.NetworkSettings{
			CreateAccessPoint: int32(1),
			Networks:          []*pb.NetworkInfo{},
		})
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	if o.LiveDataPoll {
		for {
			_, err := device.QueryLiveData(o.LiveDataInterval)
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

	if o.LiveDataOnce {
		_, err := device.QueryLiveData(o.LiveDataInterval)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	if o.Files {
		_, err := device.QueryFiles()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	if o.DownloadFile >= 0 {
		file, err := device.QueryFileInformation(uint32(o.DownloadFile))
		if file.Size == 0 {
			log.Fatalf("File is empty, ignoring")
		}

		fileName := o.WriteTo
		if fileName == "" {
			fileName = fmt.Sprintf("%s_%d", file.Name, file.Version)
		}
		log.Printf("Opening %s", fileName)

		f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
		if err != nil {
			log.Fatalf("Unable to open %s (%v)", fileName, err)
		}

		defer f.Close()

		size := int(file.Size)
		if o.Length > 0 {
			size = o.Length
		}
		bar := progress.New(size).SetUnits(progress.U_BYTES)
		bar.Start()

		writer := io.MultiWriter(f, bar)

		flags := 0
		if o.MetadataOnly {
			flags |= int(pb.DownloadFlags_DOWNLOAD_FLAG_METADATA_ONLY)
		}
		if o.PrependMetadata {
			flags |= int(pb.DownloadFlags_DOWNLOAD_FLAG_METADATA_PREPEND)
		}
		err = device.DownloadFileToWriter(uint32(o.DownloadFile), uint32(o.Offset), uint32(o.Length), uint32(flags), writer)
		if err != nil {
			log.Fatalf("Unable to download file %s (%v)", fileName, err)
		}

		// NOTE: This is wrong if they only ask for metadata.
		bar.Set(int(file.Size))
		bar.Finish()
	}

	if o.EraseFile >= 0 {
		_, err := device.EraseFile(uint32(o.EraseFile))
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	if o.Device != "" && o.Stream != "" {
		_, err := device.ConfigureIdentity(o.Device, o.Stream)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	if o.Identity {
		_, err := device.QueryIdentity()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	if o.AtlasQuery >= 0 {
		moduleQuery := &fkatlaspb.WireAtlasQuery{
			Type: fkatlaspb.QueryType_QUERY_ATLAS_COMMAND,
			AtlasCommand: &fkatlaspb.AtlasCommand{
				Sensor:  fkatlaspb.SensorType_TEMP,
				Command: "I",
			},
		}
		data, err := proto.Marshal(moduleQuery)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		reply, err := device.QueryModule(8, data)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		if reply.Module != nil {
			moduleReply := &fkatlaspb.WireAtlasReply{}

			err = proto.Unmarshal(reply.Module.Message, moduleReply)
			if err != nil {
				log.Fatalf("Error: %v", err)
			}

			log.Printf("%v", moduleReply)
		} else {
			log.Printf("%v", reply)
		}
	}

	if o.LinkIdentity {
		ctx := context.TODO()
		c, err := testing.CreateAndAuthenticate(ctx, o.Host, o.Scheme, o.Username, o.Password)
		if err != nil {
			log.Fatalf("%v", err)
		}

		webDevice, err := testing.CreateWebDevice(ctx, c, o.Project, o.DeviceName, "", "1")
		if err != nil {
			log.Fatalf("%v", err)
		}

		log.Printf("%v\n", webDevice)

		newIdentity, err := device.ConfigureIdentity(webDevice.Key, "1")
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		log.Printf("%v\n", newIdentity)
	}

	if o.Reset {
		_, err := device.Reset()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	if o.Format {
		_, err := device.Format()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}
}
