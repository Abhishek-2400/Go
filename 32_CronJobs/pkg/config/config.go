package config

import (
	"github.com/robfig/cron/v3"
)

func StartCorn() *cron.Cron {
	c := cron.New()
	return c
}
