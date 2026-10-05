package main

import (
	"fmt"
	"log"
	"net/http"
)

var appConfig Config

func main() {
	var err error
	appConfig, err = LoadConfig("config.toml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Read Secrets File Successfully.\n-------------------------------")
	fmt.Printf("Secret 1: %s\n", appConfig.Secrets.GiteaWebhookSecret)
	fmt.Printf("Secret 2: %s\n", appConfig.Secrets.GitlabWebhookSecret)
	fmt.Printf("IP Addr : %s\n", appConfig.Server.ServerIP)
	fmt.Printf("Port    : %d\n", appConfig.Server.ServerPort)

	log.Println("Checking terminal size.\n-------------------------------")
	CalculateStarPositions()

	fmt.Println("Starting Webhook Server.\n-------------------------------")
	// http.HandleFunc("/webhook", webhookHandler)
	log.Println("Server running on: 9999")
	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatal(err)
	}
}
