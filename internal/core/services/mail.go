package services

import "send-wave/internal/core/ports"

type MailService struct {
	mailSender ports.MailSender
}

func NewMailService(mailClient ports.MailSender) *MailService {
	return &MailService{
		mailSender: mailClient,
	}
}

func (m MailService) SendMail(message string) error {
	return m.mailSender.SendMail(message)
}
