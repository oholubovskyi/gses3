package services

import (
	"github.com/oholubovskyi/gses3/rate/repositories"
)

type RateService struct {
	rateRepo repositories.RateRepository
}

func NewRateService(rateRepo repositories.RateRepository) *RateService {
	return &RateService{
		rateRepo: rateRepo,
	}
}

func (s *RateService) GetBtcRate() (float32, error) {
	var btcRate = s.rateRepo.GetBtcRate()

	return btcRate, nil
}
