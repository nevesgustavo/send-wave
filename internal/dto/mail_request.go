package dto

import (
	"encoding/json"
	"fmt"
	"strings"
)

type MailMessage struct {
	Message string `json:"message"`
}

func (m MailMessage) Validate() error {
	errors := map[string]string{}

	if len(strings.Trim(m.Message, " ")) == 0 {
		errors["message"] = "message is required"
	}

	if len(errors) > 0 {
		bs, _ := json.Marshal(errors)
		return fmt.Errorf("validation errors: %v", string(bs))
	}

	return nil
}
