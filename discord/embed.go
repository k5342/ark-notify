package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"ark-notify/event"
)

func GenerateEmbedFromEvent(ae *event.ArkEvent) (discordgo.MessageEmbed) {
	embed := discordgo.MessageEmbed{}
	embed.Title = ae.GetEventTitle()
	embed.Color = ae.GetColor()
	embed.Timestamp = ae.Timestamp.Format("2006-01-02T15:04:05-07:00")
	var fields []*discordgo.MessageEmbedField
	for k, v := range ae.Info {
		f := discordgo.MessageEmbedField{}
		f.Name = k
		f.Value = v
		f.Inline = true
		fields = append(fields, &f)
	}
	if len(fields) == 0 {
		f := discordgo.MessageEmbedField{}
		f.Name = "RawLog"
		f.Value = fmt.Sprintf("```%s```", ae.RawLog)
		fields = append(fields, &f)
	}
	embed.Fields = fields
	return embed
}
