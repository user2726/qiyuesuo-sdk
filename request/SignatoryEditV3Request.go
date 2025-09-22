package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type SignatoryEditV3Request struct {
	// 合同ID与业务ID二选一，不能同时为空
	SignatoryId string `json:"signatoryId,omitempty"`
	// 合同ID与业务ID二选一，不能同时为空
	SignValidateWay string `json:"signValidateWay,omitempty"`
	// 若使用业务ID发起合同，且合同是以子公司身份创建的，则需要传递该值，用于确定合同主体
	TenantName string      `json:"tenantName,omitempty"`
	Receiver   *model.User `json:"receiver,omitempty"`
	Operator   *model.User `json:"operator,omitempty"`
}

func (obj SignatoryEditV3Request) GetUrl() string {
	return "/v3/signatory/edit"
}

func (obj SignatoryEditV3Request) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
