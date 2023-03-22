package prog

import (
	"go_cli/internal/sender"
	"go_cli/pkg/interfaceManager"
	"go_cli/pkg/logger"
	"os"
)

const (
	SelectSender       = 0
	SelectSenderInner  = 1
	SendMessage        = 2
	ConfigureMessenger = 3
)

var SenderOptions = []string{"Telegram", "Discord", "Slack", "Back"}
var SelectSenderInnerOptions = []string{"Send message", "Configure", "Back"}

type Prog struct {
	state               int
	Options             []string
	SelectedOptionIndex int
	MessageSender       sender.Sender
}

func NewProg() *Prog {
	return &Prog{
		Options:             SenderOptions,
		state:               SelectSender,
		SelectedOptionIndex: 0,
		MessageSender:       nil}
}

func (p *Prog) Up() {
	if p.SelectedOptionIndex > 0 {
		p.SelectedOptionIndex--
	} else {
		p.SelectedOptionIndex = len(p.Options) - 1
	}
}

func (p *Prog) Down() {
	if p.SelectedOptionIndex < len(p.Options)-1 {
		p.SelectedOptionIndex++
	} else {
		p.SelectedOptionIndex = 0
	}
}

func (p *Prog) handleBack() {
	if p.state == SelectSender {
		os.Exit(0)
	} else if p.state == SelectSenderInner {
		p.Options = SenderOptions
		p.SelectedOptionIndex = 0
		p.state = SelectSender
		p.MessageSender = nil
	} else if p.state == SendMessage || p.state == ConfigureMessenger {
		p.Options = SelectSenderInnerOptions
		p.SelectedOptionIndex = 0
		p.state = SelectSenderInner
	}
}

func (p *Prog) selectSender() {
	switch p.Options[p.SelectedOptionIndex] {
	case "Telegram":
		p.MessageSender = sender.NewTelegramSender()
		break
	case "Discord":
		p.MessageSender = sender.NewDiscordSender()
		break
	case "Slack":
		p.MessageSender = sender.NewSlackSender()
		break
	default:
		p.handleBack()
		return
	}
	p.state = SelectSenderInner
	p.Options = SelectSenderInnerOptions
	p.SelectedOptionIndex = 0
}

func (p *Prog) updateInterfaceAndSendMessage() {
	p.state = SendMessage
	interfaceManager.Clear()
	interfaceManager.SetText("Please enter your message: ")

	message, err := interfaceManager.GetEnteredText()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	err = p.MessageSender.Send(message)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	p.handleBack()
}

func (p *Prog) configureMessageSender() {
	p.state = ConfigureMessenger
	err := p.MessageSender.Configure()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	p.handleBack()
}

func (p *Prog) Action() {
	if p.Options[p.SelectedOptionIndex] == "Back" {
		p.handleBack()
		return
	}
	switch p.state {
	case SelectSender:
		p.selectSender()
		break
	case SelectSenderInner:
		switch p.Options[p.SelectedOptionIndex] {
		case "Send message":
			p.updateInterfaceAndSendMessage()
			break
		case "Configure":
			p.configureMessageSender()
			break
		}
	}
}
