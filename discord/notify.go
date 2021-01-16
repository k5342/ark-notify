package discord

import (
	"github.com/bwmarrin/discordgo"
	"ark-notify/event"
)

func NotifyEvent(webhookURL string, ae *event.ArkEvent) {
	emb := GenerateEmbedFromEvent(ae)
	embeds := []*discordgo.MessageEmbed{ &emb }
	SendWebhookWithEmbed(webhookURL, "", embeds)
}
