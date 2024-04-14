package mail

import "fmt"

type LocalMailSender struct {
}

func NewLocalMail() *LocalMailSender {
	return &LocalMailSender{}
}

func (l LocalMailSender) SendMail(message string) error {
	fmt.Println("sending mail with message: ", message)
	return nil
}
