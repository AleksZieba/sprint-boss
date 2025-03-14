package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/AleksZieba/sprint-boss/commands"
	"github.com/AleksZieba/sprint-boss/commands/interactive"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	token := os.Getenv("BOT_TOKEN")
	sess, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	sess.AddHandler(interactive.InteractionCreate)
	sess.AddHandler(commands.HandlerStartSprint)
	/*sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID || !strings.HasPrefix(m.Content, "/") {
			return
		}

		if m.Content == "/sprint" {
			s.ChannelMessageSend(m.ChannelID, "How many minutes do you want your sprint to last?")
		}
	}) */

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	fmt.Println("Bot server is online.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
