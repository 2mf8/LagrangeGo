package oidb

import (
	"github.com/2mf8/LagrangeGo/client/packets/pb/service/oidb"
)

func BuildGroupLeaveReq(groupUin uint32) (*OidbPacket, error) {
	body := &oidb.OidbSvcTrpcTcp0X1097_1{GroupUin: groupUin}
	return BuildOidbPacket(0x1097, 1, body, false, false)
}

func ParseGroupLeaveResp(data []byte) error {
	return CheckError(data)
}
