// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/service/oidb/OidbSvcTrpcTcp0xFE7_3.proto

package oidb

import (
	proto "github.com/RomiChan/protobuf/proto"
)

// Fetch Group Member List
type OidbSvcTrpcTcp0XFE7_3 struct {
	GroupUin uint32                     `protobuf:"varint,1,opt"`
	Field2   uint32                     `protobuf:"varint,2,opt"`
	Field3   uint32                     `protobuf:"varint,3,opt"`
	Body     *OidbSvcTrpcScp0XFE7_3Body `protobuf:"bytes,4,opt"`
	Token    proto.Option[string]       `protobuf:"bytes,15,opt"`
	_        [0]func()
}

type OidbSvcTrpcScp0XFE7_3Body struct {
	// all ture
	MemberName       bool `protobuf:"varint,10,opt"`
	MemberCard       bool `protobuf:"varint,11,opt"`
	Level            bool `protobuf:"varint,12,opt"`
	SpecialTitle     bool `protobuf:"varint,17,opt"`
	Field4           bool `protobuf:"varint,20,opt"`
	Field5           bool `protobuf:"varint,21,opt"`
	JoinTimestamp    bool `protobuf:"varint,100,opt"`
	LastMsgTimestamp bool `protobuf:"varint,101,opt"`
	ShutUpTimestamp  bool `protobuf:"varint,102,opt"`
	Field9           bool `protobuf:"varint,103,opt"`
	Field10          bool `protobuf:"varint,104,opt"`
	Field11          bool `protobuf:"varint,105,opt"`
	Field12          bool `protobuf:"varint,106,opt"`
	Permission       bool `protobuf:"varint,107,opt"`
	_                [0]func()
}

type OidbSvcTrpcTcp0XFE7_2Response struct {
	GroupUin            uint32                         `protobuf:"varint,1,opt"`
	Members             []*OidbSvcTrpcTcp0XFE7_2Member `protobuf:"bytes,2,rep"`
	Field3              uint32                         `protobuf:"varint,3,opt"`
	MemberChangeSeq     uint32                         `protobuf:"varint,5,opt"`
	MemberCardChangeSeq uint32                         `protobuf:"varint,6,opt"`
	Token               proto.Option[string]           `protobuf:"bytes,15,opt"` // for the next page
}

type OidbSvcTrpcTcp0XFE7_2Member struct {
	Uin              *OidbSvcTrpcTcp0XFE7_2Uin   `protobuf:"bytes,1,opt"`
	MemberName       string                      `protobuf:"bytes,10,opt"`
	MemberCard       *OidbSvcTrpcTcp0XFE7_2Card  `protobuf:"bytes,11,opt"`
	Level            *OidbSvcTrpcTcp0XFE7_2Level `protobuf:"bytes,12,opt"`
	SpecialTitle     proto.Option[string]        `protobuf:"bytes,17,opt"`
	JoinTimestamp    uint32                      `protobuf:"varint,100,opt"`
	LastMsgTimestamp uint32                      `protobuf:"varint,101,opt"`
	ShutUpTimestamp  proto.Option[uint32]        `protobuf:"varint,102,opt"`
	Permission       uint32                      `protobuf:"varint,107,opt"`
	_                [0]func()
}

type OidbSvcTrpcTcp0XFE7_2Uin struct {
	Uid string `protobuf:"bytes,2,opt"`
	Uin uint32 `protobuf:"varint,4,opt"`
	_   [0]func()
}

type OidbSvcTrpcTcp0XFE7_2Card struct {
	MemberCard proto.Option[string] `protobuf:"bytes,2,opt"`
	_          [0]func()
}

type OidbSvcTrpcTcp0XFE7_2Level struct {
	Infos []uint32 `protobuf:"varint,1,rep"`
	Level uint32   `protobuf:"varint,2,opt"`
}

type OidbSvcTrpcTcp0XFE7_4 struct {
	GroupUin uint32                       `protobuf:"varint,1,opt"`
	Field2   uint32                       `protobuf:"varint,2,opt"`
	Field3   uint32                       `protobuf:"varint,3,opt"`
	Body     *OidbSvcTrpcScp0XFE7_3Body   `protobuf:"bytes,4,opt"`
	Params   *OidbSvcTrpcScp0XFE7_4Params `protobuf:"bytes,5,opt"`
	_        [0]func()
}

type OidbSvcTrpcScp0XFE7_4Params struct {
	Uid string `protobuf:"bytes,2,opt"`
	_   [0]func()
}

type OidbSvcTrpcTcp0XFE7_4Response struct {
	GroupUin uint32                       `protobuf:"varint,1,opt"`
	Member   *OidbSvcTrpcTcp0XFE7_2Member `protobuf:"bytes,2,opt"`
	Field3   uint32                       `protobuf:"varint,3,opt"`
	_        [0]func()
}
