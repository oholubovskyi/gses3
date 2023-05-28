package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/oholubovskyi/gses3/config"
	rateHandlerPkg "github.com/oholubovskyi/gses3/rate/handlers"
	rateRepoPkg "github.com/oholubovskyi/gses3/rate/repositories"
	rateSvcPkg "github.com/oholubovskyi/gses3/rate/services"
	subscriptionHandlerPkg "github.com/oholubovskyi/gses3/subscription/handlers"
	subscriptionRepoPkg "github.com/oholubovskyi/gses3/subscription/repositories"
	subscriptionSvcPkg "github.com/oholubovskyi/gses3/subscription/services"
)

func main() {
	config, err := loadConfig("config.json")
	if err != nil {
		return
	}

	registerSubscriptionRoutes(*config)
	registerRateRoutes(*config)

	http.ListenAndServe(config.Server.Port, nil)
}

func loadConfig(file string) (*config.Config, error) {
	configFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	var config config.Config
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func registerSubscriptionRoutes(config config.Config) {
	var subscriptionHandler = subscriptionHandlerPkg.NewSubscriptionHandler(*subscriptionSvcPkg.NewSubscriptionService(config, subscriptionRepoPkg.SubscriptionRepository{}))

	http.HandleFunc(config.Routes.Subscribe, subscriptionHandler.Subcribe)
	http.HandleFunc(config.Routes.SendEmails, subscriptionHandler.SendEmails)
}

func registerRateRoutes(config config.Config) {
	var rateHandler = rateHandlerPkg.NewRateHandler(*rateSvcPkg.NewRateService(rateRepoPkg.RateRepository{}))

	http.HandleFunc(config.Routes.Rate, rateHandler.GetBtcRate)
}
