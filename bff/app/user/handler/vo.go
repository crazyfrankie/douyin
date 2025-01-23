package handler

type LoginReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RegisterReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
