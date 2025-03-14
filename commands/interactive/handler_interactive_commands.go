package interactive

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
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

			result := num1 + num2
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
