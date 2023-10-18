package template

const (
	PushSuccess = `
{
  "config": {
    "wide_screen_mode": true
  },
  "elements": [
    {
      "fields": [
        {
          "is_short": true,
          "text": {
            "content": "**Build Started**\n{{ .Build.StartedFormatted }}",
            "tag": "lark_md"
          }
        },
        {
          "is_short": true,
          "text": {
            "content": "**Repo**\n[{{ .Repo.Owner }}/{{ .Repo.Name }}]({{ .Repo.Url }})\n",
            "tag": "lark_md"
          }
        }
      ],
      "tag": "div"
    },
    {
      "tag": "markdown",
      "content": "👨🏻‍💻 **Built By** [@{{ .Build.CommitAuthor.Username }}](https://github.com/{{ .Build.CommitAuthor.Username }})\n🔀 **Built Branch**: [{{ .Build.Branch }}]({{ .Repo.Url }}/tree/{{ .Build.Branch }})\n🚀 **Commit**: [{{ .Build.Commit }}]({{ .Repo.Url }}/commit/{{ .Build.Commit }})\n📝 **Commit message**: \n{{ .Build.CommitMessage }}\n🛠️ **Build Details**: [#{{ .Build.Number }}]({{ .Build.Link }})\n⏱️ **Duration**: {{ .Build.CostTime }}s"
    },
    {
      "tag": "action",
      "actions": [
        {
          "tag": "button",
          "text": {
            "tag": "plain_text",
            "content": "Release"
          },
          "type": "primary",
          "multi_url": {
            "url": "https://open.feishu.cn/document",
            "android_url": "",
            "ios_url": "",
            "pc_url": ""
          }
        },
        {
          "tag": "button",
          "text": {
            "tag": "plain_text",
            "content": "See Build Details"
          },
          "type": "default",
          "multi_url": {
            "url": "{{ .Build.Link }}",
            "pc_url": "",
            "android_url": "",
            "ios_url": ""
          }
        },
        {
          "tag": "button",
          "text": {
            "tag": "plain_text",
            "content": "Dismiss"
          },
          "type": "default",
          "multi_url": {
            "url": "https://open.feishu.cn/document",
            "android_url": "",
            "ios_url": "",
            "pc_url": ""
          }
        }
      ]
    },
    {
      "tag": "hr"
    },
    {
      "elements": [
        {
          "content": "[来自未来ガジェット研究所 - 电话微波炉（暂定）](https://github.com/futuregadgetlabx)\n",
          "tag": "lark_md"
        }
      ],
      "tag": "note"
    }
  ],
  "header": {
    "template": "green",
    "title": {
      "content": "🎉【Drone CI】编译成功",
      "tag": "plain_text"
    }
  }
}
`

	PrSuccess = `
{
  "config": {
    "wide_screen_mode": true
  },
  "elements": [
    {
      "fields": [
        {
          "is_short": true,
          "text": {
            "content": "**Build Started**\n{{ .Build.StartedFormatted }}",
            "tag": "lark_md"
          }
        },
        {
          "is_short": true,
          "text": {
            "content": "**Repo**\n[{{ .Repo.Owner}}/{{ .Repo.Name }}]({{ .Repo.Url }})\n",
            "tag": "lark_md"
          }
        }
      ],
      "tag": "div"
    },
    {
      "tag": "markdown",
      "content": "👨🏻‍💻 **Built By** [@{{ .Build.CommitAuthor.Username }}](https://github.com/{{ .Build.CommitAuthor.Username }})\n📌 **Pull Request**: ***{{ .Build.SourceBranch }} ➡️ {{ .Build.TargetBranch }}***\n{{ .Build.PullRequestTitle }}\n🚀 **Commit**: [{{ .Build.Commit }}]({{ .Repo.Url }}/commit/{{ .Build.Commit }})\n📝 **Commit Message**: \n{{ .Build.CommitMessage }}\n🛠️ **Build Details**: [#{{ .Build.Number }}]({{ .Build.Link }})\n⏱️ **Duration**: {{ .Build.CostTime }}s"
    },
    {
      "tag": "action",
      "actions": [
        {
          "tag": "button",
          "text": {
            "tag": "plain_text",
            "content": "See PR Details"
          },
          "type": "primary",
          "multi_url": {
            "url": "{{ .Repo.Url }}/pull/{{ .Build.Pull }}",
            "pc_url": "",
            "android_url": "",
            "ios_url": ""
          }
        },
        {
          "tag": "button",
          "text": {
            "tag": "plain_text",
            "content": "Dismiss"
          },
          "type": "default",
          "multi_url": {
            "url": "https://open.feishu.cn/document",
            "android_url": "",
            "ios_url": "",
            "pc_url": ""
          }
        }
      ]
    },
    {
      "tag": "hr"
    },
    {
      "elements": [
        {
          "content": "[来自未来ガジェット研究所 - 电话微波炉（暂定）](https://github.com/futuregadgetlabx)",
          "tag": "lark_md"
        }
      ],
      "tag": "note"
    }
  ],
  "header": {
    "template": "green",
    "title": {
      "content": "🎉【Drone CI】编译成功",
      "tag": "plain_text"
    }
  }
}
`

	PushFailure = `
{
  "config": {
    "wide_screen_mode": true
  },
  "elements": [
    {
      "fields": [
        {
          "is_short": true,
          "text": {
            "content": "**Build Started**\n{{ .Build.StartedFormatted }}",
            "tag": "lark_md"
          }
        },
        {
          "is_short": true,
          "text": {
            "content": "**Repo**\n[{{ .Repo.Owner }}/{{ .Repo.Name }}]({{ .Repo.Url }})\n",
            "tag": "lark_md"
          }
        }
      ],
      "tag": "div"
    },
    {
      "tag": "markdown",
      "content": "👨🏻‍💻 **Built By** [@{{ .Build.CommitAuthor.Username }}](https://github.com/{{ .Build.CommitAuthor.Username }})\n🔀 **Built Branch**: [{{ .Build.Branch }}]({{ .Repo.Url }}/tree/{{ .Build.Branch }})\n🚀 **Commit**: [{{ .Build.Commit }}]({{ .Repo.Url }}/commit/{{ .Build.Commit }})\n📝 **Commit message**: \n{{ .Build.CommitMessage }}\n🛠️ **Build Details**: [#{{ .Build.Number }}]({{ .Build.Link }})\n❌ **Failed Stages**: {{ .Build.FailedStages }}\n🔥️ **Failed Steps**: {{ .Build.FailedSteps }}"
    },
    {
      "tag": "action",
      "actions": [
        {
          "tag": "button",
          "text": {
            "tag": "plain_text",
            "content": "See Build Details"
          },
          "type": "primary",
          "multi_url": {
            "url": "{{ .Build.Link }}",
            "pc_url": "",
            "android_url": "",
            "ios_url": ""
          }
        },
        {
          "tag": "button",
          "text": {
            "tag": "plain_text",
            "content": "See Commit Details"
          },
          "type": "default",
          "multi_url": {
            "url": "{{ .Repo.Url }}/commit/{{ .Build.Commit }}",
            "pc_url": "",
            "android_url": "",
            "ios_url": ""
          }
        },
        {
          "tag": "button",
          "text": {
            "tag": "plain_text",
            "content": "Dismiss"
          },
          "type": "default",
          "multi_url": {
            "url": "https://open.feishu.cn/document",
            "android_url": "",
            "ios_url": "",
            "pc_url": ""
          }
        }
      ]
    },
    {
      "tag": "hr"
    },
    {
      "elements": [
        {
          "content": "[来自未来ガジェット研究所 - 电话微波炉（暂定）](https://github.com/futuregadgetlabx)",
          "tag": "lark_md"
        }
      ],
      "tag": "note"
    }
  ],
  "header": {
    "template": "red",
    "title": {
      "content": "🚒【Drone CI】编译失败",
      "tag": "plain_text"
    }
  }
}
`

	PrFailure = `
{
  "config": {
    "wide_screen_mode": true
  },
  "elements": [
    {
      "fields": [
        {
          "is_short": true,
          "text": {
            "content": "**Build Started**\n{{ .Build.StartedFormatted }}",
            "tag": "lark_md"
          }
        },
        {
          "is_short": true,
          "text": {
            "content": "**Git Repo**\n[{{ .Repo.Owner }}/{{ .Repo.Name }}]({{ .Repo.Url }})\n",
            "tag": "lark_md"
          }
        }
      ],
      "tag": "div"
    },
    {
      "tag": "markdown",
      "content": "👨🏻‍💻 **Built By** [@{{ .Build.CommitAuthor.Username }}](https://github.com/{{ .Build.CommitAuthor.Username }})\n📌 **Pull Request**: ***{{ .Build.SourceBranch }} ➡️ {{ .Build.TargetBranch }}***\n{{ .Build.PullRequestTitle }}\n🚀 **Commit**: [{{ .Build.Commit }}]({{ .Repo.Url }}/commit/{{ .Build.Commit }})\n📝 **Commit Message**: \n{{ .Build.CommitMessage }}\n🛠️ **Build details**: [#{{ .Build.Number }}]({{ .Build.Link }})\n❌ **Failed Stages**: {{ .Build.FailedStages }}\n🔥️ **Failed Steps**: {{ .Build.FailedSteps }}"
    },
    {
      "tag": "action",
      "actions": [
        {
          "tag": "button",
          "text": {
            "tag": "plain_text",
            "content": "See Build Details"
          },
          "type": "primary",
          "multi_url": {
            "url": "{{ .Build.Link }}",
            "pc_url": "",
            "android_url": "",
            "ios_url": ""
          }
        },
        {
          "tag": "button",
          "text": {
            "tag": "plain_text",
            "content": "See PR Details"
          },
          "type": "default",
          "multi_url": {
            "url": "{{ .Repo.Url }}/pull/{{ .Build.Pull }}",
            "pc_url": "",
            "android_url": "",
            "ios_url": ""
          }
        },
        {
          "tag": "button",
          "text": {
            "tag": "plain_text",
            "content": "Dismiss"
          },
          "type": "default",
          "multi_url": {
            "url": "https://open.feishu.cn/document",
            "android_url": "",
            "ios_url": "",
            "pc_url": ""
          }
        }
      ]
    },
    {
      "tag": "hr"
    },
    {
      "elements": [
        {
          "content": "[来自未来ガジェット研究所 - 电话微波炉（暂定）](https://github.com/futuregadgetlabx)",
          "tag": "lark_md"
        }
      ],
      "tag": "note"
    }
  ],
  "header": {
    "template": "red",
    "title": {
      "content": "🚒【Drone CI】构建失败",
      "tag": "plain_text"
    }
  }
}
`
)
