package interactive

import (
	//"context"
	//"database/sql"
	"fmt"
	"log"

	//"os"
	"sync"
	"time"

	//"github.com/AleksZieba/sprint-boss/internal/database"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	sprintStates = make(map[string]bool)
	mu           sync.RWMutex
)

// setSprintOn safely updates the sprintOn state.
func setSprintState(key string, state bool) error {
	mu.Lock()
	defer mu.Unlock()
	err := error(nil)
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("failed to set sprint state: %v", r)
		}
	}()

	sprintStates[key] = state
	return err
}

func getSprintState(key string) bool {
	mu.RLock()
	defer mu.RUnlock()
	return sprintStates[key]
}

func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	godotenv.Load()
	/*dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)
	if err != nil {
		log.Println("Failed To Load DB") //turn this into a response
	} */
	if i.Type == discordgo.InteractionApplicationCommand {
		switch i.ApplicationCommandData().Name {
		case "sprint":
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
			})
			if err != nil {
				log.Println("Failed to acknowledge interaction:", err)
				return
			}
			var sprintDelay float64 = 0.0 // default value if not provided
			var sprintTime float64 = 20.0 // default value if not provided

			for _, option := range i.ApplicationCommandData().Options {
				switch option.Name {
				case "start_delay_minutes":
					sprintDelay = option.FloatValue()
				case "sprint_time":
					sprintTime = option.FloatValue()
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
			_ = nickName //REMEMBER TO REMOVE THIS
			var userName string
			if i.GuildID != "" && i.Member != nil {
				userName = i.Member.User.Username
			} else {
				userName = i.User.Username
			}
			//ctx := context.WithValue(context.Background(), "username", userName)

			var serverName string
			if i.GuildID != "" {
				// This interaction is from a guild.
				// Try to get the guild from the session state first.
				guild, err := s.State.Guild(i.GuildID)
				if err != nil || guild == nil {
					// If it's not cached, fetch it directly via the API.
					guild, err = s.Guild(i.GuildID)
					if err != nil {
						log.Println("Error fetching guild info:", err)
						serverName = "Unknown Server"
					} else {
						serverName = guild.Name
					}
				} else {
					serverName = guild.Name
				}
			} else {
				// No GuildID means it's a direct message.
				serverName = "Direct Message"
			}

			// Now you have the server name in the variable "serverName".
			fmt.Printf("Command received from: %s\n", serverName)

			/*// You can now respond or process the command accordingly.
			queryArg := database.StartSprintParams{
				UserName:   userName,
				ServerName: serverName,
			}

			ctx = context.WithValue(ctx, "servername", serverName)
			err := dbQueries.StartSprint(ctx, queryArg)
			if err != nil {
				log.Println("Sprint Start Failed") //turn this into a response
			} */

			//var embed *discordgo.MessageEmbed //THIS CAN BE MOVED NOW
			if sprintDelay > 0 {
				_, err := s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
					Content: fmt.Sprintf("The sprint starts in %g minutes and will last %g minutes", sprintDelay, sprintTime),
				})
				if err != nil {
					log.Println("Failed to send channel message:", err)
				}
			} else {
				_, err := s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
					Content: fmt.Sprintf("The sprint starts now! It will last %g minutes", sprintTime),
				})
				if err != nil {
					log.Println("Failed to send channel message:", err)
				}
			}

			if serverName != "Direct Message" {
				err := setSprintState(serverName, true)
				if err != nil {
					log.Println("Failed to turn the sprint on/off") //turn this into a response
				}
			} else {
				err := setSprintState(userName+"/dmDMdm/", true)
				if err != nil {
					log.Println("Failed to turn the sprint on/off") //turn this into a response
				}
			}
			//TODO: ADD A CHECK HERE FOR IF THE SPRINT IS CANCELLED
			if sprintDelay > 0 {
				go func() {
					time.Sleep(time.Duration(sprintDelay) * time.Minute)
					_, err := s.ChannelMessageSend(i.ChannelID, fmt.Sprintf("The sprint starts now! It will last %g minutes", sprintTime))
					if err != nil {
						log.Println("Failed to send channel message:", err)
					}

					go func() {
						time.Sleep(time.Duration(sprintTime) * time.Minute)
						if serverName != "Direct Message" {
							err = setSprintState(serverName, false)
							if err != nil {
								log.Println("Failed to turn the sprint on/off") //turn this into a response
							}
						} else {
							err = setSprintState(userName+"/dmDMdm/", false)
							if err != nil {
								log.Println("Failed to turn the sprint on/off") //turn this into a response
							}
						}
						_, err := s.ChannelMessageSend(i.ChannelID, "The Sprint is over! Please use the /word command to send it your word counts!")
						if err != nil {
							log.Println("Failed to send channel message:", err)
						}
					}()
				}()
			} else {
				// Start a goroutine that will "wake up" (set SprintState to false) after the sprint duration.
				go func() {
					time.Sleep(time.Duration(sprintTime) * time.Minute)
					if serverName != "Direct Message" {
						err := setSprintState(serverName, false)
						if err != nil {
							log.Println("Failed to turn the sprint on/off") //turn this into a response
						}
					} else {
						err := setSprintState(userName+"/dmDMdm/", false)
						if err != nil {
							log.Println("Failed to turn the sprint on/off") //turn this into a response
						}
					}
					_, err := s.ChannelMessageSend(i.ChannelID, "The Sprint is over! Please use the /word command to send it your word counts!")
					if err != nil {
						log.Println("Failed to send channel message:", err)
					}
				}()
			}
		case "ready":
			var sprintDelay float64 = 0 // default value if not provided
			var sprintTime float64 = 20 // default value if not provided

			for _, option := range i.ApplicationCommandData().Options {
				switch option.Name {
				case "start_delay_minutes":
					sprintDelay = option.FloatValue()
				case "sprint_time":
					sprintTime = option.FloatValue()
				}
			}

			result := sprintDelay + sprintTime
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("The sum is: %g", result),
				},
			})
			if err != nil {
				log.Println("Failed To write response") //turn this into a response
			}
		case "cancel":
			// This command cancels the current sprint.
			var userName string
			if i.GuildID != "" && i.Member != nil {
				userName = i.Member.User.Username
			} else {
				userName = i.User.Username
			}

			var serverName string
			if i.GuildID != "" {
				// This interaction is from a guild.
				// Try to get the guild from the session state first.
				guild, err := s.State.Guild(i.GuildID)
				if err != nil || guild == nil {
					// If it's not cached, fetch it directly via the API.
					guild, err = s.Guild(i.GuildID)
					if err != nil {
						log.Println("Error fetching guild info:", err)
						serverName = "Unknown Server"
					} else {
						serverName = guild.Name
					}
				} else {
					serverName = guild.Name
				}
			} else {
				// No GuildID means it's a direct message.
				serverName = userName + "/dmDMdm/"
			}
			/*
					if serverName != "Direct Message" {
						err := setSprintState(serverName, false)
						if err != nil {
							log.Println("Failed to turn the sprint on/off") //turn this into a response
						}
					} else {
						err := setSprintState(userName+"/dmDMdm/", false)
					if err != nil {
						log.Println("Failed to turn the sprint on/off") //turn this into a response
					}
				}
			*/
			fmt.Printf("Command received from: %s\n", serverName)

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

			var embed *discordgo.MessageEmbed

			if getSprintState(serverName) {
				embed = &discordgo.MessageEmbed{
					Title:       "Sprint Cancelled!",
					Description: fmt.Sprintf("%s has cancelled the sprint. Feel free to start another sprint at any time.", nickName),
					Color:       0xFFA500, // Orange
				}
			} else {
				embed = &discordgo.MessageEmbed{
					Title:       "No Active Sprint",
					Description: fmt.Sprintf("%s are not in the middle of a sprint 😭 feel free to start one anytime 🥳", nickName),
					Color:       0xff0000, // Red
				}
			}
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{embed},
				},
			})
			if err != nil {
				log.Println("Failed To write response") //turn this into a response
			}
		case "test":
			_, err := s.ChannelMessageSend(i.ChannelID, "/test")
			if err != nil {
				log.Println("Failed to send channel message:", err)
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
