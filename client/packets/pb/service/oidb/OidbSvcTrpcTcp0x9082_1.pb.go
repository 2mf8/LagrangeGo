// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/service/oidb/OidbSvcTrpcTcp0x9082_1.proto

package oidb

import (
	proto "github.com/RomiChan/protobuf/proto"
)

// Group Set Reaction
type OidbSvcTrpcTcp0X9082_1 struct {
	GroupUin uint32               `protobuf:"varint,2,opt"`
	Sequence uint32               `protobuf:"varint,3,opt"`
	Code     proto.Option[string] `protobuf:"bytes,4,opt"`
	Field5   bool                 `protobuf:"varint,5,opt"`
	Field6   bool                 `protobuf:"varint,6,opt"`
	Field7   bool                 `protobuf:"varint,7,opt"`
	_        [0]func()
}