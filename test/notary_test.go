package test

import (
	"fmt"
	"os"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
	"qiyuesuo/sdk/request"
	"testing"
)

/**
业务分类接口
*/

func TestChainNotary(t *testing.T) {
	req := request.ChainNotaryRequest{}
	req.NotaryName = "存证数据测试"
	req.NotaryDataID = "6331411786447082982"
	req.NotaryType = "FILE"
	req.FileType = "pdf"
	file, _ := os.Open("/Users/sgf/Downloads/文件模板三.pdf")
	req.NotaryFile = &http.FileItem{file, ""}
	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestOpenChainEvidenceProve(t *testing.T) {
	req := request.OpenChainEvidenceProveRequest{}
	req.ChainId = "3331622985832780333"
	req.ProveSource = "CUSTOM_DATA"
	user := model.UserInfoRequest{}
	user.Name = "宋一"
	user.Contact = "10000000001"
	user.ContactType = "MOBILE"
	req.OperatorInfo = &user
	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestOpenCompleteReportApply(t *testing.T) {
	req := request.OpenCompleteReportApplyRequest{}
	req.ChainId = "3331622985832780333"
	proveSource := 4
	req.ProveSource = &proveSource
	file, _ := os.Open("/Users/sgf/Downloads/文件模板三.pdf")
	var fis []*http.FileItem
	fis = append(fis, &http.FileItem{Src: file})
	req.ProveMaterial = fis
	user := model.UserInfoRequest{}
	user.Name = "宋一"
	user.Contact = "10000000001"
	user.ContactType = "MOBILE"
	req.OperatorInfo = &user
	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestOpenCompleteReportStatus(t *testing.T) {
	req := request.OpenCompleteReportStatusRequest{}
	req.ProveId = "3331627587583345265"
	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestOpenCompleteReportDownload(t *testing.T) {
	req := request.OpenCompleteReportDownloadRequest{}
	req.ProveId = "3331627587583345265"
	file, _ := os.OpenFile("/Users/sgf/develop/临时/go/go-存证报告.pdf", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer file.Close()
	err := sdkClient.Download(req, file)
	if err != nil {
		fmt.Println("error,", err.Error())
	}
}
