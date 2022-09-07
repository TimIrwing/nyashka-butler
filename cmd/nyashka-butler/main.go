package main

import (
	"context"
	"github.com/TimIrwing/nyashka-butler/internal/bot"
	"github.com/TimIrwing/nyashka-butler/internal/global"
	"github.com/TimIrwing/nyashka-butler/internal/mongodb"
	"log"
	"os"
	"strings"
)

const appName = "NyashkaButlerBot"

func main() {
	ctx := context.TODO()
	g := global.New(ctx, mongodb.New(ctx, appName))
	bot.New(getToken()).HandleUpdates(g)
}

func getToken() string {
	t, err := os.ReadFile("token")
	if err != nil {
		log.Panicf("Error getting token from file: %s\n", err)
	}
	return strings.Trim(string(t), "\r\n ")
}
