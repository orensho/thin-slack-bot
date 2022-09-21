package main

import (
	"context"
	"fmt"
	"github.com/orensho/thin-slack-bot/cmds"
	"github.com/shomali11/slacker"
	log "github.com/sirupsen/logrus"
	"os"
)

func InitBot() *slacker.Slacker {
	log.Info("starting slack thin bot")
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"), slacker.WithDebug(true))
	log.Info("slack thin bot started")

	// Terminate if bot cannot be started
	bot.Err(func(err string) {
		log.Errorf("failed to init bot: %s", err)
	})

	// Define a default response (if nothing else matches)
	bot.DefaultCommand(func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
		response.Reply(fmt.Sprintf("Sorry. the command \"%s\" is not understood. Please type \"*help*\" to get the list of supported commands", botCtx.Event().Text))
	})

	log.Debugf("Registering bot commands")
	loadBotcommands(bot)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return bot
}

func loadBotcommands(bot *slacker.Slacker) {
	botCmdsFactory := cmds.NewCmdsFactory(bot)
	botCmdsFactory.Create()
}
