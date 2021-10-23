package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go-kbs-notification/config"
	"go-kbs-notification/constants"
	"go-kbs-notification/handlers/rest"
	"go-kbs-notification/infra/redis"
	"go-kbs-notification/repositories"
	"go-kbs-notification/services"

	"github.com/gorilla/mux"
)

type SMSNotificationRequest struct {
	PhoneNumber string
	Message     string
}

func main() {
	cfg := config.NewConfig()

	redisConnOptions := redis.ConnOptions{
		Address: cfg.Databases.Redis.Host,
		Port:    cfg.Databases.Redis.Port,
		Timeout: 1 * time.Second,
	}

	redisClient := redis.NewClient(redis.ClientRedis)
	redisClient.Register(constants.RedisKBS, redisConnOptions)

	fetaureFlagRepository := repositories.NewFeatureFlagRepository(redisClient)
	featureFlagService := services.NewFeatureFlagService(fetaureFlagRepository)
	featureFlagRestHandler := rest.NewFeatureFlagRestHandler(featureFlagService)

	smsNotificationService := services.NewSMSNotificationService(fetaureFlagRepository)
	smsNotificationRestHandler := rest.NewSMSNotificationRestHandler(smsNotificationService)

	router := mux.NewRouter()
	router.Handle("/v1/feature-flag", featureFlagRestHandler.StoreFeatureFlag()).Methods(http.MethodPost)
	router.Handle("/v1/feature-flag", featureFlagRestHandler.GetFeatureFlag()).Methods(http.MethodGet)

	router.Handle("/v1/notification/sms", smsNotificationRestHandler.SendSMS()).Methods(http.MethodPost)

	// Setup http server
	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:" + cfg.App.Port,
		WriteTimeout: time.Duration(cfg.App.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.App.ReadTimeout) * time.Second,
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		log.Println("We received an interrupt signal, shut down.")
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
		log.Println("Bye.")
	}()
	log.Println("Listening on port " + cfg.App.Port)
	log.Fatal(srv.ListenAndServe())
	<-idleConnsClosed
	// BuildConfig()

	// req = getRequest()
	// cfg = GetConfig()
	// cfg.Vendor.SendSMS(req)
	//

	/*
		buat gatway struct berisikan mekanisme untuk mengirimkan data ke vendor-vendor

		buat facade dengan struct untuk menentukan vendor yang akan digunakan dengan flagging

		send notification sms
	*/
}
