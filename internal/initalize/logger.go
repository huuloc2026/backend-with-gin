package initalize

import (
	"github.com/huuloc2026/go-backend/global"
	"github.com/huuloc2026/go-backend/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
