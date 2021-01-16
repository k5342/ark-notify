package main

import (
	"ark-notify/webhook"
	"github.com/bwmarrin/discordgo"
)

func (an *ArkNotifier) NotifyEvent(ae *ArkEvent) {
	emb := GenerateEmbedFromEvent(ae)
	embeds := []*discordgo.MessageEmbed{ &emb }
	webhook.SendWebhookWithEmbed(an.WebhookURL, "", embeds)
}
