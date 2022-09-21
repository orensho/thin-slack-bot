package cmds

import (
	"github.com/shomali11/slacker"
)

type BotCommand interface {
	GetHandler() *slacker.CommandDefinition
	GetQuestion() string
}

type BotCmdsFactory struct {
	botCommands []BotCommand
	bot         *slacker.Slacker
}

func NewCmdsFactory(bot *slacker.Slacker) BotCmdsFactory {
	var botCommands []BotCommand
	// Declare all bot commands here
	botCommands = append(botCommands, NewEchoCommand(bot))
	botCommands = append(botCommands, NewTimeCommand(bot))
	botCommands = append(botCommands, NewCreateCVECommand(bot))

	return BotCmdsFactory{botCommands: botCommands, bot: bot}
}

func (b *BotCmdsFactory) Create() {
	for _, botCommand := range b.botCommands {
		b.bot.Command(botCommand.GetQuestion(), botCommand.GetHandler())
	}
}
