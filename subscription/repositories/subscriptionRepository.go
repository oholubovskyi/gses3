package repositories

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/oholubovskyi/gses3/config"
	"github.com/oholubovskyi/gses3/subscription/models"
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

	decoder := json.NewDecoder(file)
	var subscriptionData models.SubscriptionData

	err = decoder.Decode(&subscriptionData)
	if err != nil {
		return nil, err
	}

	var emails = subscriptionData.Emails

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
	subscriptionsFile, err := os.OpenFile(path, os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer subscriptionsFile.Close()

	encoder := json.NewEncoder(subscriptionsFile)

	var subscriptionData = &models.SubscriptionData{
		Emails: emails,
	}
	err = encoder.Encode(subscriptionData)
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
