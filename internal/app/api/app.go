package api

import (
	"github.com/sirupsen/logrus"
	log "github.com/status-report/internal/app/utils/log"
)

// NewApp returns configured app
func NewApp(level logrus.Level, name string, serviceDetails interface{}) *App {
	return &App{
		log: log.ServiceLogger(level, name, serviceDetails),
	}
}
