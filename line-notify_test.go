package main

import (
	"os"
	"testing"
)

func TestSetDebugEnv(t *testing.T) {
	os.Setenv("PLUGIN_DEBUG", "true")

	if v := getBoolEnv("PLUGIN_DEBUG"); v {
		t.Log("PASS")
	} else {
		t.Error("Fail, debug not true")
	}
}

func TestNullDebugEnv(t *testing.T) {
	os.Unsetenv("PLUGIN_DEBUG")

	v := getBoolEnv("PLUGIN_DEBUG")

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
		Token: "1234567890",
		Message: "Test",
	}

	err := send(info)

	if err == nil {
		t.Log("PASS")
	} else {
		t.Error("Fail, send error")
	}
}