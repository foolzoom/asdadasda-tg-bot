package main

import (
	"encoding/json"
	tele "gopkg.in/telebot.v3"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	go supportMonitoring()
	go listenPort()

	pref := getTgBotConfig()
	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}
	initTgMessageHandlers(b)
	log.Println("Starting...")
	b.Start()
}

func supportMonitoring() {
	handler := http.HandlerFunc(handleRequest)
	http.Handle("/monitor", handler)
}

func handleRequest(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Status OK"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	_, _ = w.Write(jsonResp)
	return
}

func listenPort() {
	port := os.Getenv("PORT")
	log.Println("Listening port: " + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}

func getTgBotConfig() tele.Settings {
	pref := tele.Settings{
		Token:  os.Getenv("ASDADASDA_TG_BOT_API_KEY"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	return pref
}

func initTgMessageHandlers(bot *tele.Bot) {
	bot.Handle("/ping", func(c tele.Context) error {
		return c.Send("Pong!")
	})
	bot.Handle("/pong", func(c tele.Context) error {
		return c.Send("Ping!")
	})
}
