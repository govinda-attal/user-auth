package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"github.com/govinda-attal/user-auth/handler"
	"github.com/govinda-attal/user-auth/handler/mw/usrtoken"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/auth", handler.AuthenticateUser).Methods("POST")
	r.HandleFunc("/verify", handler.VerifyUser).Methods("POST")
	r.HandleFunc("/register", handler.RegisterUser).Methods("POST")
	r.HandleFunc("/confirm", handler.ConfirmUser).Methods("POST")

	n := negroni.New()
	n.Use(negroni.HandlerFunc(usrtoken.ValidateUserLogon))
	n.UseHandler(r)

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      n,
	}
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	signal.Notify(c, os.Interrupt)
	// Block until we receive our signal.
	<-c
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait until the timeout deadline.
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
