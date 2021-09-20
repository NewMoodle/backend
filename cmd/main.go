package main

import (
	"github.com/ZhansultanS/myLMS/backend/internal/app"
	"github.com/ZhansultanS/myLMS/backend/internal/config"
)

func main() {
	app.Start(config.EnvLocal)
}
