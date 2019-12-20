package fkdevice

import (
	"io"
	"net"
	"net/http"
	"time"
)

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
	return dr.Target.Read(p)
}

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
