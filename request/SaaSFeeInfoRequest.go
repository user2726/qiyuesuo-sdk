package request

import (
	"qiyuesuo/sdk/http"
)

type SaaSFeeInfoRequest struct {
	CompanyName string `json:"companyName,omitempty"`
}

func (obj SaaSFeeInfoRequest) GetUrl() string {
	return "/saas/v2/fee/info"
}

func (obj SaaSFeeInfoRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("companyName", obj.CompanyName)
	return parameter
}
