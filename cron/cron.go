package cron

import (
	"97HW/service"
)

func UpdateDataAndNotify() error {
	if err := service.UpdateWeatherData(); err != nil {
		return err
	}

	if err := service.SendNotifications(); err != nil {
		return err
	}

	return nil
}
