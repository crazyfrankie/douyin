package handler

type MessageActionReq struct {
	UserId     int64  `json:"user_id"`
	ToUserId   int64  `json:"to_user_id"`
	Content    string `json:"content"`
	ActionType int64  `json:"action_type"`
}

type MessageChatReq struct {
	UserId     int64 `json:"user_id"`
	ToUserId   int64 `json:"to_user_id"`
	PreMsgTime int64 `json:"pre_msg_time"`
}
