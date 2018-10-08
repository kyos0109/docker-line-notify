package main

import (
	"log"
	"os"
	"net/http"
	"net/url"
	"strings"
	"strconv"
	"errors"
	"fmt"
)

const lineNotifyURL = "https://notify-api.line.me/api/notify"

type LineInfo struct {
	Token   string
	Message string
	Debug   bool
}

func main() {
	info := LineInfo{
			Token:   getToken("PLUGIN_TOKEN", "token_secret"),
			Message: os.Getenv("PLUGIN_MESSAGE"),
			Debug:   getBoolEnv("PLUGIN_DEBUG"),
		}

	if err := send(info) ; err != nil {
		fmt.Println(err.Error())
	}
}

func getToken(key ...string) string {
	for _, item := range key {
		if v, ok := os.LookupEnv(item); ok {
			return v
		}
	}
	return ""
}

func getBoolEnv(key string) bool {
    if v, ok := os.LookupEnv(key); ok {
    	if strings.ToLower(v) == "true" {
        	return true
    	}
    }
    return false
}

func send(l LineInfo) error {

	if l.Token == "" || l.Message == "" || len(l.Message) == 0 {
		return errors.New("error env: PLUGIN_TOKEN or PLUGIN_MESSAGE is empty.")
	}

	data := url.Values{}
	data.Add("message", l.Message)

	req, err := http.NewRequest(
		"POST",
		lineNotifyURL,
		strings.NewReader(data.Encode()),
	)

	if err != nil {
		return errors.New("error request : " + err.Error())
	}

	req.Header.Add("Authorization", "Bearer " + l.Token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return errors.New("error response: " + err.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		log.Println("send...OK")
	}

	if l.Debug {
		log.Println(resp)
	}

	return nil
}