// Copyright © 2018 Govinda Attal govinda.attal@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"

	"github.com/govinda-attal/user-auth/handler"
	"github.com/govinda-attal/user-auth/internal/provider"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the user authentication micro service",
	Run:   startServer,
}

func startServer(cmd *cobra.Command, args []string) {
	provider.Setup()
	r := mux.NewRouter()
	s := r.PathPrefix("/users/v1").Subrouter().StrictSlash(true)
	s.HandleFunc("/auth",
		handler.ErrorHandler(handler.AuthenticateUser)).Methods("POST")
	s.HandleFunc("/verify",
		handler.ErrorHandler(handler.VerifyUser)).Methods("POST")
	s.HandleFunc("/register",
		handler.ErrorHandler(handler.RegisterUser)).Methods("POST")
	s.HandleFunc("/confirm",
		handler.ErrorHandler(handler.ConfirmUser)).Methods("POST")
	s.HandleFunc("/info",
		handler.ValidateUserLogon(handler.ErrorHandler(handler.FetchUserInfo))).Methods("GET")

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
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
	log.Println("shutting down ...")
	provider.Cleanup()
	os.Exit(0)
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
