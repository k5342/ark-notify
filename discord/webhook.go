package discord

import (
	"net/http"
	"bytes"
	"encoding/json"
	"github.com/bwmarrin/discordgo"
)

func SendWebhook(webhook_url string, msg string) {
	SendWebhookWithEmbed(webhook_url, msg, nil)
}

func SendWebhookWithEmbed(webhook_url string, msg string, embeds []*discordgo.MessageEmbed) {
        data := map[string]interface{}{
                "content": msg,
		"embeds": embeds,
        }
        json, err := json.Marshal(data)
        if err != nil {
		panic(err)
        }
        req, err := http.NewRequest("POST", webhook_url, bytes.NewBuffer(json))
        if err != nil {
		panic(err)
        }
        req.Header.Set("Content-Type", "application/json")
        client := http.Client{}
        _, err = client.Do(req)
        if err != nil {
		panic(err)
        }
}
