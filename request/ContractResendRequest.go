package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type ContractResendRequest struct {
	// 接口返回值
	ResendContract *model.ResendContract
}

func (obj ContractResendRequest) GetUrl() string {
	return "/v2/contract/resend"
}

func (obj ContractResendRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj.ResendContract)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
