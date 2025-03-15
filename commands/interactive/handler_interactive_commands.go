package interactive

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/AleksZieba/sprint-boss/internal/database"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)
	if err != nil {
		log.Fatal("Failed To Load DB") //turn this into a response
	}
	if i.Type == discordgo.InteractionApplicationCommand {
		switch i.ApplicationCommandData().Name {
		case "sprint":
			var num1 int64 = 0  // default value if not provided
			var num2 int64 = 20 // default value if not provided

			for _, option := range i.ApplicationCommandData().Options {
				switch option.Name {
				case "start_delay_minutes":
					num1 = option.IntValue()
				case "sprint_time":
					num2 = option.IntValue()
				}
			}
			// func WithValue(parent Context, key, val any) Context - use with context.Background() as "Context"
			// If the interaction occurred in a guild, use the Member.User field.

			var nickName string
			if i.GuildID != "" && i.Member != nil {
				if i.Member.Nick != "" {
					nickName = i.Member.Nick
				} else {
					nickName = i.Member.User.Username
				}
			} else {
				nickName = i.User.Username
			}

			var userName string
			if i.GuildID != "" && i.Member != nil {
				userName = i.Member.User.Username
			} else {
				userName = i.User.Username
			}
			err := dbQueries.StartSprint(context.WithValue(context.Background(), "username", userName)) //TODO the query function input
			if err != nil {
				log.Fatal("Sprint Start Failed") //turn this into a response
			}
			var embed *discordgo.MessageEmbed
			if num1 > 0 {
				embed = &discordgo.MessageEmbed{
					Title:       "Sprint Set",
					Description: fmt.Sprintf("%s's Sprint starts in %d minutes and will last %d minutes", nickName, num1, num2),
					Color:       0x00ff00, // Green
				}
			} else {
				embed = &discordgo.MessageEmbed{
					Title:       "Sprint Starts Now",
					Description: fmt.Sprintf("%s, Your sprint starts now and will last %d minutes", nickName, num2),
					Color:       0x00ff00, // Green
				}
			}
			// 0xff0000 Red
			//result := fmt.Sprintf("Sprint starts in %v minutes and will last %v minutes", num1, num2)
			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{embed},
				},
			})
			if err != nil {
				log.Fatal("Failed To write response") //turn this into a response
			}

		case "ready":
			var num1 int64 = 0  // default value if not provided
			var num2 int64 = 20 // default value if not provided

			for _, option := range i.ApplicationCommandData().Options {
				switch option.Name {
				case "start_delay_minutes":
					num1 = option.IntValue()
				case "sprint_time":
					num2 = option.IntValue()
				}
			}

			result := num1 + num2
			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("The sum is: %d", result),
				},
			})
			if err != nil {
				log.Fatal("Failed To write response") //turn this into a response
			}
		}
	}
}

/* func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
    var displayName string

    // Check if the message is from a guild and if a nickname is set.
    if m.GuildID != "" && m.Member != nil && m.Member.Nick != "" {
        displayName = m.Member.Nick
    } else {
        displayName = m.Author.Username
    }

    fmt.Printf("Display name: %s\n", displayName)
} */
