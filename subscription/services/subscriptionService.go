package services

import (
	"errors"
	"fmt"
	"net/smtp"

	"github.com/oholubovskyi/gses3/config"
	"github.com/oholubovskyi/gses3/rate/services"
	"github.com/oholubovskyi/gses3/subscription/repositories"
)

type SubscriptionService struct {
	config           config.Config
	subscribtionRepo repositories.SubscriptionRepository
	rateSvc          services.RateService
}

func NewSubscriptionService(
	config config.Config,
	subscribtionRep repositories.SubscriptionRepository,
	rateSvc services.RateService) *SubscriptionService {
	return &SubscriptionService{
		config:           config,
		subscribtionRepo: subscribtionRep,
		rateSvc:          rateSvc,
	}
}

func (s *SubscriptionService) Subscribe(email string) error {
	isExists, err := s.subscribtionRepo.IsEmailExist(email)
	if err != nil {
		return err
	}

	if !*isExists {
		err = s.subscribtionRepo.AddEmail(email)
		if err != nil {
			return err
		}

		return nil
	}

	var emailExistsErrorMsg = "email exists already"
	return errors.New(emailExistsErrorMsg)
}

func (s *SubscriptionService) SendEmails() error {
	emails, err := s.subscribtionRepo.GetAllEmails()
	if err != nil {
		return err
	}

	btcRate, err := s.rateSvc.GetBtcRate()
	if err != nil {
		return err
	}

	err = SendEmails(s.config, emails, fmt.Sprintf("%f", btcRate))
	if err != nil {
		return err
	}

	return nil
}

func SendEmails(config config.Config, emails []string, msg string) error {
	auth := smtp.PlainAuth("", config.Smtp.Sender, config.Smtp.Password, config.Smtp.SmtpServer)

	err := smtp.SendMail(fmt.Sprintf("%s:%d", config.Smtp.SmtpServer, config.Smtp.SmtpPort), auth, config.Smtp.Sender, emails, []byte(msg))
	if err != nil {
		return err
	}

	return nil
}
