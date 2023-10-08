package crontab

import "github.com/robfig/cron/v3"

type Interface interface {
	Rule() string
	Execute() func()
}

type TaskInterface interface {
	Tasks() Tasks
}

type Tasks []Interface

type Crontab struct {
	instance *cron.Cron
}

func New() *Crontab {
	c := &Crontab{
		// Enable "second" support,
		// doesn't impact "@every" schedule.
		instance: cron.New(cron.WithParser(cron.NewParser(
			cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
		))),
	}
	return c
}

func (c *Crontab) AddFunc(cmd ...Interface) {
	if len(cmd) == 0 {
		return
	}
	for _, job := range cmd {
		c.instance.AddJob(job.Rule(), cron.FuncJob(job.Execute()))
	}
}

func (c *Crontab) Start() {
	c.instance.Start()
}

func (c *Crontab) Stop() {
	// Stop the scheduler (does not stop any jobs already running).
	c.instance.Stop()
}
