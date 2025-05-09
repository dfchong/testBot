package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const token = "YOUR_TELEGRAM_BOT_TOKEN"

type Update struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

func setWebhook(webhookURL string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/setWebhook", token)
	resp, err := http.PostForm(url, map[string][]string{
		"url": {webhookURL},
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Println("Set webhook response status:", resp.Status)
	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	var update Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Println("Failed to decode update:", err)
		return
	}

	log.Printf("Received message: %s\n", update.Message.Text)
	sendMessage(update.Message.Chat.ID, "You said: "+update.Message.Text)
}

func sendMessage(chatID int64, text string) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	payload := map[string]interface{}{
		"chat_id": chatID,
		"text":    text,
	}
	data, _ := json.Marshal(payload)
	_, err := http.Post(url, "application/json", 
	                     http.NoBodyFromReader(data))
	if err != nil {
		log.Println("Error sending message:", err)
	}
}

func main() {
	webhookURL := "https://yourdomain.com" // same as Caddy domain
	if err := setWebhook(webhookURL); err != nil {
		log.Fatal("Failed to set webhook:", err)
	}

	http.HandleFunc("/", handler)
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
