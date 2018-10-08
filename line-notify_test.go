package main

import (
	"os"
	"testing"
)

const plugin_debug_key string = "PLUGIN_DEBUG"
const plugin_message_key string = "PLUGIN_MESSAGE"
const plugin_token_key string = "PLUGIN_TOKEN"
const token_secret_key string = "token_secret"

func TestSetDebugEnv(t *testing.T) {
	os.Setenv(plugin_debug_key, "true")

	if v := getBoolEnv(plugin_debug_key); v {
		t.Log("PASS")
	} else {
		t.Error("Fail, debug not true")
	}
}

func TestNullDebugEnv(t *testing.T) {
	os.Unsetenv(plugin_debug_key)

	v := getBoolEnv(plugin_debug_key)

	if v == false {
		t.Log("PASS")
	} else {
		t.Error("Fail, debug env not false")
	}
}

func TestNullSendInfo(t *testing.T) {
	info := LineInfo{}

	err := send(info)

	if err != nil {
		t.Log("PASS")
	} else {
		t.Error("Fail, error not found")
	}
}

func TestSend(t *testing.T) {
	info := LineInfo{
		Token:   "1234567890",
		Message: "Test",
	}

	err := send(info)

	if err == nil {
		t.Log("PASS")
	} else {
		t.Error("Fail, send error")
	}
}

func TestTokenSecret(t *testing.T) {
	os.Setenv(token_secret_key, "1234567890")

	if v := getToken(plugin_token_key, token_secret_key); v != "" {
		t.Log("PASS")
	} else {
		t.Error("Fail, token not found.")
	}
}

func TestTokenEnv(t *testing.T) {
	os.Setenv(plugin_token_key, "1234567890")

	if v := getToken(plugin_token_key, token_secret_key); v != "" {
		t.Log("PASS")
	} else {
		t.Error("Fail, token not found.")
	}
}
