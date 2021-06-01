package helper

import (
	"github.com/andypangaribuan/gcp-instance-scheduler/app/env"
	"github.com/andypangaribuan/gcp-instance-scheduler/app/store"
	"github.com/andypangaribuan/gcp-instance-scheduler/model"
	"github.com/andypangaribuan/vision-go/vis"
	"github.com/pkg/errors"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/20
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
var mutex sync.Mutex

func schedulerConfigHeader() string {
	header := ""
	for i, m := range store.MapSchedulerConfig {
		if i != 0 {
			header += "\t"
		}
		header += m.Name
	}
	return header
}


func lineToSchedulerConfigModel(line string) model.SchedulerConfigModel {
	isActive := func(txt string) bool {
		if txt == "1" {
			return true
		}
		return false
	}

	toDays := func(txt string) map[int]interface{} {
		arr := strings.Split(txt, ",")
		m := make(map[int]interface{}, 0)

		for _, a := range arr {
			d, err := strconv.Atoi(a)
			if err != nil {
				log.Fatalf(":: ERROR\n%v\n", errors.WithStack(err))
			}
			m[d] = nil
		}

		return m
	}

	arr := strings.Split(line, "\t")
	return model.SchedulerConfigModel{
		Active:             isActive(arr[0]),
		Name:               strings.TrimSpace(arr[1]),
		ProjectId:          strings.TrimSpace(arr[2]),
		Zone:               strings.TrimSpace(arr[3]),
		Instance:           strings.TrimSpace(arr[4]),
		Type:               strings.TrimSpace(arr[5]),
		StartTime:          strings.TrimSpace(arr[6]),
		StopTime:           strings.TrimSpace(arr[7]),
		Days:               toDays(strings.TrimSpace(arr[8])),
		ServiceAccountPath: strings.TrimSpace(arr[9]),
		Description:        strings.TrimSpace(arr[10]),
	}
}


func InitializeFiles() {
	filePaths := []string{env.ActionLogFile, env.ErrorLogFile}
	for _, filePath := range filePaths {
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			f, err := os.Create(filePath)
			if err != nil {
				log.Fatalf(":: ERROR\n%v\n", errors.WithStack(err))
			}
			f.Close()
		}
	}

	filePath := env.SchedulerConfigFile
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		f, err := os.Create(filePath)
		if err != nil {
			log.Fatalf(":: ERROR\n%v\n", errors.WithStack(err))
		}
		defer f.Close()

		_, err = f.WriteString(schedulerConfigHeader())
		if err != nil {
			log.Fatalf(":: ERROR\n%v\n", errors.WithStack(err))
		}
	}
}


func GetSchedulerConfig() (configs []model.SchedulerConfigModel) {
	mutex.Lock()
	defer mutex.Unlock()

	configs = make([]model.SchedulerConfigModel, 0)

	scanner, err := vis.Util.ScanFileLines(env.SchedulerConfigFile)
	if err != nil {
		log.Fatalf(":: ERROR\n%v\n", errors.WithStack(err))
	}

	for fn := range scanner {
		_, index, line := fn()
		if index == 0 || strings.TrimSpace(line) == "" {
			continue
		}

		configs = append(configs, lineToSchedulerConfigModel(line))
	}

	return
}

// Sunday:0, Monday:1 ...
func GetDayNow() int {
	weekday := time.Now().Weekday()
	return int(weekday)
}

func GetTimeNow() string {
	return vis.Convert.TimeToStr(time.Now(), "HH:mm")
}
