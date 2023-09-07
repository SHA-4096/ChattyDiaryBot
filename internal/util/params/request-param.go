package util

type SendMessageParam struct {
	Message_type string `json:"message_type"` //消息类型, 支持 private、group , 分别对应私聊、群组, 如不传入, 则根据传入的 *_id 参数判断
	User_id      int    `json:"user_id"`      //对方 QQ 号 ( 消息类型为 private 时需要 )
	Group_id     int    `json:"group_id"`     //群号 ( 消息类型为 group 时需要 )
	Message      string `json:"message"`      //要发送的内容
	Auto_escape  bool   `json:"auto_escape"`  //消息内容是否作为纯文本发送 ( 即不解析 CQ 码 ) , 只在 message 字段是字符串时有效
}
