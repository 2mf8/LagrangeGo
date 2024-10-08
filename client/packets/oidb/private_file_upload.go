package oidb

import (
	"github.com/2mf8/LagrangeGo/client/packets/pb/service/oidb"
	"github.com/2mf8/LagrangeGo/message"
	"github.com/2mf8/LagrangeGo/utils/crypto"
)

func BuildPrivateFileUploadReq(selfUid string, targetUid string, file *message.FileElement) (*OidbPacket, error) {
	md510MCheckSum, _ := crypto.ComputeMd5AndLengthWithLimit(file.FileStream, 10*1024*1024)
	body := &oidb.OidbSvcTrpcTcp0XE37_1700{
		Command: 1700,
		Seq:     0,
		Upload: &oidb.ApplyUploadReqV3{
			SenderUid:      selfUid,
			ReceiverUid:    targetUid,
			FileSize:       uint32(file.FileSize),
			FileName:       file.FileName,
			Md510MCheckSum: md510MCheckSum,
			Sha1CheckSum:   file.FileSha1,
			LocalPath:      "/",
			Md5CheckSum:    file.FileMd5,
			Sha3CheckSum:   []byte{},
		},
		BusinessId:               3,
		ClientType:               1,
		FlagSupportMediaPlatform: 1,
	}
	return BuildOidbPacket(0xE37, 1700, body, false, false)
}

func ParsePrivateFileUploadResp(data []byte) (*oidb.OidbSvcTrpcTcp0XE37Response, error) {
	return ParseTypedError[oidb.OidbSvcTrpcTcp0XE37Response](data)
}
