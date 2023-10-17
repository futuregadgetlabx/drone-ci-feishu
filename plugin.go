package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/futuregadgetlabx/drone-feishu/consts"
	"github.com/futuregadgetlabx/drone-feishu/request"
	feishuTemplate "github.com/futuregadgetlabx/drone-feishu/template"
	"github.com/google/uuid"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type (
	Repo struct {
		Owner string
		Name  string
		Url   string
	}

	Build struct {
		Tag               string
		Event             string
		Number            int
		Parent            int
		Commit            string
		CommitMessage     string
		CommitAuthor      CommitAuthor
		Ref               string
		Branch            string
		Pull              string
		PullRequestTitle  string
		Status            string
		Link              string
		Started           int64
		StartedFormatted  string
		Created           int64
		CreatedFormatted  string
		Finished          int64
		FinishedFormatted string
		CostTime          int64
		FailedStages      string
		FailedSteps       string
		SourceBranch      string
		TargetBranch      string
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

var eventStatusMap = map[string]map[string]string{
	"push": {
		"success": feishuTemplate.PushSuccess,
		"failure": feishuTemplate.PushFailure,
	},
	"pull_request": {
		"success": feishuTemplate.PrSuccess,
		"failure": feishuTemplate.PrFailure,
	},
}

func buildCommitMessage(m string) string {
	tmp := strings.ReplaceAll(m, "\r\n", "\\n")
	return strings.ReplaceAll(tmp, "\n", "\\n")
}

func (p Plugin) Exec() {
	// Get tenant access token
	tokenReq := request.GetTokenReq{
		AppID:     p.Feishu.AppID,
		AppSecret: p.Feishu.AppSecret,
	}
	reqBody, _ := json.Marshal(tokenReq)
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

	var originTmpl string
	if statusMap, ok := eventStatusMap[p.Build.Event]; ok {
		if originTmpl, ok = statusMap[p.Build.Status]; !ok {
			log.Fatal("unknown status")
		}
	} else {
		log.Fatal("unknown event")
	}

	p.Build.CreatedFormatted = time.Unix(p.Build.Created, 0).Format("2006-01-02 15:04:05")
	p.Build.StartedFormatted = time.Unix(p.Build.Started, 0).Format("2006-01-02 15:04:05")
	p.Build.FinishedFormatted = time.Unix(p.Build.Finished, 0).Format("2006-01-02 15:04:05")
	// DRONE_BUILD_FINISHED 与 DRONE_BUILD_STARTED 总是相同，且原因未知所以改为计算 Created 时间
	p.Build.CostTime = p.Build.Finished - p.Build.Started
	tmpl, err := template.New("template").Parse(originTmpl)
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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	// 读取响应的内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 打印响应内容
	fmt.Println("响应内容:", string(body))
}
