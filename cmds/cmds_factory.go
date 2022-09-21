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
}

func NewCmdsFactory(bot *slacker.Slacker) BotCmdsFactory {
	var botCommands []BotCommand

	// Declare all bot commands here - each command must return a BotCommand interface
	botCommands = append(botCommands, NewEchoCommand(bot))
	botCommands = append(botCommands, NewTimeCommand(bot))

	return BotCmdsFactory{botCommands: botCommands}
}

func (b *BotCmdsFactory) Create(bot *slacker.Slacker) {
	for _, botCommand := range b.botCommands {
		bot.Command(botCommand.GetQuestion(), botCommand.GetHandler())
	}
}
