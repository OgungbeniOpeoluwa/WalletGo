package util

import (
	"encoding/base64"
)

func Encode(api string) string {
	encoder := base64.StdEncoding.EncodeToString([]byte(api))
	return encoder
}
