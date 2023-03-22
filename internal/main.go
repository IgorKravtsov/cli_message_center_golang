package app

import (
	"errors"
	"fmt"
	"github.com/eiannone/keyboard"
	"github.com/joho/godotenv"
	"go_cli/internal/prog"
	"go_cli/internal/sender"
	"go_cli/pkg/interfaceManager"
	"go_cli/pkg/logger"
	"os"
)

func InitSenders() error {
	tgSender := sender.NewTelegramSender()
	if tgSender == nil {
		return errors.New("Telegram sender is nil")
	}
	dsSender := sender.NewDiscordSender()
	if dsSender == nil {
		return errors.New("Discord sender is nil")
	}
	slackSender := sender.NewSlackSender()
	if slackSender == nil {
		return errors.New("Slack sender is nil")
	}

	return nil
}

func printApp(app *prog.Prog) {
	if app.MessageSender != nil {
		fmt.Println("Selected messenger:", app.MessageSender.GetName())
	}
	fmt.Println("Please choose an option:")
	interfaceManager.PrintOptions(app.Options, app.SelectedOptionIndex)
}

func Run() {
	if err := godotenv.Load(); err != nil {
		logger.Error("error loading env variables: %s", err.Error())
	}

	if err := InitSenders(); err != nil {
		logger.Error(err)
		return
	}
	var app = prog.NewProg()

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	printApp(app)

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyArrowUp {
			app.Up()
		} else if key == keyboard.KeyArrowDown {
			app.Down()
		} else if key == keyboard.KeyEnter {
			app.Action()
		} else if key == keyboard.KeyEsc {
			os.Exit(0)
		}

		interfaceManager.Clear()
		printApp(app)
	}
}
