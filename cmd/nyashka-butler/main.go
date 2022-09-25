package main

import (
	"context"
	"github.com/TimIrwing/nyashka-butler/internal/bot"
	"github.com/TimIrwing/nyashka-butler/internal/mongodb"
	"log"
	"os"
	"strings"
)

func main() {
	db := mongodb.Init(context.TODO())
	bot.Start(getToken(), db)
}

func getToken() string {
	t, err := os.ReadFile("token")
	if err != nil {
		log.Fatalf("Error getting token from file: %s\n", err)
	}
	return strings.Trim(string(t), "\r\n ")
}
