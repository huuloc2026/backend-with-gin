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
	if m.Host == "" || m.User == "" || m.DBName == "" || m.Port == 0 {
		panic(fmt.Sprintf("Database config is invalid: %+v", m))
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable search_path=public",
		m.Host, m.User, m.Password, m.DBName, m.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	global.Logger.Info("Postgresql connected")
	global.Mdb = db
	// set pool
	setPool(global.Mdb)
	// Tạo schema nếu chưa có
	db.Exec("CREATE SCHEMA IF NOT EXISTS public;")
	// db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
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
		global.Logger.Error("Database migration failed")
		panic(fmt.Sprintf("Failed to migrate database tables: %v", err))
	}
}
