package env

import "github.com/andypangaribuan/vision-go/vis"


/* ============================================
	Created by andy pangaribuan on 2021/05/20
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
var (
	ApiPort             int
	SchedulerDelay      int //in second
	SchedulerConfigFile string
	ActionLogFile       string
	ErrorLogFile        string
)


func Load() {
	ApiPort = vis.Env.GetIntEnv("API_PORT")
	SchedulerDelay = vis.Env.GetIntEnv("SCHEDULER_DELAY")
	SchedulerConfigFile = vis.Env.GetStr("SCHEDULER_CONFIG_FILE")
	ActionLogFile = vis.Env.GetStr("ACTION_LOG_FILE")
	ErrorLogFile = vis.Env.GetStr("ERROR_LOG_FILE")
}
