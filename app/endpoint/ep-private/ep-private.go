package ep_private

import (
	"fmt"
	"github.com/andypangaribuan/gcp-instance-scheduler/app/env"
	"github.com/andypangaribuan/vision-go/core/api"
	"github.com/andypangaribuan/vision-go/vis"
	"net/http"
	"strconv"
	"time"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/20
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
func Time(c api.Context) error {
	return c.ResponseStr(http.StatusOK, vis.Convert.TimeToStrFull(time.Now()))
}


func Day(c api.Context) error {
	weekday := time.Now().Weekday()
	day := int(weekday)
	return c.ResponseStr(http.StatusOK, strconv.Itoa(day))
}


func ClearConsole(c api.Context) error {
	vis.Term.ClearScreen()
	return c.ResponseStr(http.StatusOK, "done")
}


func LogStatus(c api.Context) error {
	msg := "print: "
	if env.PrintLog {
		msg += "active"
	} else {
		msg += "inactive"
	}
	return c.ResponseStr(http.StatusOK, msg)
}


func ReverseLogStatus(c api.Context) error {
	msg := "old status: %v, new status: %v"
	if env.PrintLog {
		msg = fmt.Sprintf(msg, "active", "inactive")
	} else {
		msg = fmt.Sprintf(msg, "inactive", "active")
	}
	env.PrintLog = !env.PrintLog
	return c.ResponseStr(http.StatusOK, msg)
}
