package helper

import (
	"context"
	"github.com/andypangaribuan/gcp-instance-scheduler/model"
	"github.com/pkg/errors"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
	sqladmin "google.golang.org/api/sqladmin/v1beta4"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/20
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
func StartStopGCE(isStart bool, config model.SchedulerConfigModel) error {
	ctx := context.Background()

	service, err := compute.NewService(ctx, option.WithCredentialsFile(config.ServiceAccountPath))
	if err != nil {
		return errors.WithStack(err)
	}

	instance, err := service.Instances.Get(config.ProjectId, config.Zone, config.Instance).Do()
	if err != nil {
		return errors.WithStack(err)
	}

	if instance.Status == "RUNNING" && isStart {
		return nil
	}
	if instance.Status == "TERMINATED" && !isStart {
		return nil
	}

	if isStart {
		_, err := service.Instances.Start(config.ProjectId, config.Zone, config.Instance).Do()
		if err != nil {
			return errors.WithStack(err)
		}
	} else {
		_, err := service.Instances.Stop(config.ProjectId, config.Zone, config.Instance).Do()
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}


func StartStopCloudSQL(isStart bool, config model.SchedulerConfigModel) error {
	ctx := context.Background()

	service, err := sqladmin.NewService(ctx, option.WithCredentialsFile(config.ServiceAccountPath))
	if err != nil {
		return errors.WithStack(err)
	}

	instance, err := service.Instances.Get(config.ProjectId, config.Instance).Do()
	if err != nil {
		return errors.WithStack(err)
	}

	if instance.Settings.ActivationPolicy == "ALWAYS" && isStart {
		return nil
	}
	if instance.Settings.ActivationPolicy == "NEVER" && !isStart {
		return nil
	}

	action := ""
	if isStart {
		action = "ALWAYS"
	} else {
		action = "NEVER"
	}

	setting := &sqladmin.DatabaseInstance{
		Settings: &sqladmin.Settings{
			ActivationPolicy: action,
		},
	}

	_, err = service.Instances.Patch(config.ProjectId, config.Instance, setting).Do()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

