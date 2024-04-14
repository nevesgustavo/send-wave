package ports

type MailSender interface {
	SendMail(message string) error
}
