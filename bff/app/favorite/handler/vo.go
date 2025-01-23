package handler

type FavoriteActionReq struct {
	VideoId    int64 `json:"video_id"`
	ActionType int32 `json:"action_type"`
}
