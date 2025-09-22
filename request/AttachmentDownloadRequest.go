package request

import (
	"qiyuesuo/sdk/http"
)

type AttachmentDownloadRequest struct {
	// 合同ID，合同ID与业务ID二选一，不能同时为空
	ContractId string `json:"contractId,omitempty"`
	// 业务ID，合同ID与业务ID二选一，不能同时为空
	BizId string `json:"bizId,omitempty"`
	// 子公司名称，若使用业务ID下载合同，且合同是以子公司身份创建的，则需要传递该值，用于确定合同主体
	TenantName string `json:"tenantName,omitempty"`
	Origin     *bool  `json:"origin,omitempty"`
}

func (obj AttachmentDownloadRequest) GetUrl() string {
	return "/v2/attachment/download"
}

func (obj AttachmentDownloadRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("contractId", obj.ContractId)
	parameter.AddParam("bizId", obj.BizId)
	parameter.AddParam("tenantName", obj.TenantName)
	parameter.AddParam("origin", obj.Origin)
	return parameter
}
