package util

import "fmt"

//
//Marshal a pictite into a CQ code
//"file" support the following
//file: file:///C:\\Users\Alice\Pictures\1.png
//url https://www.baidu.com/img/xxxxxx.png
//base64 base64://your-base64
//Example ;[CQ:image,file=http://baidu.com/1.jpg,type=show,id=40004]
//cache:0/1
//
func MarshalImage(file string) string {
	return fmt.Sprintf("[CQ:image,file=%s]", file)
}
