package request

import (
	"qiyuesuo/sdk/http"
)

type SaasPlatformFeeRequest struct {
}

func (obj SaasPlatformFeeRequest) GetUrl() string {
	return "/saas/v2/fee/platform/info"
}

func (obj SaasPlatformFeeRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	return parameter
}
