package client

// from https://github.com/Mrs4s/MiraiGo/blob/master/client/entities.go

import (
	"errors"

	"github.com/2mf8/LagrangeGo/message"
)

var (
	ErrAlreadyOnline  = errors.New("already online")
	ErrMemberNotFound = errors.New("member not found")
	ErrNotExists      = errors.New("not exists")
)

type ClientDisconnectedEvent struct {
	Message string
}

type groupMessageReceiptEvent struct {
	Rand int32
	Seq  int32
	Msg  *message.GroupMessage
}
