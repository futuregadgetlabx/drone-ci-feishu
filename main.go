package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
	"log"
	"os"
	"time"
)

var (
	version = "0.0.1"
)

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
			Value:  "refs/heads/main",
			Usage:  "git commit ref",
			EnvVar: "DRONE_COMMIT_REF",
		},
		cli.StringFlag{
			Name:   "commit.branch",
			Value:  "main",
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
			Name:   "commit.pull.title",
			Usage:  "git pull request title",
			EnvVar: "DRONE_PULL_REQUEST_TITLE",
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
		cli.Int64Flag{
			Name:   "build.finished",
			Usage:  "build finished",
			EnvVar: "DRONE_BUILD_FINISHED",
		},
		cli.StringFlag{
			Name:   "build.tag",
			Usage:  "build tag",
			EnvVar: "DRONE_TAG",
		},
		cli.StringFlag{
			Name:   "build.failed.stages",
			Usage:  "build failed stages",
			EnvVar: "DRONE_FAILED_STAGES",
		},
		cli.StringFlag{
			Name:   "build.failed.steps",
			Usage:  "build failed steps",
			EnvVar: "DRONE_FAILED_STEPS",
		},
		cli.StringFlag{
			Name:   "build.source_branch",
			Usage:  "build source branch",
			EnvVar: "DRONE_SOURCE_BRANCH",
		},
		cli.StringFlag{
			Name:   "build.target_branch",
			Usage:  "build target branch",
			EnvVar: "DRONE_TARGET_BRANCH",
		},
	}

	if _, err := os.Stat("/run/drone/env"); err == nil {
		_ = godotenv.Overload("/run/drone/env")
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// run executes the application
func run(c *cli.Context) {
	plugin := Plugin{
		Repo: Repo{
			Owner: c.String("repo.owner"),
			Name:  c.String("repo.name"),
			Url:   c.String("repo.url"),
		},
		Build: Build{
			Number: c.Int("build.number"),
			Event:  c.String("build.event"),
			Status: c.String("build.status"),
			Commit: c.String("commit.sha")[:8],
			Ref:    c.String("commit.ref"),
			Branch: c.String("commit.branch"),
			CommitAuthor: CommitAuthor{
				Username: c.String("commit.author"),
				Name:     c.String("commit.author.name"),
				Email:    c.String("commit.author.email"),
				Avatar:   c.String("commit.author.avatar"),
			},
			Pull:             c.String("commit.pull"),
			PullRequestTitle: c.String("commit.pull.title"),
			CommitMessage:    buildCommitMessage(c.String("commit.message")),
			Link:             c.String("build.link"),
			Started:          time.Unix(c.Int64("build.started"), 0).Format("2006-01-02 15:04:05"),
			Created:          time.Unix(c.Int64("build.created"), 0).Format("2006-01-02 15:04:05"),
			Finished:         c.Int64("build.finished"),
			CostTime:         (c.Int64("build.finished") - c.Int64("build.started")) / 1000,
			FailedStages:     c.String("build.failed.stages"),
			FailedSteps:      c.String("build.failed.steps"),
			TargetBranch:     c.String("build.target_branch"),
			SourceBranch:     c.String("build.source_branch"),
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

	plugin.Exec()
}
