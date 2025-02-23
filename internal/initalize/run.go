package initalize

import (
	"github.com/huuloc2026/go-backend/global"
)

func Run() {
	// Load configuration
	LoadConfig()
	//fmt.Println("Loading config...", global.Config.Database.Port)
	InitLogger()
	global.Logger.Info("Logger initialized")
	InitPostgresql()
	InitRedis()
	InitRouter()
	r := InitRouter()
	r.Run()
}
