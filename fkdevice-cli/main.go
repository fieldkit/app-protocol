package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"strings"

	fkc "github.com/fieldkit/app-protocol/fkdevice"

	pb "github.com/fieldkit/app-protocol"
)

type options struct {
	Address   string
	Port      int
	Status    bool
	HexEncode bool
	Name      string
	Save      string

	GetReadings    bool
	TakeReadings   bool
	StartRecording bool
	StopRecording  bool

	Schedule string
	Networks bool

	LoraAppKey            string
	LoraAppEui            string
	LoraAppSessionKey     string
	LoraNetworkSessionKey string
	LoraDeviceAddress     string
	LoraUplinkCounter     int
	LoraDownlinkCounter   int
}

func main() {
	o := options{}

	flag.StringVar(&o.Address, "address", "", "ip address of the device")
	flag.IntVar(&o.Port, "port", 80, "port number")
	flag.BoolVar(&o.HexEncode, "hex", false, "hex encoding")
	flag.BoolVar(&o.Status, "status", false, "device status")
	flag.StringVar(&o.Name, "name", "", "name")
	flag.BoolVar(&o.GetReadings, "get", false, "")
	flag.BoolVar(&o.TakeReadings, "take", false, "")
	flag.BoolVar(&o.StartRecording, "start-recording", false, "")
	flag.BoolVar(&o.StopRecording, "stop-recording", false, "")
	flag.StringVar(&o.Save, "save", "", "save")
	flag.StringVar(&o.LoraAppKey, "lora-app-key", "", "lora-app-key")
	flag.StringVar(&o.LoraAppEui, "lora-app-eui", "", "lora-app-eui")
	flag.StringVar(&o.LoraAppSessionKey, "lora-app-session-key", "", "lora-app-session-key")
	flag.StringVar(&o.LoraNetworkSessionKey, "lora-network-session-key", "", "lora-network-session-key")
	flag.StringVar(&o.LoraDeviceAddress, "lora-device-address", "", "lora-device-address")
	flag.IntVar(&o.LoraUplinkCounter, "lora-uplink-counter", 0, "lora-uplink-counter")
	flag.IntVar(&o.LoraDownlinkCounter, "lora-downlink-counter", 0, "lora-downlink-counter")
	flag.StringVar(&o.Schedule, "schedule", "", "schedule")
	flag.BoolVar(&o.Networks, "networks", false, "")

	flag.Parse()

	if o.Address == "" {
		flag.Usage()
		os.Exit(2)
	}

	device := &fkc.DeviceClient{
		Address: o.Address,
		Port:    o.Port,
		Callbacks: &fkc.LogJsonCallbacks{
			Save: o.Save,
		},
		HexEncode: o.HexEncode,
	}

	if o.Status {
		_, err := device.QueryStatus()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	if o.Name != "" {
		_, err := device.ConfigureName(o.Name)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	if o.GetReadings {
		_, err := device.QueryGetReadings()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	if o.TakeReadings {
		_, err := device.QueryTakeReadings()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	if o.StartRecording {
		_, err := device.QueryStartRecording()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	if o.StopRecording {
		_, err := device.QueryStopRecording()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	if o.LoraAppKey != "" && o.LoraAppEui != "" {
		_, err := device.ConfigureLoraOtaa(o.LoraAppKey, o.LoraAppEui)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	if o.LoraAppSessionKey != "" && o.LoraNetworkSessionKey != "" && o.LoraDeviceAddress != "" {
		_, err := device.ConfigureLoraAbp(o.LoraAppSessionKey, o.LoraNetworkSessionKey, o.LoraDeviceAddress, uint32(o.LoraUplinkCounter), uint32(o.LoraDownlinkCounter))
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	if o.Schedule != "" {
		f := strings.Split(o.Schedule, ",")
		if len(f) == 4 {
			readings, _ := strconv.Atoi(f[0])
			network, _ := strconv.Atoi(f[1])
			gps, _ := strconv.Atoi(f[2])
			lora, _ := strconv.Atoi(f[3])
			_, err := device.ConfigureSchedule(uint32(readings), uint32(network), uint32(gps), uint32(lora))
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
		}
	}

	if o.Networks {
		networks := []*pb.NetworkInfo{
			&pb.NetworkInfo{
				Ssid:     "Cottonwood",
				Password: "asdfasdf",
			},
			&pb.NetworkInfo{
				Ssid:     "Conservify",
				Password: "Okavang0",
			},
		}
		_, err := device.ConfigureWifiNetworks(networks)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}
}
