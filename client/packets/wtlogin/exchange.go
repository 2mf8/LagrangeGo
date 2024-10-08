package wtlogin

import (
	"encoding/hex"

	"github.com/2mf8/LagrangeGo/client/packets/pb/login"
	"github.com/2mf8/LagrangeGo/internal/proto"
	"github.com/2mf8/LagrangeGo/utils"
	"github.com/2mf8/LagrangeGo/utils/binary"
	"github.com/2mf8/LagrangeGo/utils/crypto"
	"github.com/2mf8/LagrangeGo/utils/crypto/ecdh"
)

var encKey, _ = hex.DecodeString("e2733bf403149913cbf80c7a95168bd4ca6935ee53cd39764beebe2e007e3aee")

func BuildKexExchangeRequest(uin uint32, guid string) ([]byte, error) {
	encl, err := crypto.AESGCMEncrypt(proto.DynamicMessage{
		1: uin,
		2: guid,
	}.Encode(), ecdh.P256().SharedKey())
	if err != nil {
		return nil, err
	}

	p2Hash := crypto.SHA256Digest(
		binary.NewBuilder(nil).
			WriteBytes(ecdh.P256().PublicKey()).
			WriteU32(1).
			WriteBytes(encl).
			WriteU32(0).
			WriteU32(uint32(utils.TimeStamp())).
			ToBytes(),
	)
	encP2Hash, err := crypto.AESGCMEncrypt(p2Hash, encKey)
	if err != nil {
		return nil, err
	}

	return proto.DynamicMessage{
		1: ecdh.P256().PublicKey(),
		2: 1,
		3: encl,
		4: utils.TimeStamp(),
		5: encP2Hash,
	}.Encode(), nil
}

func ParseKeyExchangeResponse(response []byte) (key, sign []byte, err error) {
	var p login.SsoKeyExchangeResponse
	err = proto.Unmarshal(response, &p)
	if err != nil {
		return
	}

	shareKey, err := ecdh.P256().Exange(p.PublicKey)
	if err != nil {
		return
	}

	var decPb login.SsoKeyExchangeDecrypted
	data, err := crypto.AESGCMDecrypt(p.GcmEncrypted, shareKey)
	if err != nil {
		return
	}
	err = proto.Unmarshal(data, &decPb)
	if err != nil {
		return
	}

	return decPb.GcmKey, decPb.Sign, nil
}
