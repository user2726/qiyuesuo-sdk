package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
)

type SaaSFeeUpdateRequest struct {
	CompanyName    string `json:"companyName,omitempty"`
	InitialBalance string `json:"initialBalance,omitempty"`
	InitialPrice   string `json:"initialPrice,omitempty"`
	PayMode        string `json:"payMode,omitempty"`
}

func (obj SaaSFeeUpdateRequest) GetUrl() string {
	return "/saas/v2/fee/update"
}

func (obj SaaSFeeUpdateRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
