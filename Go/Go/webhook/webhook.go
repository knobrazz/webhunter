package webhook

import (
    "bytes"
    "encoding/json"
    "net/http"
)

type WebhookClient struct {
    SlackURL    string
    DiscordURL  string
}

func (w *WebhookClient) SendToSlack(message map[string]interface{}) error {
    return w.send(w.SlackURL, message)
}

func (w *WebhookClient) SendToDiscord(message map[string]interface{}) error {
    return w.send(w.DiscordURL, message)
}

func (w *WebhookClient) send(url string, message map[string]interface{}) error {
    payload, err := json.Marshal(message)
    if err != nil {
        return err
    }

    resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    return nil
}

