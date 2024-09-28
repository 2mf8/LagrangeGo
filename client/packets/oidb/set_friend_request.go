package oidb

import (
	"github.com/2mf8/LagrangeGo/client/packets/pb/service/oidb"
	"github.com/2mf8/LagrangeGo/utils"
)

func BuildSetFriendRequest(accept bool, targetUid string) (*OidbPacket, error) {
	body := oidb.OidbSvcTrpcTcp0XB5D_44{
		Accept:    utils.Ternary[uint32](accept, 3, 5),
		TargetUid: targetUid,
	}
	return BuildOidbPacket(0xb5d, 44, &body, false, false)
}

func ParseSetFriendRequestResp(data []byte) error {

	return CheckError(data)
}
