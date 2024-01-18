package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
)

func main() {
	setupConfig()
	err := sentry.Init(sentry.ClientOptions{
		Dsn: viper.GetString("SENTRY.DSN"),
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("failed to init sentry: %s", err.Error())
	}
	defer sentry.Flush(2 * time.Second)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := fmt.Errorf("simple error: %s", "something went wrong")
		sentry.CaptureException(err)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Error sample sent for Sentry!"))
	})

	port := 8080
	log.Printf("Server started on :%d", port)
	if err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}

func setupConfig() error {
	config := flag.String("config", "config.yaml", "load config file")
	flag.Parse()

	viper.SetConfigName(*config)
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
