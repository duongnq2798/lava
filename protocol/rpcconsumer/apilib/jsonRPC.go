package apilib

import (
	"fmt"

	spectypes "github.com/lavanet/lava/x/spec/types"
)

type JsonRPCAPIParser struct{}

func (apip *JsonRPCAPIParser) ParseMsg(url string, data []byte, connectionType string) (APIMessage, error) {
	return nil, nil
}

func (apip *JsonRPCAPIParser) SetSpec(spec spectypes.Spec) {}

func NewJrpcAPIParser() (apiParser *JsonRPCAPIParser, err error) {
	return nil, fmt.Errorf("not implemented")
}

type JsonRPCAPIListener struct{}

func (apil *JsonRPCAPIListener) Serve() {}

func NewJrpcAPIListener() (apiListener *JsonRPCAPIListener) {
	return nil
}