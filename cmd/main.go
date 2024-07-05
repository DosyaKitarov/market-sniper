package main

import (
	"context"
	"crypto/tls"
	"flag"
	"github.com/DosyaKitarov/market-sniper/internal/app/database"
	"github.com/DosyaKitarov/market-sniper/internal/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type application struct {
	logger *slog.Logger
	client *mongo.Client
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	client := database.ConnectToDb()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()

	app := &application{
		client: &client}

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:         *addr,
		Handler:      app.routes(),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: time.Minute,
	}

	logger.Info(ctx, "Starting server on ", *addr)

	err := srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")

	logger.Error(ctx, "server.ListenAndServeTLS() failed: %v", err)
	os.Exit(1)
}
