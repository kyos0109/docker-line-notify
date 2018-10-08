package main

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"text/template"
)

const lineNotifyURL = "https://notify-api.line.me/api/notify"

type LineInfo struct {
	Token   string
	Message string
	Debug   bool
}

type ProjectStatus struct {
	BuildStatus string
	RepoName	string
	RepoBranch  string
	BuildNum    string
	CommitID    string
	Author      string
	CommitMsg   string
}

func main() {

	for _, pair := range os.Environ() {
		log.Println(pair)
	}

	info := LineInfo{
		Token:   getToken("PLUGIN_TOKEN", "TOKEN_SECRET"),
		Message: getMessage("PLUGIN_MESSAGE"),
		Debug:   getBoolEnv("PLUGIN_DEBUG"),
	}

	if err := send(info); err != nil {
		log.Fatal(err.Error())
	}
}

func getMessage(msg string) string {

	repo := ProjectStatus{
		BuildStatus: os.Getenv("CI_BUILD_STATUS"),
		RepoName:    os.Getenv("CI_REPO_NAME"),
		RepoBranch:  os.Getenv("DRONE_COMMIT_BRANCH"),
		BuildNum:    os.Getenv("DRONE_BUILD_NUMBER"),
		CommitID:    os.Getenv("DRONE_COMMIT_SHA"),
		Author:      os.Getenv("DRONE_COMMIT_AUTHOR"),
		CommitMsg:   os.Getenv("DRONE_COMMIT_MESSAGE"),
	}

	if v, ok := os.LookupEnv(msg); ok {
		t := template.New("drone message")
		t, err := t.Parse(v)

		if err != nil {
			log.Fatal("Parse:", err)
			return ""
		}

		var tpl bytes.Buffer
		if err := t.Execute(&tpl, repo); err != nil {
			log.Fatal("Execute:", err)
			return ""
		}
		return tpl.String()
	}
	return ""
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

	req.Header.Add("Authorization", "Bearer "+l.Token)
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
