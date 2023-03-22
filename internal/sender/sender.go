package sender

type Sender interface {
	GetName() string
	Send(msg string) error
	Configure() error
}
