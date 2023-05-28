package handlers

import (
	"fmt"
	"net/http"

	"github.com/oholubovskyi/gses3/rate/services"
)

type RateHandler struct {
	exchangeRateSvc services.RateService
}

func NewRateHandler(exchangeRateSvc services.RateService) *RateHandler {
	return &RateHandler{
		exchangeRateSvc: exchangeRateSvc,
	}
}

func (s *RateHandler) GetBtcRate(w http.ResponseWriter, req *http.Request) {
	var rate, err = s.exchangeRateSvc.GetBtcRate()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%f", rate)))
}
