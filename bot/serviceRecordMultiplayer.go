package bot

import (
	"fmt"
	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
	"toughcrab.com/halo/autocode"
	"toughcrab.com/halo/autocode/generated"
)

const route = "serviceRecord"
const description = "responds with serviceRecord"

type ServiceRecordMultiplayer struct {
	AutoCodeClient *autocode.Client
}

func (s *ServiceRecordMultiplayer) Handle(ctx *exrouter.Context) {
	// index 0 is the command the user ran
	discordInput := strings.Fields(ctx.Msg.Content)

	if len(discordInput) < 2 {
		_, _ = ctx.Reply(fmt.Sprintf("Please provide your gamertag i.e `!%s GamerTag`", route))
		return
	}

	gamertag := discordInput[1]

	// get the input gamertag's service record from the halo infinite stats API
	serviceRecord, err := s.AutoCodeClient.ServiceRecordMultiplayer(gamertag, "matchmade:pvp")
	if err != nil {
		_, _ = ctx.Reply(fmt.Sprintf("Couldn't fetch service record for : %s", gamertag))
		return
	}

	// embed the stats fields for the message
	fields := s.CreateEmbedFields(serviceRecord)

	// send an embedded message back to the channel it was requested from
	_, err = ctx.Ses.ChannelMessageSendEmbed(ctx.Msg.ChannelID, &discordgo.MessageEmbed{
		URL: fmt.Sprintf(
			"https://halotracker.com/halo-infinite/profile/xbl/%s/overview?experience=overall",
			gamertag),
		Title:       "Halo Multiplayer PvP Stats",
		Description: fmt.Sprintf("Stats for **%s**", gamertag),
		Color:       2067276,
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "Halo Stats for " + gamertag,
			IconURL: "https://cdnb.artstation.com/p/assets/images/images/011/374/989/large/ludwig-gaias-halo-infinite-v2-par-playbox.jpg?1529272697",
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL:      "https://cdnb.artstation.com/p/assets/images/images/011/374/989/large/ludwig-gaias-halo-infinite-v2-par-playbox.jpg?1529272697",
			ProxyURL: "",
			Width:    100,
			Height:   50,
		},
		Fields: fields,
	})

	if err != nil {
		log.Print("Something went wrong when sending the message to discord.", err)
	}
}

func (s *ServiceRecordMultiplayer) GetCommand() string {
	return route
}

func (s *ServiceRecordMultiplayer) GetDescription() string {
	return description
}

func (s *ServiceRecordMultiplayer) CreateEmbedFields(serviceRecord *generated.ServiceRecordMultiplayer) []*discordgo.MessageEmbedField {
	var embedFields []*discordgo.MessageEmbedField

	statFields := [][]string{
		{"K / D", fmt.Sprintf("%f", serviceRecord.Data.Core.Kdr)},
		{"Kills", fmt.Sprintf("%d", serviceRecord.Data.Core.Summary.Kills)},
		{"Deaths", fmt.Sprintf("%d", serviceRecord.Data.Core.Summary.Deaths)},
		{"Betrayals", fmt.Sprintf("%d", serviceRecord.Data.Core.Summary.Betrayals)},
		{"Assists", fmt.Sprintf("%d", serviceRecord.Data.Core.Summary.Assists)},
		{"Medals", fmt.Sprintf("%d", serviceRecord.Data.Core.Summary.Medals)},
		{"Destroyed Vehicles", fmt.Sprintf("%d", serviceRecord.Data.Core.Summary.Vehicles.Destroys)},
		{"Hijacked Vehicles", fmt.Sprintf("%d", serviceRecord.Data.Core.Summary.Vehicles.Hijacks)},
	}

	for _, ch := range statFields {
		embedFields = append(embedFields, &discordgo.MessageEmbedField{
			Name:   ch[0],
			Value:  ch[1],
			Inline: false,
		})
	}

	return embedFields

}

func NewServiceRecordMultiplayer(client *autocode.Client) *ServiceRecordMultiplayer {
	return &ServiceRecordMultiplayer{
		AutoCodeClient: client,
	}
}
