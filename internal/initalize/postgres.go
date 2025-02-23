package initalize

import (
	"fmt"
	"time"

	"github.com/huuloc2026/go-backend/global"
	"github.com/huuloc2026/go-backend/internal/po"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgresql() {
	m := global.Config.Database
	dsn := "host=%s user=%s  password=%s  dbname=%s  port=%v sslmode=disable search_path=public"
	var s = fmt.Sprintf(dsn, m.Host, m.User, m.Password, m.DBName, m.Port)
	db, err := gorm.Open(postgres.Open(s), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	global.Logger.Info("Postgresql connected")
	global.Mdb = db
	// set pool
	setPool(global.Mdb)
	// Tạo schema nếu chưa có
	db.Exec("CREATE SCHEMA IF NOT EXISTS public;")
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	migrateTables(global.Mdb)

}

func setPool(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to get database instance")
	}

	sqlDB.SetMaxOpenConns(50)                  // Limit the maximum number of open connections
	sqlDB.SetMaxIdleConns(10)                  // Set the number of idle connections
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // Limit how long a connection can be reused
	sqlDB.SetConnMaxIdleTime(10 * time.Minute) // Set the max idle time before closing connection
}

func migrateTables(db *gorm.DB) {
	err := db.AutoMigrate(
		&po.User{},
	// &models.Product{},
	// &models.Order{},
	)
	if err != nil {
		panic("Failed to migrate database tables")
	}
}
