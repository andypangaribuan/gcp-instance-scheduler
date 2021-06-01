package model

/* ============================================
	Created by andy pangaribuan on 2021/05/20
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
type SchedulerConfigModel struct {
	Active             bool
	Name               string
	ProjectId          string
	Zone               string
	Instance           string
	Type               string
	StartTime          string
	StopTime           string
	Days               map[int]interface{}
	ServiceAccountPath string
	Description        string
}

