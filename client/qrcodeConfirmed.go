package client

import (
	"github.com/LagrangeDev/LagrangeGo/packets/tlv"
	"github.com/LagrangeDev/LagrangeGo/packets/wtlogin"
	"github.com/LagrangeDev/LagrangeGo/utils"
	"github.com/LagrangeDev/LagrangeGo/utils/binary"
)

func (c *QQClient) QrcodeConfirmed() error {
	app := c.appInfo
	device := c.deviceInfo
	body := binary.NewBuilder(nil).
		WriteU16(0x09).
		WriteTlv(
			binary.NewBuilder(nil).WriteBytes(c.t106).Pack(0x106),
			tlv.T144(c.sig.Tgtgt, app, device),
			tlv.T116(app.SubSigmap),
			tlv.T142(app.PackageName, 0),
			tlv.T145(utils.MustParseHexStr(device.Guid)),
			tlv.T18(0, app.AppClientVersion, int(c.Uin), 0, 5, 0),
			tlv.T141([]byte("Unknown"), nil),
			tlv.T177(app.WTLoginSDK, 0),
			tlv.T191(0),
			tlv.T100(5, app.AppID, app.SubAppID, 8001, app.MainSigmap, 0),
			tlv.T107(1, 0x0d, 0, 1),
			tlv.T318(nil),
			binary.NewBuilder(nil).WriteBytes(c.t16a).Pack(0x16a),
			tlv.T166(5),
			tlv.T521(0x13, "basicim"),
		).ToBytes()

	response, err := c.SendUniPacketAndAwait(
		"wtlogin.login",
		wtlogin.BuildLoginPacket(c.Uin, "wtlogin.login", app, body))

	if err != nil {
		networkLogger.Error(err)
		return err
	}

	return wtlogin.DecodeLoginResponse(response.Data, c.sig)
}
