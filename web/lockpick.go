package web

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"strings"
	"github.com/apex/log"
)

// Lockpick will try different passwords
type Lockpick struct {
	Address            string
	PayloadTemplate    string
	UnsuccessfulString string
	Username           string
}

// TryLogin single login attempt
func (l *Lockpick) TryLogin(password string) (bool, error) {
	request := gorequest.New()
	payload := strings.Replace(l.PayloadTemplate, "{{username}}", l.Username, 1)
	payload = strings.Replace(payload, "{{password}}", password, 1)
	_, body, err := request.Post(l.Address).
		Type("form").
		Send(payload).
		End()

	if err != nil {
		return false, fmt.Errorf("Error sending request %v", err)
	}

	if strings.Contains(body, l.UnsuccessfulString) {
		log.WithField("Login", l.Username).WithField("Password", password).Info("Failed")
		return false, nil
	}
	log.WithField("Login", l.Username).WithField("Password", password).Info("SUCCESS")
	return true, nil
}
