// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/service/oidb/OidbSvcTrpcTcp0x6D8.proto

package oidb

// GroupFSView
type OidbSvcTrpcTcp0X6D8 struct {
	List  *OidbSvcTrpcTcp0X6D8List  `protobuf:"bytes,2,opt"`
	Count *OidbSvcTrpcTcp0X6D8Count `protobuf:"bytes,3,opt"`
	Space *OidbSvcTrpcTcp0X6D8Space `protobuf:"bytes,4,opt"`
	_     [0]func()
}

type OidbSvcTrpcTcp0X6D8List struct {
	GroupUin        uint32 `protobuf:"varint,1,opt"`
	AppId           uint32 `protobuf:"varint,2,opt"` // 7
	TargetDirectory string `protobuf:"bytes,3,opt"`
	FileCount       uint32 `protobuf:"varint,5,opt"`  // 20
	SortBy          uint32 `protobuf:"varint,9,opt"`  // 1
	StartIndex      uint32 `protobuf:"varint,13,opt"` // default 0
	Field17         uint32 `protobuf:"varint,17,opt"` // 2
	Field18         uint32 `protobuf:"varint,18,opt"` // 0
	_               [0]func()
}

type OidbSvcTrpcTcp0X6D8Count struct {
	GroupUin uint32 `protobuf:"varint,1,opt"`
	AppId    uint32 `protobuf:"varint,2,opt"` // 7
	BusId    uint32 `protobuf:"varint,3,opt"` // 6
	_        [0]func()
}

type OidbSvcTrpcTcp0X6D8Space struct {
	GroupUin uint32 `protobuf:"varint,1,opt"`
	AppId    uint32 `protobuf:"varint,2,opt"` // 7
	_        [0]func()
}

type OidbSvcTrpcTcp0X6D8_1Response struct {
	List  *OidbSvcTrpcTcp0X6D8_1ResponseList  `protobuf:"bytes,2,opt"`
	Count *OidbSvcTrpcTcp0X6D8_1ResponseCount `protobuf:"bytes,3,opt"`
	Space *OidbSvcTrpcTcp0X6D8_1ResponseSpace `protobuf:"bytes,4,opt"`
	_     [0]func()
}

type OidbSvcTrpcTcp0X6D8_1ResponseList struct {
	RetCode       int32                                `protobuf:"varint,1,opt"`
	RetMsg        string                               `protobuf:"bytes,2,opt"`
	ClientWording string                               `protobuf:"bytes,3,opt"`
	IsEnd         bool                                 `protobuf:"varint,4,opt"`
	Items         []*OidbSvcTrpcTcp0X6D8_1ResponseItem `protobuf:"bytes,5,rep"`
}

type OidbSvcTrpcTcp0X6D8_1ResponseCount struct {
	FileCount  uint32 `protobuf:"varint,4,opt"`
	LimitCount uint32 `protobuf:"varint,6,opt"`
	IsFull     bool   `protobuf:"varint,7,opt"`
	_          [0]func()
}

type OidbSvcTrpcTcp0X6D8_1ResponseSpace struct {
	TotalSpace uint64 `protobuf:"varint,4,opt"`
	UsedSpace  uint64 `protobuf:"varint,5,opt"`
	Field6     uint32 `protobuf:"varint,6,opt"`
	_          [0]func()
}

type OidbSvcTrpcTcp0X6D8_1ResponseItem struct {
	Type       uint32                                   `protobuf:"varint,1,opt"`
	FolderInfo *OidbSvcTrpcTcp0X6D8_1ResponseFolderInfo `protobuf:"bytes,2,opt"`
	FileInfo   *OidbSvcTrpcTcp0X6D8_1ResponseFileInfo   `protobuf:"bytes,3,opt"`
	_          [0]func()
}

type OidbSvcTrpcTcp0X6D8_1ResponseFolderInfo struct {
	FolderId          string `protobuf:"bytes,1,opt"`
	ParentDirectoryId string `protobuf:"bytes,2,opt"`
	FolderName        string `protobuf:"bytes,3,opt"`
	CreateTime        uint32 `protobuf:"varint,4,opt"`
	ModifiedTime      uint32 `protobuf:"varint,5,opt"`
	CreatorUin        uint32 `protobuf:"varint,6,opt"`
	CreatorName       string `protobuf:"bytes,7,opt"`
	TotalFileCount    uint32 `protobuf:"varint,8,opt"`
	_                 [0]func()
}

type OidbSvcTrpcTcp0X6D8_1ResponseFileInfo struct {
	FileId          string `protobuf:"bytes,1,opt"`
	FileName        string `protobuf:"bytes,2,opt"`
	FileSize        uint64 `protobuf:"varint,3,opt"`
	BusId           uint32 `protobuf:"varint,4,opt"`
	UploadedSize    uint64 `protobuf:"varint,5,opt"`
	UploadedTime    uint32 `protobuf:"varint,6,opt"`
	ExpireTime      uint32 `protobuf:"varint,7,opt"`
	ModifiedTime    uint32 `protobuf:"varint,8,opt"`
	DownloadedTimes uint32 `protobuf:"varint,9,opt"`
	FileSha1        []byte `protobuf:"bytes,10,opt"`
	FileMd5         []byte `protobuf:"bytes,12,opt"`
	UploaderName    string `protobuf:"bytes,14,opt"`
	UploaderUin     uint32 `protobuf:"varint,15,opt"`
	ParentDirectory string `protobuf:"bytes,16,opt"`
	Field17         uint32 `protobuf:"varint,17,opt"`
	Field22         string `protobuf:"bytes,22,opt"`
}
