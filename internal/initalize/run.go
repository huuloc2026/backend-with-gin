package initalize

import (
	"fmt"

	"github.com/huuloc2026/go-backend/global"
	"go.uber.org/zap"
)

func Run() {
	// Load configuration
	LoadConfig()
	fmt.Println("Loading config...", global.Config.Security.JWT.Key)
	InitLogger()
	global.Logger.Info("Logger initialized", zap.String("author", "Huuloc"))
	InitPostgresql()
	InitRedis()
	InitRouter()
	r := InitRouter()
	r.Run()
}
