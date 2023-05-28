package repositories

type RateRepository struct {
}

func (s *RateRepository) GetBtcRate() float32 {
	return 12345
}
