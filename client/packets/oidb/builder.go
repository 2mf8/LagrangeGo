package oidb

import (
	"errors"
	"fmt"

	"github.com/2mf8/LagrangeGo/client/packets/pb/service/oidb"
	"github.com/2mf8/LagrangeGo/internal/proto"
	"github.com/2mf8/LagrangeGo/utils"
)

// var oidbLogger = utils.GetLogger("oidb")

type OidbPacket struct {
	Cmd       string
	Data      []byte
	ExtraData []byte
}

func BuildOidbPacket(cmd, subCmd uint32, body any, isLafter, isUid bool) (*OidbPacket, error) {
	bodyData, err := proto.Marshal(body)
	if err != nil {
		return nil, err
	}
	oidbPkt := &oidb.OidbSvcTrpcTcpBase{
		Command:    cmd,
		SubCommand: subCmd,
		Body:       bodyData,
		Reserved:   int32(utils.Bool2Int(isUid)),
	}
	if isLafter {
		oidbPkt.Lafter = &oidb.OidbLafter{}
	}

	data, err := proto.Marshal(oidbPkt)
	if err != nil {
		return nil, err
	}

	return &OidbPacket{
		Cmd:  fmt.Sprintf("OidbSvcTrpcTcp.0x%02x_%d", cmd, subCmd),
		Data: data,
	}, nil
}

func ParseOidbPacket(b []byte, pkt any) (oidbBaseResp oidb.OidbSvcTrpcTcpBase, err error) {
	err = proto.Unmarshal(b, &oidbBaseResp)
	if err != nil {
		return
	}
	if oidbBaseResp.ErrorCode != 0 {
		return oidbBaseResp, errors.New(oidbBaseResp.ErrorMsg)
	}
	if pkt == nil {
		return
	}
	err = proto.Unmarshal(oidbBaseResp.Body, pkt)
	return
}

func CheckError(data []byte) error {
	baseResp, err := ParseOidbPacket(data, nil)
	if err != nil {
		return err
	}
	if baseResp.ErrorCode != 0 {
		return errors.New(baseResp.ErrorMsg)
	}
	return nil
}

func CheckTypedError[T any](data []byte) error {
	var resp T
	baseResp, err := ParseOidbPacket(data, &resp)
	if err != nil {
		return err
	}
	if baseResp.ErrorCode != 0 {
		return errors.New(baseResp.ErrorMsg)
	}
	return nil
}

func ParseTypedError[T any](data []byte) (*T, error) {
	var resp T
	baseResp, err := ParseOidbPacket(data, &resp)
	if err != nil {
		return nil, err
	}
	if baseResp.ErrorCode != 0 {
		return nil, errors.New(baseResp.ErrorMsg)
	}
	return &resp, nil
}
