package bytesimage

import (
	"io"
	"net/http"
	"os"
	"time"
)

func GetImageBytes(path string) ([]byte, error) {
	if StartsWith(path, "http") {
		cli := &http.Client{
			Timeout: time.Second * 5,
		}
		resp, err := cli.Get(path)
		if err != nil {
			return nil, err
		}
		v, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		mimeType := http.DetectContentType(v)
		switch mimeType {
		case "image/jpeg":
			//baseImage := "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(v)
			return v, nil
		case "image/png":
			//baseImage := "data:image/png;base64," + base64.StdEncoding.EncodeToString(v)
			return v, nil
		default:
			return nil, nil
		}
	} else {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		v, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}
		mimeType := http.DetectContentType(v)
		switch mimeType {
		case "image/jpeg":
			//baseImage := "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(v)
			return v, nil
		case "image/png":
			//baseImage := "data:image/png;base64," + base64.StdEncoding.EncodeToString(v)
			return v, nil
		default:
			return nil, nil
		}
	}
}

func StartsWith(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
