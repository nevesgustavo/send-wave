package ports

type MailService interface {
	SendMail(message string) error
}
