package request

import (
	"qiyuesuo/sdk/http"
)

type CompanyFeeQueryRequest struct {
	CompanyName string `json:"companyName,omitempty"`
}

func (obj CompanyFeeQueryRequest) GetUrl() string {
	return "/v2/company/fee/info"
}

func (obj CompanyFeeQueryRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("companyName", obj.CompanyName)
	return parameter
}
