package dto

type MailResponse struct {
	Details     map[string]int `json:"details"`
	ElapsedTime string         `json:"elapsed_time"`
}
