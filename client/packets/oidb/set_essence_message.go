package oidb

import (
	"github.com/2mf8/LagrangeGo/client/packets/pb/service/oidb"
	"github.com/2mf8/LagrangeGo/utils"
)

func BuildSetEssenceMessageReq(groupUin, seq, random uint32, isSet bool) (*OidbPacket, error) {
	body := oidb.OidbSvcTrpcTcp0XEAC{
		GroupUin: groupUin,
		Sequence: seq,
		Random:   random,
	}
	return BuildOidbPacket(0xEAC, utils.Ternary[uint32](isSet, 1, 2), &body, false, false)
}

func ParseSetEssenceMessageResp(data []byte) error {
	return CheckError(data)
}
