package interactive

import "github.com/bwmarrin/discordgo"

var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "sprint",
		Description: "Start a sprint",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "start_delay_minutes",
				Description: "number of minutes till the sprint starts",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "sprint_time",
				Description: "when your sprint starts and how long it is",
				Required:    false,
			},
		},
	},
	{
		Name:        "ready",
		Description: "Start a sprint",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "start_delay_minutes",
				Description: "number of minutes till the sprint starts",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "sprint_time",
				Description: "when your sprint starts and how long it is",
				Required:    false,
			},
		},
	},
}
