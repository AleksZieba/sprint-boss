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
			var userName string
			if i.GuildID != "" && i.Member != nil {
				userName = i.Member.User.Username
			} else {
				// For DM interactions, the user information is directly in i.User.
				userName = i.User.Username
			}

			err := dbQueries.StartSprint(context.WithValue(context.Background(), "username", userName)) //TODO the query function input
			if err != nil {
				log.Fatal("Sprint Start Failed") //turn this into a response
			}

			result := fmt.Sprintf("Sprint starts in %v minutes and will last %v minutes", num1, num2)
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("The sum is: %d", result),
				},
			})

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
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("The sum is: %d", result),
				},
			})
		}
	}
}
