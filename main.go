package main

import (
	"github.com/apex/log"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"bufio"
	"github.com/skarnecki/lockpick/web"
)

var (
	address = kingpin.Flag("address", "Full address to password form").Required().Short('a').String()
	username = kingpin.Flag("username", "Payload template.").Short('u').Required().String()
	dictionary = kingpin.Flag("dictionary", "Path to dictionary file.").Short('d').Required().ExistingFile()
	payload = kingpin.Flag("payload", "Payload template.").Default(`{"username": "{{username}}", "password": "{{password}}", "Login": "Login"}`).Short('p').String()
	unsuccessfulMessage = kingpin.Flag("message", "Message after unsuccesful login attempt.").Short('m').Default(`Login failed`).String()
)

// Loggable should perform login attempt
type Loggable interface {
	TryLogin (string) (bool, error)
}

func main() {
	kingpin.Parse()

	lock := &web.Lockpick{
		Address: *address,
		PayloadTemplate: *payload,
		UnsuccessfulString: *unsuccessfulMessage,
		Username: *username,
	}

	err := walkDict(*dictionary, lock)
	if err != nil {
		log.WithField("Dictionary", *dictionary).
			WithField("Address", lock.Address).
			WithField("PayloadTemplate", lock.PayloadTemplate).
			WithField("UnsuccessfulString", lock.UnsuccessfulString).
			WithField("Username", lock.Username).WithError(err).
			Error("Something went wrong")

	}
}

func walkDict(path string, handler Loggable) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result, err := handler.TryLogin(scanner.Text())
		if err != nil {
			return err
		}

		if result {
			return nil
		}
	}

	return scanner.Err()
}

