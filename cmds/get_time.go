package cmds

import (
	"github.com/shomali11/slacker"
	"strings"
	"time"
)

type TimeCommand struct {
	bot *slacker.Slacker
}

func (c *TimeCommand) GetHandler() *slacker.CommandDefinition {
	return &slacker.CommandDefinition{
		Description: "Get time in local or UTC ",
		Examples:    []string{" time in [local|UTC]"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			reply := ""
			t := time.Now()
			format := strings.ToLower(request.Param("format"))
			switch format {
			case "local":
				reply = (t.In(time.Local)).String()
			case "utc":
				reply = (t.In(time.UTC)).String()
			default:
				reply = "format not supported"
			}

			response.Reply(reply)
		},
	}
}
func (c *TimeCommand) GetQuestion() string {
	return "time in {format}"
}

func NewTimeCommand(bot *slacker.Slacker) BotCommand {
	return &TimeCommand{bot: bot}
}
