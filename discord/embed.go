package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"ark-notify/event"
)

func GenerateEmbedFromEvent(ae *event.ArkEvent) (discordgo.MessageEmbed) {
	embed := discordgo.MessageEmbed{}
	switch ae.Kind {
	case event.KillEvent:
		embed.Title = "Killed"
		embed.Color = 0xb36f6f
	case event.TameEvent:
		embed.Title = "Tamed"
		embed.Color = 0x4db329
	case event.AdminCmdEvent:
		embed.Title = "AdminCmd"
		embed.Color = 0x828282
	case event.JoinEvent:
		embed.Title = "User Joined"
		embed.Color = 0x3496fe
	case event.LeaveEvent:
		embed.Title = "User Left"
		embed.Color = 0x8dacce
	case event.DefaultEvent:
		fallthrough
	default:
		embed.Title = "New Event"
		embed.Color = 0xdedede
	}
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
