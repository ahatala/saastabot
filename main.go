package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: saastabot slack-bot-token\n")
		os.Exit(1)
	}

	start := time.Now()
	ws, id := slackConnect(os.Args[1])
	fmt.Println(id + " ready, ^C exits")

	for {
		m, err := getMessage(ws)
		if err != nil {
			log.Println(err)
		}

		if m.Type == "message" && strings.HasPrefix(m.Text, "<@"+id+">") {
			parts := strings.Fields(m.Text)
			if len(parts) == 2 && parts[1] == "report" {
				go func(m Message) {
					hostname, _ := os.Hostname()
					m.Text = fmt.Sprintf("Been running %s on %s\n", time.Since(start), hostname)
					postMessage(ws, m)
				}(m)
			}
		}
	}
}
