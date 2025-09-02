package base64image

import "encoding/base64"

// base64String base64编码的图片字符串，不包含 data:image/png;base64, 前缀
func GetImageByBase64(base64String string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(base64String)
}
