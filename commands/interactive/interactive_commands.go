package interactive

import "github.com/bwmarrin/discordgo"

var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "sprint",
		Description: "Start a sprint",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "tillStart",
				Description: "number of minutes till the sprint starts",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "sprintTime",
				Description: "when your sprint starts and how long it is",
				Required:    true,
			},
		},
	},
	//{}
}
