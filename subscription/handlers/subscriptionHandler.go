package handlers

import (
	"net/http"

	"github.com/oholubovskyi/gses3/subscription/services"
)

type SubscriptionHandler struct {
	subscribtionSvc services.SubscriptionService
}

func NewSubscriptionHandler(subscribtionSvc services.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{
		subscribtionSvc: subscribtionSvc,
	}
}

func (s *SubscriptionHandler) Subcribe(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var email = req.FormValue("email")
	var err = s.subscribtionSvc.Subscribe(email)

	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *SubscriptionHandler) SendEmails(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	s.subscribtionSvc.SendEmails()
	w.WriteHeader(http.StatusOK)
}
