package fkdevice

import (
	"encoding/json"
	"io/ioutil"
	"log"

	pb "github.com/fieldkit/app-protocol"
)

type DeviceClientLoggingCallbacks interface {
	Sent(query *pb.HttpQuery)
	Received(reply *pb.HttpReply)
}

type LogJsonCallbacks struct {
	Save string
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

	if cb.Save != "" {
		err := ioutil.WriteFile(cb.Save, []byte(replyJson), 0644)
		if err != nil {
			panic(err)
		}
	}
}
