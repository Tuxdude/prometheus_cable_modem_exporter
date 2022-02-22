package main

import (
	"github.com/tuxdude/zzzlog"
	"github.com/tuxdude/zzzlogi"
)

var (
	log = buildLogger()
)

func buildLogger() zzzlogi.Logger {
	config := zzzlog.NewConsoleLoggerConfig()
	config.MaxLevel = zzzlog.LvlDebug
	return zzzlog.NewLogger(config)
}
