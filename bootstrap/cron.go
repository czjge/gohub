package bootstrap

import (
	tasks "github.com/czjge/gohub/app/tasks"
	"github.com/czjge/gohub/pkg/crontab"
)

func SetupCron() {

	go func() {

		cron := crontab.New()

		cron.AddFunc(tasks.New().Tasks()...)

		cron.Start()

		defer cron.Stop()

		// Block the goroutine.
		select {}
	}()
}
