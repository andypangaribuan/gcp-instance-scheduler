package main

import (
	"github.com/andypangaribuan/gcp-instance-scheduler/app/endpoint/ep-private"
	"github.com/andypangaribuan/gcp-instance-scheduler/app/env"
	"github.com/andypangaribuan/gcp-instance-scheduler/app/helper"
	"github.com/andypangaribuan/gcp-instance-scheduler/app/worker"
	"github.com/andypangaribuan/vision-go"
	"github.com/andypangaribuan/vision-go/core/api"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/20
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
func main() {
	vision.Initialize()
	env.Load()
	helper.InitializeFiles()

	go worker.Scheduler()

	e := api.BuildEcho(env.ApiPort, nil)
	endpoints(e)
	e.Serve()
}


func endpoints(e *api.EchoApi) {
	e.POST("/private/time", ep_private.Time)
}
