package webhook

import (
	"net/http"
	"bytes"
	"encoding/json"
)

func SendWebhook(webhook_url string, msg string) {
        data := map[string]interface{}{
                "content": msg,
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

