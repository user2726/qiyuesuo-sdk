package test

import (
	"fmt"
	"os"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/request"
	"testing"
)

/**
业务分类接口
*/

func TestIdcardOcr(t *testing.T) {
	req := request.IdCardOcrRequest{}
	file, _ := os.Open("D:\\file\\inet_idcard.png")
	req.Image = &http.FileItem{file, ""}

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestBankCardOcr(t *testing.T) {
	req := request.BankCardOcrRequest{}
	file, _ := os.Open("D:\\file\\银行卡照片.png")
	req.Image = &http.FileItem{file, ""}

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}
