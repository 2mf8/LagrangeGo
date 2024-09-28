package event

import (
	"github.com/2mf8/LagrangeGo/client/auth"
	"github.com/2mf8/LagrangeGo/client/packets/pb/message"
)

func ParseSelfRenameEvent(event *message.SelfRenameMsg, sig *auth.SigInfo) *Rename {
	sig.Nickname = event.Body.RenameData.NickName
	return &Rename{
		Uin:      event.Body.Uin,
		Nickname: event.Body.RenameData.NickName,
	}
}
