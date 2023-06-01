package file

import (
	"errors"
	"net/http"
	"os"
)

func GetFileType(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", errors.New("不可识别的文件")
	}
	defer file.Close()
	buffer := make([]byte, 512)
	n, err := file.Read(buffer)
	if err != nil {
		return "", errors.New("不可识别的文件")
	}
	contentType := http.DetectContentType(buffer[:n])
	return contentType, nil
}
