package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"

	fkc "github.com/fieldkit/app-protocol/fkdevice"

	pb "github.com/fieldkit/app-protocol"
	atlaspb "github.com/fieldkit/atlas-protocol"
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

	Module int

	Atlas          string
	AtlasClear     bool
	AtlasStatus    bool
	AtlasCalibrate int
	AtlasValue     float64

	FactoryReset bool

	TransmissionUrl   string
	TransmissionToken string
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

	flag.IntVar(&o.Module, "module", -1, "module")
	flag.StringVar(&o.Atlas, "atlas", "", "atlas")
	flag.BoolVar(&o.AtlasClear, "atlas-clear", false, "atlas-clear")
	flag.BoolVar(&o.AtlasStatus, "atlas-status", false, "atlas-status")
	flag.IntVar(&o.AtlasCalibrate, "atlas-calibrate", 0, "atlas-calibrate")
	flag.Float64Var(&o.AtlasValue, "atlas-value", 0, "atlas-value")

	flag.StringVar(&o.TransmissionUrl, "transmission-url", "", "transmission url")
	flag.StringVar(&o.TransmissionToken, "transmission-token", "", "transmission token")

	flag.BoolVar(&o.FactoryReset, "factory-reset", false, "factory reset")

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

	if o.FactoryReset {
		_, err := device.FactoryReset()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	if o.Atlas != "" {
		if o.Module < 0 {
			log.Fatalf("Error: module required")
		}

		query := &atlaspb.WireAtlasQuery{}
		rawQuery, err := proto.Marshal(query)
		if err != nil {
			panic(err)
		}

		rawReply, err := device.ModuleQuery(uint32(o.Module), rawQuery)
		if err != nil {
			panic(err)
		}

		log.Printf("%v", rawReply)

		buffer := proto.NewBuffer(rawReply)

		var reply atlaspb.WireAtlasReply
		err = buffer.DecodeMessage(&reply)
		if err != nil {
			panic(err)
		}

		log.Printf("%v", reply)
	}

	if o.TransmissionUrl != "" || o.TransmissionToken != "" {
		_, err := device.ConfigureTransmission(o.TransmissionUrl, o.TransmissionToken)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	if o.Module >= 0 {
		if o.AtlasStatus {
			query := &atlaspb.WireAtlasQuery{
				Type: atlaspb.QueryType_QUERY_NONE,
				Calibration: &atlaspb.AtlasCalibrationCommand{
					Operation: atlaspb.CalibrationOperation_CALIBRATION_STATUS,
				},
			}

			reply, err := atlasQuery(device, uint32(o.Module), query)
			if err != nil {
				panic(err)
			}

			log.Printf("reply: %+v", reply)
		}

		if o.AtlasClear {
			query := &atlaspb.WireAtlasQuery{
				Type: atlaspb.QueryType_QUERY_NONE,
				Calibration: &atlaspb.AtlasCalibrationCommand{
					Operation: atlaspb.CalibrationOperation_CALIBRATION_CLEAR,
				},
			}

			reply, err := atlasQuery(device, uint32(o.Module), query)
			if err != nil {
				panic(err)
			}

			log.Printf("reply: %+v", reply)
		}

		if o.AtlasCalibrate > 0 {
			query := &atlaspb.WireAtlasQuery{
				Type: atlaspb.QueryType_QUERY_NONE,
				Calibration: &atlaspb.AtlasCalibrationCommand{
					Operation: atlaspb.CalibrationOperation_CALIBRATION_SET,
					Which:     uint32(o.AtlasCalibrate),
					Value:     float32(o.AtlasValue),
				},
			}

			reply, err := atlasQuery(device, uint32(o.Module), query)
			if err != nil {
				panic(err)
			}

			log.Printf("reply: %+v", reply)
		}
	}
}

func atlasQuery(device *fkc.DeviceClient, module uint32, query *atlaspb.WireAtlasQuery) (reply *atlaspb.WireAtlasReply, err error) {
	rawQuery, err := proto.Marshal(query)
	if err != nil {
		return nil, err
	}

	log.Printf("query: %v", rawQuery)

	rawReply, err := device.ModuleQuery(module, rawQuery)
	if err != nil {
		return nil, err
	}

	log.Printf("reply: %v", rawReply)

	reply = &atlaspb.WireAtlasReply{}
	buffer := proto.NewBuffer(rawReply)

	err = buffer.DecodeMessage(reply)
	if err != nil {
		return nil, err
	}

	return
}
