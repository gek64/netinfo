package preload

import (
	"encoding/json"
	"github.com/gek64/gek/gCrypto"
	"golang.org/x/crypto/chacha20poly1305"
	"netinfo/internal/netinfo"
	"time"
)

const AssociatedDataSize = 8

func newPreload(id string) (preload []byte, err error) {
	var preloadStrut netinfo.Data

	netInterfaces, err := netinfo.GetNetInterfaces()
	if err != nil {
		return nil, err
	}

	preloadStrut.ID = id
	preloadStrut.UpdatedAt = time.Now()
	preloadStrut.NetInterfaces = netInterfaces

	return json.Marshal(preloadStrut)
}

func newEncryptedPreload(id string, key []byte, associatedDataSize uint) (preload []byte, err error) {
	plaintext, err := newPreload(id)
	if err != nil {
		return nil, err
	}
	return gCrypto.NewChaCha20Poly1305(key, associatedDataSize).Encrypt(plaintext)
}

func GetPreload(id string, key []byte) (preload []byte, err error) {
	// 通过密钥长度判断是否使用加密
	switch len(key) {
	case 0:
		return newPreload(id)
	default:
		key = gCrypto.KeyZeroPadding(key, chacha20poly1305.KeySize)
		key = gCrypto.KeyCropping(key, chacha20poly1305.KeySize)
		return newEncryptedPreload(id, key, AssociatedDataSize)
	}
}
