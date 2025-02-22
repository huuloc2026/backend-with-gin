package initalize

import (
	"fmt"

	"github.com/huuloc2026/go-backend/global"
)

func Run() {
	// Load configuration
	LoadConfig()
	fmt.Println("Loading config...", global.Config.Security.JWT.Key)
	InitLogger()
	InitPostgresql()
	InitRedis()
	InitRouter()
	r := InitRouter()
	r.Run()
}
