package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
	"log"
	"os"
)

const Send_Message_API = "https://open.feishu.cn/open-apis/im/v1/messages?receive_id_type=user_id"

var (
	version = "0.0.1"
	template       = `{"config":{"wide_screen_mode":true},"elements":[{"fields":[{"is_short":true,"text":{"content":"**构建时间**\n2023-10-11 18:11","tag":"lark_md"}},{"is_short":true,"text":{"content":"**仓库地址**\n[futuregadgetlabx/ohmyhelper-fe](https://github.com/futuregadgetlabx/ohmyhelper-fe)\n","tag":"lark_md"}}],"tag":"div"},{"tag":"markdown","content":"👨🏻\u200d💻 提交人: <at id=cruii></at>\n✅ Git Commit: [f421355](https://github.com/futuregadgetlabx/ohmyhelper-fe/commit/f421355)\n🛠️ 构建任务: [](https://drone.company.com/octocat/hello-world/42)\n✉️ PR Title: Update .drone.yml\n⏱️ 构建耗时: 128s"},{"tag":"action","actions":[{"tag":"button","text":{"tag":"plain_text","content":"部署上线"},"type":"primary","multi_url":{"url":"futuregadgetlabx.com","pc_url":"","android_url":"","ios_url":""}},{"tag":"button","text":{"tag":"plain_text","content":"我已知悉"},"type":"default"}]},{"tag":"hr"},{"elements":[{"content":"[来自未来ガジェット研究所](hhttps://github.com/futuregadgetlabx)","tag":"lark_md"}],"tag":"note"}],"header":{"template":"green","title":{"content":"【Drone CI】代码编译构建成功","tag":"plain_text"}}}`
)

func run(c *cli.Context) error {
	plugin := Plugin{
		Repo: Repo{
			Owner: c.String("repo.owner"),
			Name:  c.String("repo.name"),
			Url:   c.String("repo.url"),
		},
		Build: Build{
			Tag:    c.String("build.tag"),
			Number: c.Int("build.number"),
			Parent: c.Int("build.parent"),
			Event:  c.String("build.event"),
			Status: c.String("build.status"),
			Commit: c.String("commit.sha"),
			Ref:    c.String("commit.ref"),
			Branch: c.String("commit.branch"),
			Author: Author{
				Username: c.String("commit.author"),
				Name:     c.String("commit.author.name"),
				Email:    c.String("commit.author.email"),
				Avatar:   c.String("commit.author.avatar"),
			},
			Pull:     c.String("commit.pull"),
			Message:  newCommitMessage(c.String("commit.message")),
			DeployTo: c.String("build.deployTo"),
			Link:     c.String("build.link"),
			Started:  c.Int64("build.started"),
			Created:  c.Int64("build.created"),
		},
		Feishu: Feishu{
			UserID:    c.String("user_id"),
			ChatID:    c.String("chat_id"),
			AppID:     c.String("app_id"),
			AppSecret: c.String("app_secret"),
		},
	}
	if plugin.Build.Commit == "" {
		plugin.Build.Commit = "0000000000000000000000000000000000000000"
	}

	fmt.Printf("%+v", plugin)
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "Feishu Notification"
	app.Usage = "Feishu Notification"
	app.Authors = []cli.Author{
		{
			Name:  "Cruii",
			Email: "cruii811@gmail.com",
		},
	}
	app.Action = run
	app.Version = fmt.Sprintf("%s", version)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "user_id",
			Usage:  "飞书用户ID，与飞书群聊ID同时配置时不触发单聊",
			EnvVar: "PLUGIN_USER_ID",
		},
		cli.StringFlag{
			Name:   "chat_id",
			Usage:  "飞书群聊ID",
			EnvVar: "PLUGIN_CHAT_ID",
		},
		cli.StringFlag{
			Name:   "app_id",
			Usage:  "飞书机器人App ID",
			EnvVar: "PLUGIN_APP_ID",
		},
		cli.StringFlag{
			Name:   "app_secret",
			Usage:  "飞书机器人App Secret",
			EnvVar: "PLUGIN_APP_SECRET",
		},
		cli.StringFlag{
			Name:   "repo.owner",
			Usage:  "repository owner",
			EnvVar: "DRONE_REPO_OWNER",
		},
		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "repository name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "repo.url",
			Usage:  "Provides the repository link for the current running build.",
			EnvVar: "DRONE_REPO_LINK",
		},
		cli.StringFlag{
			Name:   "commit.sha",
			Usage:  "git commit sha",
			EnvVar: "DRONE_COMMIT_SHA",
			Value:  "00000000",
		},
		cli.StringFlag{
			Name:   "commit.ref",
			Value:  "refs/heads/master",
			Usage:  "git commit ref",
			EnvVar: "DRONE_COMMIT_REF",
		},
		cli.StringFlag{
			Name:   "commit.branch",
			Value:  "master",
			Usage:  "git commit branch",
			EnvVar: "DRONE_COMMIT_BRANCH",
		},
		cli.StringFlag{
			Name:   "commit.author",
			Usage:  "git author username",
			EnvVar: "DRONE_COMMIT_AUTHOR",
		},
		cli.StringFlag{
			Name:   "commit.author.email",
			Usage:  "git author email",
			EnvVar: "DRONE_COMMIT_AUTHOR_EMAIL",
		},
		cli.StringFlag{
			Name:   "commit.author.avatar",
			Usage:  "git author avatar",
			EnvVar: "DRONE_COMMIT_AUTHOR_AVATAR",
		},
		cli.StringFlag{
			Name:   "commit.author.name",
			Usage:  "git author name",
			EnvVar: "DRONE_COMMIT_AUTHOR_NAME",
		},
		cli.StringFlag{
			Name:   "commit.pull",
			Usage:  "git pull request",
			EnvVar: "DRONE_PULL_REQUEST",
		},
		cli.StringFlag{
			Name:   "commit.message",
			Usage:  "commit message",
			EnvVar: "DRONE_COMMIT_MESSAGE",
		},
		cli.StringFlag{
			Name:   "build.event",
			Value:  "push",
			Usage:  "build event",
			EnvVar: "DRONE_BUILD_EVENT",
		},
		cli.IntFlag{
			Name:   "build.number",
			Usage:  "build number",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.IntFlag{
			Name:   "build.parent",
			Usage:  "build parent",
			EnvVar: "DRONE_BUILD_PARENT",
		},
		cli.StringFlag{
			Name:   "build.status",
			Usage:  "build status",
			Value:  "success",
			EnvVar: "DRONE_BUILD_STATUS",
		},
		cli.StringFlag{
			Name:   "build.link",
			Usage:  "build link",
			EnvVar: "DRONE_BUILD_LINK",
		},
		cli.Int64Flag{
			Name:   "build.started",
			Usage:  "build started",
			EnvVar: "DRONE_BUILD_STARTED",
		},
		cli.Int64Flag{
			Name:   "build.created",
			Usage:  "build created",
			EnvVar: "DRONE_BUILD_CREATED",
		},
		cli.StringFlag{
			Name:   "build.tag",
			Usage:  "build tag",
			EnvVar: "DRONE_TAG",
		},
		cli.StringFlag{
			Name:   "build.deployTo",
			Usage:  "environment deployed to",
			EnvVar: "DRONE_DEPLOY_TO",
		},
	}

	if _, err := os.Stat("/run/drone/env"); err == nil {
		_ = godotenv.Overload("/run/drone/env")
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
