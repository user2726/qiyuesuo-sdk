package request

import (
	"qiyuesuo/sdk/http"
)

type IdCardOcrRequest struct {
	Image    *http.FileItem `json:"image,omitempty"`
	ImageUrl string         `json:"imageUrl,omitempty"`
}

func (obj IdCardOcrRequest) GetUrl() string {
	return "/v2/ocr/idcard"
}

func (obj IdCardOcrRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	parameter.AddFiles("image", obj.Image)
	parameter.AddParam("imageUrl", obj.ImageUrl)
	return parameter
}
