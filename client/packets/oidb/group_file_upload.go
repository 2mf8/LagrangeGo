package oidb

import (
	"errors"

	"github.com/LagrangeDev/LagrangeGo/client/packets/pb/service/oidb"
	"github.com/LagrangeDev/LagrangeGo/message"
)

func BuildGroupFileUploadReq(groupUin uint32, file *message.FileElement, targetDirectory string) (*Packet, error) {
	body := &oidb.OidbSvcTrpcTcp0X6D6{
		File: &oidb.OidbSvcTrpcTcp0X6D6Upload{
			GroupUin:        groupUin,
			AppId:           4,
			BusId:           102,
			Entrance:        6,
			TargetDirectory: targetDirectory,
			FileName:        file.FileName,
			LocalDirectory:  "/" + file.FileName,
			FileSize:        file.FileSize,
			FileSha1:        file.FileSha1,
			FileSha3:        []byte{},
			FileMd5:         file.FileMd5,
			Field15:         true,
		},
	}
	return BuildOidbPacket(0x6D6, 0, body, false, true)
}

func ParseGroupFileUploadResp(data []byte) (*oidb.OidbSvcTrpcTcp0X6D6Response, error) {
	var resp oidb.OidbSvcTrpcTcp0X6D6Response
	if _, err := ParseOidbPacket(data, &resp); err != nil {
		return nil, err
	}
	if resp.Upload.RetCode != 0 {
		return nil, errors.New(resp.Upload.ClientWording)
	}
	return &resp, nil
}
