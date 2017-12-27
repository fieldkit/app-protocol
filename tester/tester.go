package main

import (
	"context"
	"flag"
	pb "github.com/fieldkit/app-protocol"
	fkc "github.com/fieldkit/app-protocol/device"
	"github.com/fieldkit/cloud/testing"
	"log"
	"os"
	"time"
)

type options struct {
	Address string
	Port    int

	Scan        bool
	Status      bool
	Network     bool
	ConfigureAp bool
	Schedules   bool

	DownloadDataSet int
	EraseDataSet    int

	LiveDataPoll     bool
	LiveDataInterval int

	Files        bool
	DownloadFile int

	Identity bool
	Device   string
	Stream   string

	Reset bool

	Host         string
	Scheme       string
	Username     string
	Password     string
	Project      string
	DeviceName   string
	LinkIdentity bool
}

func main() {
	o := options{}

	flag.StringVar(&o.Address, "address", "", "ip address of the device")
	flag.IntVar(&o.Port, "port", 54321, "port number")

	flag.BoolVar(&o.Scan, "scan", false, "scan the device's capabilities and data sets")
	flag.BoolVar(&o.Schedules, "schedules", false, "query for schedules")
	flag.BoolVar(&o.Status, "status", false, "device status")

	flag.BoolVar(&o.Network, "network", false, "query the device's network settings")
	flag.BoolVar(&o.ConfigureAp, "network-configure-ap", false, "force network ap mode (empty configured networks)")

	flag.BoolVar(&o.Files, "files", false, "scan the device's files")
	flag.IntVar(&o.DownloadFile, "download-file", -1, "download file")

	flag.IntVar(&o.DownloadDataSet, "download-ds", -1, "download data")
	flag.IntVar(&o.EraseDataSet, "erase-ds", -1, "erase data")

	flag.IntVar(&o.LiveDataInterval, "live-data-interval", 1000, "interval to poll (0 disables)")
	flag.BoolVar(&o.LiveDataPoll, "live-data-poll", false, "send live data poll")

	flag.BoolVar(&o.Identity, "identity", false, "retrieve the device's identity")
	flag.StringVar(&o.Device, "device", "", "device identity")
	flag.StringVar(&o.Stream, "stream", "", "stream identity")

	flag.BoolVar(&o.Reset, "reset", false, "reset the device")

	flag.StringVar(&o.Host, "host", "127.0.0.1:8080", "hostname to use")
	flag.StringVar(&o.Scheme, "scheme", "http", "scheme to use")
	flag.StringVar(&o.Username, "username", "demo-user", "username to use")
	flag.StringVar(&o.Password, "password", "asdfasdfasdf", "password to use")
	flag.StringVar(&o.Project, "project", "www", "project")
	flag.StringVar(&o.DeviceName, "device-name", "test-device", "device name")
	flag.BoolVar(&o.LinkIdentity, "link", false, "link device identity to website")
	flag.Parse()

	if o.Address == "" {
		flag.Usage()
		os.Exit(2)
	}

	device := &fkc.DeviceClient{
		Address: o.Address,
		Port:    o.Port,
	}

	if o.Scan {
		_, err := device.QueryCapabilities()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		dataSets, err := device.QueryDataSets()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		for _, ds := range dataSets.DataSets.DataSets {
			_, err := device.QueryDataSet(ds.Id)
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
		}
	}

	if o.Status {
		_, err := device.QueryStatus()
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

	if o.DownloadDataSet >= 0 {
		ds, err := device.QueryDataSet(uint32(o.DownloadDataSet))
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		device.DownloadDataSet(uint32(o.DownloadDataSet), ds.DataSets.DataSets[0].Pages)
	}

	if o.EraseDataSet >= 0 {
		_, err := device.QueryDataSet(uint32(o.EraseDataSet))
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		_, err = device.EraseDataSet(uint32(o.EraseDataSet))
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
		_, err := device.DownloadFile(o.DownloadFile)
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

	if o.LinkIdentity {
		ctx := context.TODO()
		c, err := testing.CreateAndAuthenticate(ctx, o.Host, o.Scheme, o.Username, o.Password)
		if err != nil {
			log.Fatalf("%v", err)
		}

		webDevice, err := testing.CreateWebDevice(ctx, c, o.Project, o.DeviceName)
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
}
