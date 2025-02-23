package global

import (
	"github.com/huuloc2026/go-backend/pkg/logger"
	"github.com/huuloc2026/go-backend/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.Logger
	Mdb    *gorm.DB
)
