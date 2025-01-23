package handler

import "mime/multipart"

type PublishReq struct {
	Title string                `json:"title"`
	Data  *multipart.FileHeader `json:"data"`
}
