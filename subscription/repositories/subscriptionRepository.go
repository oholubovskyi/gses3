package repositories

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/oholubovskyi/gses3/config"
)

type SubscriptionRepository struct {
	config config.Config
}

func NewSubscriptionRepository(config config.Config) *SubscriptionRepository {
	return &SubscriptionRepository{
		config: config,
	}
}

func (s *SubscriptionRepository) GetAllEmails() ([]string, error) {
	file, err := os.Open(s.config.Storage.Subscriptions)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	jsonBytes, err := ioutil.ReadAll(file)

	var emails []string
	err = json.Unmarshal(jsonBytes, &emails)
	if err != nil {
		return nil, err
	}

	return emails, nil
}

func (s *SubscriptionRepository) IsEmailExist(email string) (*bool, error) {
	emails, err := s.GetAllEmails()
	if err != nil {
		return nil, err
	}

	var isExists = stringInSlice(email, emails)
	return &isExists, nil
}

func (s *SubscriptionRepository) AddEmail(email string) error {
	emails, err := s.GetAllEmails()	
	if err != nil {
		return err
	}

	if stringInSlice(email, emails) {
		return errors.New("email subscribed already")
	}

	emails = append(emails, email)
	err = saveToFile(s.config.Storage.Subscriptions, emails)
	if err != nil {
		return err
	}

	return nil
}

func saveToFile(path string, emails []string) error {
	subscriptionsFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer subscriptionsFile.Close()

	emailsJson, err := json.Marshal(emails)
	if err != nil {
		return err
	}

	_, err = subscriptionsFile.Write(emailsJson)
	if err != nil {
		return err
	}

	err = subscriptionsFile.Sync()
	if err != nil {
		return err
	}

	return nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
