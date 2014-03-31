package main

import (
	"fmt"
	"github.com/thoj/go-ircevent"
)

// http://waterpigs.co.uk/articles/a-simple-go-irc-bot/
func main() {
	var roomName = "#devcloud"
	con := irc.IRC("my-bot-test", "my-bot-test")
	err := con.Connect("irc.freenode.net:6667")
	if err != nil {
		fmt.Println("Failed connecting")
		return
	}
}
