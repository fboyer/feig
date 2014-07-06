package feig

import (
	"fmt"

	"github.com/fboyer/feig/feisc"
	"github.com/fboyer/feig/fetcp"
)

type RFReader interface {
	Open() error
	Read()
	Close()
}

type Reader struct {
	address      string
	port         int
	socketHnd    int
	socketParams map[string]string
	readerHnd    int
	readerParams map[string]string
}

func NewReader(address string, port int, socketParams, readerParams map[string]string) *Reader {
	return &Reader{
		address:      address,
		port:         port,
		socketParams: socketParams,
		readerParams: readerParams,
	}
}

func (r *Reader) Open() error {
	r.socketHnd = fetcp.Connect(r.address, r.port)
	if r.socketHnd < 0 {
		return fmt.Errorf("unable to establish a connection to the feig reader at %s:%d", r.address, r.port)
	}

	for param, value := range r.socketParams {
		result := fetcp.SetSocketParam(r.socketHnd, param, value)
		if result < 0 {
			r.Close()
			return fmt.Errorf("unable to apply the value %s to the parameter %s", value, param)
		}
	}

	r.readerHnd = feisc.NewReader(r.socketHnd)
	if r.readerHnd < 0 {
		r.Close()
		return fmt.Errorf("unable to create a new feig reader")
	}

	for param, value := range r.readerParams {
		result := feisc.SetReaderParam(r.readerHnd, param, value)
		if result < 0 {
			r.Close()
			return fmt.Errorf("unable to apply the value %s to the parameter %s", value, param)
		}
	}

	return nil
}

func (r *Reader) Read() (string, int, error) {
	respData, respLen, result := feisc.SendIsoCmd(r.readerHnd, 255, "0100", 4, 1)
	if result < 0 {
		return "", 0, fmt.Errorf("read error")
	}

	return respData, respLen, nil
}

func (r *Reader) Close() error {
	var result int
	if r.readerHnd > 0 {
		result = feisc.DeleteReader(r.readerHnd)
		if result < 0 {
			fmt.Errorf("unable to delete the reader cleanly")
		}
		r.readerHnd = 0
	}

	if r.socketHnd > 0 {
		result = fetcp.Disconnect(r.socketHnd)
		if result < 0 {
			fmt.Errorf("unable to disconnect from the reader cleanly")
		}
		r.socketHnd = 0
	}

	return nil
}
