package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"os/signal"
	"syscall"

	"github.com/jordan-wright/email"
)

var (
	GtmID = os.Getenv("MYSTIC_CASE_GTM_ID")
)

type ContextKey string
type Message struct {
	BoxName string
	Message []byte
}

func listenForEmail(emailChan chan Message) {
	host := os.Getenv("MYSTIC_CASE_SMTP_HOST")
	port := os.Getenv("MYSTIC_CASE_SMTP_PORT")
	username := os.Getenv("MYSTIC_CASE_USERNAME")
	password := os.Getenv("MYSTIC_CASE_PASSWORD")
	to := os.Getenv("MYSTIC_CASE_TO")
	from := os.Getenv("MYSTIC_CASE_FROM")

	log.Print("starting email worker")

	for message := range emailChan {
		e := email.Email{
			From:    from,
			To:      []string{to},
			Subject: fmt.Sprintf("User feedback on the %s box", message.BoxName),
			HTML:    message.Message,
		}

		log.Print("authorising on SMTP server...")
		auth := smtp.PlainAuth("", username, password, host)

		log.Print("sending email...")
		err := e.Send(fmt.Sprintf("%s:%s", host, port), auth)
		if err != nil {
			log.Printf("something went wrong: %s", err.Error())
		}
		// return smtp.SendMail(fmt.Sprintf("%s:%s", host, port), auth, from, []string{to}, message)
	}

	log.Print("Stopping email worker")
}

func main() {
	log.Print("starting server...")
	var emailChan = make(chan Message, 10)
	var ctx = context.WithValue(context.Background(), ContextKey("emailChan"), emailChan)

	initRoutes(ctx)

	go listenForEmail(emailChan)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	srv := http.Server{
		Addr: fmt.Sprintf(":%s", port),
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT)
		<-sigint

		// We received an interrupt signal, shut down.
		log.Print("Shutting down")
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(emailChan)
		close(idleConnsClosed)
	}()

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}

	<-idleConnsClosed
}
