package main

import (
	"ark-notify/discord"
	"ark-notify/event"
)

func (an *ArkNotifier) NotifyEvent(ae *event.ArkEvent) {
	discord.NotifyEvent(an.WebhookURL, ae)
}
