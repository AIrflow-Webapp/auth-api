package auth

import (
	"encoding/json"
)

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (login *LoginDTO) Decode(body []byte) error {
	if err := json.Unmarshal(body, login); err != nil {
		return err
	}
	return nil
}
