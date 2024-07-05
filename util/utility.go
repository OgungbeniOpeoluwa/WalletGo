package util

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateCryptoID() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

//func HashMonifyWebhookRequest(req request.MonnifyWebhookRequest) {
//	hash := sha512.New()
//	//hash.Write([]byte())
//}
