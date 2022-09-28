package main

import (
	"context"
	"flag"
	"github.com/TimIrwing/nyashka-butler/internal/bot"
	"github.com/TimIrwing/nyashka-butler/internal/mongodb"
)

func main() {
	tPtr := flag.String("token", "", "token from @BotFather")
	uriPtr := flag.String("mongouri", mongodb.DefaultURI, "custom mongo uri, mainly used for auth")
	flag.Parse()

	db := mongodb.Init(context.TODO(), *uriPtr)
	bot.Start(*tPtr, db)
}
