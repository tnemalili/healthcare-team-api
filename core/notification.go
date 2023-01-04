package core

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func SMSHandler(to string, msg string) bool {

	headers := make(map[string]interface{})
	url := os.Getenv("SMS_URL")
	log.Infof("[notifications.SMSHandler] SMS URL: %v", url)
	log.Infof("[notifications.SMSHandler] SMS: %v", msg)
	body := make(map[string]interface{})
	body["message"] = msg
	body["to"] = to
	response, err := HTTPClient.REQUEST(url, "POST", headers, body)
	if err != nil {
		log.Errorf("[notifications.SMSHandler] %v", err)
		return false
	}
	return response.StatusCode != http.StatusBadRequest
}
