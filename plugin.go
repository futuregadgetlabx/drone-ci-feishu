package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/futuregadgetlabx/drone-feishu/consts"
	"github.com/futuregadgetlabx/drone-feishu/request"
	"github.com/google/uuid"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

type (
	Repo struct {
		Owner string
		Name  string
		Url   string
	}

	Build struct {
		Tag              string
		Event            string
		Number           int
		Parent           int
		Commit           string
		CommitMessage    string
		CommitAuthor     CommitAuthor
		Ref              string
		Branch           string
		Pull             string
		PullRequestTitle string
		Status           string
		Link             string
		Started          string
		Created          string
		Finished         int64
		CostTime         int64
		FailedStages     string
		FailedSteps      string
	}

	CommitAuthor struct {
		Username string
		Name     string
		Email    string
		Avatar   string
	}

	CommitMessage struct {
		msg   string
		Title string
		Body  string
	}

	Feishu struct {
		UserID    string
		ChatID    string
		AppID     string
		AppSecret string
	}

	Plugin struct {
		Repo   Repo
		Build  Build
		Feishu Feishu
	}
)

func (p Plugin) Exec() error {
	// Get tenant access token
	tokenReq := request.GetTokenReq{
		AppID:     p.Feishu.AppID,
		AppSecret: p.Feishu.AppSecret,
	}
	reqBody, _ := json.Marshal(tokenReq)
	fmt.Println(reqBody)
	resp, err := http.Post(consts.GetTenantToken, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatalf("request feishu tenant_access_token error: %v", err)
	}
	respBody, _ := io.ReadAll(resp.Body)
	var tokenResp request.GetTokenResp
	err = json.Unmarshal(respBody, &tokenResp)
	if err != nil {
		log.Fatalf("unmarshal data error: %v", err)
	}
	if tokenResp.Code != 0 {
		log.Fatalf("request feishu tenant_access_token error: %v", tokenResp.Msg)
	}

	var filePath string
	if p.Build.Event == "push" {
		if p.Build.Status == "success" {
			filePath = "template/compile_success.json"
		} else {
			filePath = "template/compile_failure.json"
		}
	} else if p.Build.Event == "pull_request" {

	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New("template").Parse(string(file))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", p)
	var filledTemplate bytes.Buffer
	err = tmpl.Execute(&filledTemplate, p)
	if err != nil {
		log.Fatal(err)
	}

	sendMsgReq := request.SendMsgReq{
		ReceiveID: p.Feishu.UserID,
		MsgType:   "interactive",
		Content:   filledTemplate.String(),
		UUID:      uuid.New().String(),
	}
	sendMsgReqBody, _ := json.Marshal(sendMsgReq)
	httpPost, err := http.NewRequest("POST", consts.SendMsg, bytes.NewBuffer(sendMsgReqBody))
	if err != nil {
		log.Fatal(err)
	}

	httpPost.Header.Set("Authorization", "Bearer "+tokenResp.TenantAccessToken)
	// 发送请求
	client := &http.Client{}
	resp, err = client.Do(httpPost)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// 读取响应的内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return nil
	}

	// 打印响应内容
	fmt.Println("响应内容:", string(body))
	return nil
}
