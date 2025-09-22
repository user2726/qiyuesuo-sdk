package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
)

type SaaSFeeInitRequest struct {
	CompanyName    string `json:"companyName,omitempty"`
	InitialBalance string `json:"initialBalance,omitempty"`
	InitialPrice   string `json:"initialPrice,omitempty"`
	PayMode        string `json:"payMode,omitempty"`
}

func (obj SaaSFeeInitRequest) GetUrl() string {
	return "/saas/v2/fee/init"
}

func (obj SaaSFeeInitRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
