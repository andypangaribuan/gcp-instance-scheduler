package worker

import (
	"fmt"
	"github.com/andypangaribuan/gcp-instance-scheduler/app/env"
	"github.com/andypangaribuan/gcp-instance-scheduler/app/helper"
	"github.com/andypangaribuan/gcp-instance-scheduler/model"
	"github.com/andypangaribuan/vision-go/vis"
	"os"
	"strings"
	"time"
)

/* ============================================
	Created by andy pangaribuan on 2021/05/20
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
func Scheduler() {
	for {
		configs := helper.GetSchedulerConfig()
		day := helper.GetDayNow()
		tm := helper.GetTimeNow()

		var (
			idxStarts = make([]int, 0)
			idxStops  = make([]int, 0)
		)


		for idx, conf := range configs {
			if conf.Active {
				if _, ok := conf.Days[day]; ok {
					if conf.StartTime == tm {
						idxStarts = append(idxStarts, idx)
					}
					if conf.StopTime == tm {
						idxStops = append(idxStops, idx)
					}
				}
			}
		}

		if len(idxStarts) > 0 || len(idxStops) > 0 {
			go doStartStop(configs, idxStarts, idxStops)
			for {
				if tm != helper.GetTimeNow() {
					break
				}
				time.Sleep(time.Second)
			}
		}

		time.Sleep(time.Second * time.Duration(env.SchedulerDelay))
	}
}


func doStartStop(configs []model.SchedulerConfigModel, idxStarts, idxStops []int) {
	actionLogFile := env.ActionLogFile
	errorLogFile := env.ErrorLogFile

	tmNow := func() string {
		return vis.Convert.TimeToStrFull(time.Now())
	}

	appendToFile := func(txt, filePath string) {
		f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return
		}
		defer f.Close()

		_, _ = f.WriteString(txt)
	}

	isFirstLog := true
	printLog := func(txt string) {
		if env.PrintLog {
			if isFirstLog {
				isFirstLog = false
				fmt.Printf("\n\n")
			}

			fmt.Printf(txt)
		}
	}


	for _, idx := range idxStarts {
		config := configs[idx]
		switch strings.ToLower(config.Type) {
		case "vm":
			if err := helper.StartStopGCE(true, config); err == nil {
				msg := fmt.Sprintf("%v starting vm success, %v", tmNow(), config.Name)
				msg += "\n"
				appendToFile(msg, actionLogFile)
				printLog(msg)
			} else {
				msg := fmt.Sprintf("%v starting vm failed, %v\n", tmNow(), config.Name)
				msg += *vis.Log.Stack(err)
				msg += "\n\n"
				appendToFile(msg, errorLogFile)
				printLog(msg)
			}

		case "sql":
			if err := helper.StartStopCloudSQL(true, config); err == nil {
				msg := fmt.Sprintf("%v starting cloud sql success, %v", tmNow(), config.Name)
				msg += "\n"
				appendToFile(msg, actionLogFile)
				printLog(msg)
			} else {
				msg := fmt.Sprintf("%v starting cloud sql failed, %v\n", tmNow(), config.Name)
				msg += *vis.Log.Stack(err)
				msg += "\n\n"
				appendToFile(msg, errorLogFile)
				printLog(msg)
			}
		}
	}


	for _, idx := range idxStops {
		config := configs[idx]
		switch strings.ToLower(config.Type) {
		case "vm":
			if err := helper.StartStopGCE(false, config); err == nil {
				msg := fmt.Sprintf("%v stopping vm success, %v", tmNow(), config.Name)
				msg += "\n"
				appendToFile(msg, actionLogFile)
				printLog(msg)
			} else {
				msg := fmt.Sprintf("%v stopping vm failed, %v\n", tmNow(), config.Name)
				msg += *vis.Log.Stack(err)
				msg += "\n\n"
				appendToFile(msg, errorLogFile)
				printLog(msg)
			}

		case "sql":
			if err := helper.StartStopCloudSQL(false, config); err == nil {
				msg := fmt.Sprintf("%v stopping cloud sql success, %v", tmNow(), config.Name)
				msg += "\n"
				appendToFile(msg, actionLogFile)
				printLog(msg)
			} else {
				msg := fmt.Sprintf("%v stopping cloud sql failed, %v\n", tmNow(), config.Name)
				msg += *vis.Log.Stack(err)
				msg += "\n\n"
				appendToFile(msg, errorLogFile)
				printLog(msg)
			}
		}
	}
}
