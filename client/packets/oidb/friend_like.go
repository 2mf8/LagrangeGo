package oidb

import (
	"github.com/2mf8/LagrangeGo/client/packets/pb/service/oidb"
	"github.com/2mf8/LagrangeGo/internal/proto"
)

func BuildFriendLikeReq(uid string, count uint32) (*OidbPacket, error) {
	body := &oidb.OidbSvcTrpcTcp0X7E5_104{
		TargetUid: proto.Some(uid),
		Field2:    71,
		Field3:    count,
	}
	return BuildOidbPacket(0x7E5, 104, body, false, false)
}

func ParseFriendLikeResp(data []byte) error {
	return CheckError(data)
}
