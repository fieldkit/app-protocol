package main

import (
	"flag"
	"log"
	"os"

	fkc "github.com/fieldkit/app-protocol/fkdevice"
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

	LoraAppKey string
	LoraAppEui string
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
		_, err := device.ConfigureLora(o.LoraAppKey, o.LoraAppEui)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}
}
