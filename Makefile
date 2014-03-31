all: bot

bot: bot.go
	go build bot.go

clean:
	rm bot
