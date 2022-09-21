package cmds

import (
	"fmt"
	"github.com/shomali11/slacker"
	"github.com/slack-go/slack"
)

type CreateCVECommand struct {
	bot *slacker.Slacker
}

func (c *CreateCVECommand) GetHandler() *slacker.CommandDefinition {
	return &slacker.CommandDefinition{
		Description: "create and dispatch a new CVE",
		Examples:    []string{"cve create [hello]"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			number := request.Param("number")

			client := botCtx.Client()
			ev := botCtx.Event()

			if ev.Channel != "" {
				// sed-perry-cve-2022-1388
				channel, err := client.CreateConversation(fmt.Sprintf("sec-perry-cve-%s", number), true)
				if err != nil {
					fmt.Printf("fail to create channel")
				}
				_, _, _, err = client.JoinConversation(channel.ID)
				if err != nil {
					fmt.Printf("fail to join channel")
				}

				//client.InviteUsersToConversation()
				_, _, err = client.PostMessage(channel.ID, slack.MsgOptionText("Creating CVE.", false))
				if err != nil {
					fmt.Printf("fail to post message")
				}
			}

			response.Reply(fmt.Sprintf("successfully create sec-perry-cve-%s", number))
		},
	}
}
func (c *CreateCVECommand) GetQuestion() string {
	return "cve create {number}"
}

func NewCreateCVECommand(bot *slacker.Slacker) BotCommand {
	return &CreateCVECommand{bot: bot}
}
