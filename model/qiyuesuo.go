package model

type CreateContractDraftReq struct {
	ContractSubject string `json:"contractSubject"`
	Receiver string `json:"receiver"`
	ReceiverContact string `json:"receiverContact"`
	
}

type AddContractDocumentsUsingFilesReq struct {
	ContractId string                                        `json:"contractId"`
	ProjectName string `json:"projectName"`
	Id int `json:"id"`
	Documents  []*AddContractDocumentsUsingFilesReqDocuments `json:"documents"`
}

type AddContractDocumentsUsingFilesReqDocuments struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type InitiateContractReq struct {
	ContractId string `json:"contractId"`
}
type SignaturePageReq struct {
	ContractId string `json:"contractId"`
	Contact    string `json:"contact"`
}
type BrowsePageReq struct {
	ContractId string `json:"contractId"`
}
type ContractDetailsReq struct {
	ContractId string `json:"contractId"`
}
// QiYueSuoCreateContractResponse 契约锁创建合同的响应结构体
type QiYueSuoCreateContractResponse struct {
	Code        string `json:"code"`        // 响应码，"0"表示成功
	Message     string `json:"message"`     // 响应消息
	ResponseCode string `json:"responseCode"` // 响应业务码，"00000000"表示成功
	Result      struct {
		Category struct {
			ID   string `json:"id"`   // 分类ID
			Name string `json:"name"` // 分类名称
		} `json:"category"`
		Creator struct {
			Name string `json:"name"` // 创建者名称
		} `json:"creator"`
		ExpireTime string `json:"expireTime"` // 过期时间
		ID         string `json:"id"`         // 合同ID
		Ordinal    bool   `json:"ordinal"`    // 是否有序
		Signatories []struct {
			DelaySet   bool   `json:"delaySet"`   // 是否延迟签署
			ID         string `json:"id"`         // 签署方ID
			Receiver   struct {
				Contact     string `json:"contact"`     // 联系方式
				ContactType string `json:"contactType"` // 联系方式类型
			} `json:"receiver"`
			SerialNo  int    `json:"serialNo"`  // 签署顺序号
			Sponsor   bool   `json:"sponsor"`   // 是否发起方
			Status    string `json:"status"`    // 签署状态
			TenantName  string `json:"tenantName"`  // 租户名称
			TenantType string `json:"tenantType"` // 租户类型
		} `json:"signatories"`
		SN          string `json:"sn"`           // 合同编号
		Status      string `json:"status"`       // 合同状态
		Subject     string `json:"subject"`      // 合同标题
		TenantName  string `json:"tenantName"`   // 租户名称
	} `json:"result"`
}

// QiYueAddFile 契约锁文件上传
type QiYueAddFile struct {
	Code        string `json:"code"`        // 响应码，"0"表示成功
	Message     string `json:"message"`     // 响应消息
	ResponseCode string `json:"responseCode"` // 响应业务码，"00000000"表示成功
	Result      struct {
		DocumentId string `json:"documentId"`
	} `json:"result"`
}

// QiYueSuoContractPageResponse 合同详情
type QiYueSuoContractPageResponse struct {
	Code        string `json:"code"`        // 响应码，"0"表示成功
	Message     string `json:"message"`     // 响应消息
	ResponseCode string `json:"responseCode"` // 响应业务码，"00000000"表示成功
	Result      struct{
		PageUrl string `json:"pageUrl"`
		Status string  `json:"status"`
	} `json:"result"`
}


// 以下是更详细的拆分结构体，便于单独使用

// CategoryInfo 分类信息
type CategoryInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CreatorInfo 创建者信息
type CreatorInfo struct {
	Name string `json:"name"`
}

// ReceiverInfo 接收者信息
type ReceiverInfo struct {
	Contact     string `json:"contact"`
	ContactType string `json:"contactType"`
}

// SignatoryInfo 签署方信息
type SignatoryInfo struct {
	DelaySet   bool        `json:"delaySet"`
	ID         string      `json:"id"`
	Receiver   ReceiverInfo `json:"receiver"`
	SerialNo   int         `json:"serialNo"`
	Sponsor    bool        `json:"sponsor"`
	Status     string      `json:"status"`
	TenantName string      `json:"tenantName"`
	TenantType string      `json:"tenantType"`
}

// ContractResult 合同创建结果
type ContractResult struct {
	Category    CategoryInfo    `json:"category"`
	Creator     CreatorInfo     `json:"creator"`
	ExpireTime  string          `json:"expireTime"`
	ID          string          `json:"id"`
	Ordinal     bool            `json:"ordinal"`
	Signatories []SignatoryInfo `json:"signatories"`
	SN          string          `json:"sn"`
	Status      string          `json:"status"`
	Subject     string          `json:"subject"`
	TenantName  string          `json:"tenantName"`
}

// QiYueSuoResponse 统一的契约锁响应结构
type QiYueSuoResponse struct {
	Code        string         `json:"code"`
	Message     string         `json:"message"`
	ResponseCode string         `json:"responseCode"`
	Result      ContractResult `json:"result"`
}

// WithdrawReportReq 撤回报告请求参数
type WithdrawReportReq struct {
	ContractId string `json:"contractId"`
	Contact    string `json:"contact"`
	Reason     string `json:"reason"` 
}

// ContractResponse 合同详情响应结构体
type ContractDetailsResponse struct {
	ResponseCode string   `json:"responseCode"` // 响应码
	Message      string   `json:"message"`      // 响应消息
	Result       ContractDetails `json:"result"`       // 合同详情
}
// Contract 合同详情结构体
type ContractDetails struct {
	ID                  string      `json:"id"`                  // 合同ID
	BizId               string      `json:"bizId"`               // 业务ID
	TenantName          string      `json:"tenantName"`          // 发起方名称
	Subject             string      `json:"subject"`             // 合同主题（合同名称）
	Description         string      `json:"description"`         // 合同描述
	SN                  string      `json:"sn"`                  // 合同编号
	Ordinal             bool        `json:"ordinal"`             // 是否顺序签署
	Category            Category    `json:"category"`            // 业务分类
	Creator             User        `json:"creator"`             // 创建人
	Status              string      `json:"status"`              // 合同状态
	Signatories         []Signatory `json:"signatories"`         // 签署方
	Documents           []Document  `json:"documents"`           // 合同文档
	Comments            string      `json:"comments"`            // 合同已撤回/已退回状态时的原因说明
	InvalidReason       string      `json:"invalidReason"`       // 合同作废原因
	ExpireTime          string      `json:"expireTime"`          // 截止签署时间
	PublishTime         string      `json:"publishTime"`         // 合同发起时间
	EndTime             string      `json:"endTime"`             // 文件的到期时间
	Tags                []string    `json:"tags"`                // 标签列表
	SignFlowStrategy    string      `json:"signFlowStrategy"`    // 签署策略
	RelatedContractIds  []string    `json:"relatedContractIds"`  // 关联合同ID
	CustomFields        []CustomField `json:"customFields"`       // 合同自定义参数
	CopySendReceivers   []Receiver   `json:"copySendReceivers"`   // 合同抄送人
	BusinessData        string       `json:"businessData"`        // 自定义业务参数
	InvalidExpireTime   string       `json:"invalidExpireTime"`   // 作废申请的截止日期
	InvalidFailType     string       `json:"invalidFailType"`     // 作废失败的原因
}

// CustomField 自定义参数结构体
type CustomField struct {
	Key   string `json:"key"`   // 自定义参数键
	Value string `json:"value"` // 自定义参数值
}
// Receiver 抄送人结构体
type Receiver struct {
	ID   string `json:"id"`   // 抄送人ID
	Name string `json:"name"` // 抄送人名称
}

type DownloadContractResponse struct {
	ResponseCode string `json:"responseCode"` // 响应码
	Message      string `json:"message"`      // 响应消息
	Result       Result `json:"result"`       // 返回数据
}

type Result struct {
	DownloadUrls []DownloadUrl `json:"downloadUrls"` // 下载链接列表
}

type DownloadUrl struct {
	ContractId    string       `json:"contractId"`    // 合同ID
	DocumentId    string       `json:"documentId"`    // 文档ID
	Title         string       `json:"title"`         // 文件名称
	DownloadItems string       `json:"downloadItems"` // 文件类型
	DownloadURL   string       `json:"downloadUrl"`   // 下载链接(有效期60分钟)
}


