package global

import (
	"github.com/huuloc2026/go-backend/pkg/logger"
	"github.com/huuloc2026/go-backend/pkg/setting"
	"github.com/redis/go-redis/v9"

	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.Logger
	Rdb    *redis.Client
	Mdb    *gorm.DB
)
