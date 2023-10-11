package main

import "strings"

type (
	Repo struct {
		Owner string
		Name  string
		Url   string
	}

	Build struct {
		Tag      string
		Event    string
		Number   int
		Parent   int
		Commit   string
		Ref      string
		Branch   string
		Author   Author
		Pull     string
		Message  Message
		DeployTo string
		Status   string
		Link     string
		Started  int64
		Created  int64
	}

	Author struct {
		Username string
		Name     string
		Email    string
		Avatar   string
	}

	Message struct {
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

func (a Author) String() string {
	return a.Username
}

func newCommitMessage(m string) Message {
	splitMsg := strings.Split(m, "\n")

	return Message{
		msg:   m,
		Title: strings.TrimSpace(splitMsg[0]),
		Body:  strings.TrimSpace(strings.Join(splitMsg[1:], "\n")),
	}
}

func (m Message) String() string {
	return m.msg
}
