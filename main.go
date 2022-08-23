package main

import (
	"context"
	"flag"
	"log"

	tgClient "github.com/arturgumerov/linkbot/clients/telegram"
	eventconsumer "github.com/arturgumerov/linkbot/events/consumer/event-consumer"
	"github.com/arturgumerov/linkbot/events/telegram"
	"github.com/arturgumerov/linkbot/storage/sqlite"
)

const (
	tgBotHost         = "api.telegram.org"
	sqliteStoragePath = "data/sqlite/storage.db"
	batchSize         = 100
)

func main() {

	//token = flags.Get(token)

	//tgClient = telegram.New(token)

	//s:= files.New(storagePath)
	s, err := sqlite.New(sqliteStoragePath)
	if err != nil {
		log.Fatal("can't connect to storage: ", err)
	}

	if err := s.Init(context.TODO()); err != nil {
		log.Fatal("can't init storage: ", err)
	}

	eventsProcessor := telegram.New(tgClient.New(tgBotHost, mustToken()), s)
	log.Print("service started")
	//fetcher = fetcher.New(tgClient)
	consumer := eventconsumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal()
	}

	//processor = processor.New(tgClient)

	//consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
