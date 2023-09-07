package controller

type MessagePrivate struct {
	Time         int    `json:"time"`         //事件发生的时间戳
	Self_id      int    `json:"self_id"`      //收到事件的机器人 QQ 号
	Post_type    string `json:"post_type"`    //参考	message	上报类型
	Message_type string `json:"message_type"` //消息类型
	Sub_type     string `json:"sub_type"`     //参考	friend、group、group_self、other	消息子类型, 如果是好友则是 friend, 如果是群临时会话则是 group, 如果是在群中自身发送则是 group_self
	Message_id   int    `json:"message_id"`   //消息 ID
	User_id      int    `json:"user_id"`      //发送者 QQ 号
	Message      string `json:"message"`      //消息内容
	Raw_message  string `json:"raw_message"`  //原始消息内容
	Font         int    `json:"font"`         //字体
	//Sender       object  `json:"sender"`       //发送人信息
	Target_id   int `json:"target_id"`   //接收者 QQ 号
	Temp_source int `json:"temp_source"` //临时会话来源
}

type Message struct {
	MsgType string `json:"type"`
	Data    Data   `json:"data"`
}

type Data struct {
	Text string `json:"text"`
	QQ   string `json:"qq"`
}
