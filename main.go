package main

import (
	"github.com/go-zoox/cli"
)

func main() {
	app := cli.NewSingleProgram(&cli.SingleProgramConfig{
		Name:    "chatgpt-for-chatbot-feishu",
		Usage:   "chatgpt-for-chatbot-feishu is a portable chatgpt server",
		Version: Version,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Usage:   "server port",
				Aliases: []string{"p"},
				EnvVars: []string{"PORT"},
				Value:   8080,
			},
			&cli.StringFlag{
				Name:    "api-path",
				Usage:   "custom api path, default: /",
				EnvVars: []string{"API_PATH"},
				Value:   "/",
			},
			&cli.StringFlag{
				Name:     "chatgpt-api-key",
				Usage:    "ChatGPT API Key",
				EnvVars:  []string{"CHATGPT_API_KEY"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "app-id",
				Usage:    "Feishu App ID",
				EnvVars:  []string{"APP_ID"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "app-secret",
				Usage:    "Feishu App SECRET",
				EnvVars:  []string{"APP_SECRET"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "encrypt-key",
				Usage:   "enable encryption if you need",
				EnvVars: []string{"ENCRYPT_KEY"},
			},
			&cli.StringFlag{
				Name:    "verification-token",
				Usage:   "enable token verification if you need",
				EnvVars: []string{"VERIFICATION_TOKEN"},
			},
			&cli.StringFlag{
				Name:    "report-url",
				Usage:   "Set error report url",
				EnvVars: []string{"REPORT_URL"},
			},
			&cli.StringFlag{
				Name:    "site-url",
				Usage:   "The Site URL",
				EnvVars: []string{"SITE_URL"},
			},
			&cli.StringFlag{
				Name:    "openai-model",
				Usage:   "Custom open ai model",
				EnvVars: []string{"OPENAI_MODEL"},
			},
			&cli.StringFlag{
				Name:    "feishu-base-uri",
				Usage:   "Custom feishu base uri for selfhosted Feishu",
				EnvVars: []string{"FEISHU_BASE_URI"},
			},
			&cli.StringFlag{
				Name:    "conversation-context",
				Usage:   "Custom chatgpt conversation context",
				EnvVars: []string{"CHATGPT_CONTEXT_MESSAGE"},
			},
			&cli.StringFlag{
				Name:    "conversation-language",
				Usage:   "Custom chatgpt conversation lanuage",
				EnvVars: []string{"CHATGPT_LANGUAGE"},
			},
			&cli.StringFlag{
				Name:    "logs-dir",
				Usage:   "The logs dir for save logs",
				EnvVars: []string{"LOGS_DIR"},
				Value:   "/tmp/chatgpt-for-chatbot-feishu",
			},
			&cli.StringFlag{
				Name:    "offline-message",
				Usage:   "The message to use for offline status",
				EnvVars: []string{"OFFLINE_MESSAGE"},
				Value:   "robot is offline",
			},
			&cli.StringFlag{
				Name:    "admin-email",
				Usage:   "Sets the admin with admin email, who can run commands",
				EnvVars: []string{"ADMIN_EMAIL"},
			},
			&cli.StringFlag{
				Name:    "bot-name",
				Usage:   "Sets the bot name, default: ChatGPT",
				EnvVars: []string{"BOT_NAME"},
			},
		},
	})

	app.Command(func(ctx *cli.Context) (err error) {
		return ServeFeishuBot(&FeishuBotConfig{
			LogsDir:           ctx.String("logs-dir"),
			Port:              ctx.Int64("port"),
			APIPath:           ctx.String("api-path"),
			ChatGPTAPIKey:     ctx.String("chatgpt-api-key"),
			AppID:             ctx.String("app-id"),
			AppSecret:         ctx.String("app-secret"),
			EncryptKey:        ctx.String("encrypt-key"),
			VerificationToken: ctx.String("verification-token"),
			ReportURL:         ctx.String("report-url"),
			SiteURL:           ctx.String("site-url"),
			OpenAIModel:       ctx.String("openai-model"),
			FeishuBaseURI:     ctx.String("feishu-base-uri"),
			ChatGPTContext:    ctx.String("chatgpt-context"),
			ChatGPTLanguage:   ctx.String("chatgpt-language"),
			OfflineMessage:    ctx.String("offline-message"),
			AdminEmail:        ctx.String("admin-email"),
			BotName:           ctx.String("bot-name"),
		})
	})

	app.Run()
}
