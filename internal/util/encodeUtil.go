package util

import (
	"encoding/base64"
	"reflect"
	"unsafe"
)

func Encode(data string) string {
	content := *(*[]byte)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&data))))
	coder := base64.NewEncoding(getBase64Table())
	return coder.EncodeToString(content)
}

func getBase64Table() string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	return str
}

func Decode(data string) string {
	coder := base64.NewEncoding(getBase64Table())
	result, _ := coder.DecodeString(data)
	return *(*string)(unsafe.Pointer(&result))
}
