package user

import (
	"encoding/json"
	"io"
)

type Object struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	TimesReceived int64  `json:"times_received"`
}

func (u *Object) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(u)
}

func (u *Object) Encode() ([]byte, error) {
	return json.Marshal(u)
}
