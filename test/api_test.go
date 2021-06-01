package test

import (
	"github.com/andypangaribuan/vision-go"
	"github.com/andypangaribuan/vision-go/vis"
	"net/http"
	"testing"
)


/* ============================================
	Created by andy pangaribuan on 2021/06/01
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
func init() {
	vision.Initialize()
}


func postCall(t *testing.T, path string, header map[string]string, body map[string]interface{}) {
	url := "http://localhost:33301" + path
	httpData, httpCode, err := vis.Http.Post(url, header, body, false, nil)
	logFatal(t, err)

	if httpCode == http.StatusOK {
		t.Logf("success, response: %v", string(httpData))
	} else {
		t.Logf("failed: code: %v, response: %v", httpCode, string(httpData))
		t.FailNow()
	}
}


func TestApi_Time(t *testing.T) {
	postCall(t, "/private/time", nil, nil)
}

func TestApi_Day(t *testing.T) {
	postCall(t, "/private/day", nil, nil)
}

func TestApi_ClearConsole(t *testing.T) {
	postCall(t, "/private/clear-console", nil, nil)
}

func TestApi_LogStatus(t *testing.T) {
	postCall(t, "/private/log-status", nil, nil)
}

func TestApi_ReverseLogStatus(t *testing.T) {
	postCall(t, "/private/reverse-log-status", nil, nil)
}
