package cmds

import (
	"github.com/shomali11/slacker"
)

type EchoCommand struct {
	bot *slacker.Slacker
}

func (c *EchoCommand) GetHandler() *slacker.CommandDefinition {
	return &slacker.CommandDefinition{
		Description: "Echo a word!",
		Examples:    []string{"echo [hello]"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			word := request.Param("word")
			response.Reply(word)
		},
	}
}
func (c *EchoCommand) GetQuestion() string {
	return "echo {word}"
}

func NewEchoCommand(bot *slacker.Slacker) BotCommand {
	return &EchoCommand{bot: bot}
}
