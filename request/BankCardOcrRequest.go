package request

import (
	"qiyuesuo/sdk/http"
)

type BankCardOcrRequest struct {
	Image    *http.FileItem `json:"image,omitempty"`
	ImageUrl string         `json:"imageUrl,omitempty"`
}

func (obj BankCardOcrRequest) GetUrl() string {
	return "/v2/ocr/bankcard"
}

func (obj BankCardOcrRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	parameter.AddFiles("image", obj.Image)
	parameter.AddParam("imageUrl", obj.ImageUrl)
	return parameter
}
