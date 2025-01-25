package handler

type CommentReq struct {
	VideoID     int64
	ActionType  int64
	CommentText string
	CommentID   int64
}
