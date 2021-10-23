package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go-kbs-soccer/config"
	"go-kbs-soccer/handlers/rest"
	"go-kbs-soccer/repositories"
	"go-kbs-soccer/services"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.NewConfig()

	teamRepository := repositories.NewTeamRepository()
	teamService := services.NewTeamService(teamRepository)
	teamRestHandler := rest.NewTeamRestHandler(teamService)

	playerRepository := repositories.NewPlayerRepository()
	playerService := services.NewPlayerService(playerRepository)
	playerHandler := rest.NewPlayerRestHandler(playerService)

	router := mux.NewRouter()

	router.Handle("/v1/team", teamRestHandler.StoreTeam()).Methods(http.MethodPost)
	router.Handle("/v1/team", teamRestHandler.GetTeam()).Methods(http.MethodGet)

	router.Handle("/v1/player", playerHandler.StorePlayer()).Methods(http.MethodPost)
	router.Handle("/v1/player", playerHandler.GetPlayer()).Methods(http.MethodGet)

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
}
