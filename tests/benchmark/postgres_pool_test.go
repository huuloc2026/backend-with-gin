package benchmark

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	"github.com/huuloc2026/go-backend/global"

	"github.com/huuloc2026/go-backend/internal/initalize"
	"github.com/huuloc2026/go-backend/internal/po"
	"gorm.io/gorm"
)

// Setup function to initialize DB connection before benchmarking
func setupDB() *gorm.DB {

	if global.Mdb == nil {
		//todo:
		initalize.LoadConfig()
		initalize.InitPostgresql()
	}
	return global.Mdb

}

// Generate a random user for testing
func randomUser() po.User {
	return po.User{
		Name:     fmt.Sprintf("User%d", rand.Intn(100000)),
		Email:    fmt.Sprintf("user%d@example.com", rand.Intn(100000)),
		Password: "password123",
	}
}

// Benchmark Insert User
func BenchmarkInsertUser(b *testing.B) {
	db := setupDB()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		user := randomUser()
		if err := db.Create(&user).Error; err != nil {
			b.Errorf("Failed to insert user: %v", err)
		}
	}
}

// Benchmark Query User
func BenchmarkQueryUser(b *testing.B) {
	db := setupDB()
	var user po.User
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := db.First(&user).Error; err != nil && err != gorm.ErrRecordNotFound {
			b.Errorf("Failed to query user: %v", err)
		}
	}
}

// Benchmark Update User
func BenchmarkUpdateUser(b *testing.B) {
	db := setupDB()
	var user po.User
	db.First(&user) // Lấy một user bất kỳ để cập nhật
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := db.Model(&user).Update("name", "UpdatedUser"+strconv.Itoa(i)).Error; err != nil {
			b.Errorf("Failed to update user: %v", err)
		}
	}
}

// Benchmark Delete User
func BenchmarkDeleteUser(b *testing.B) {
	db := setupDB()
	var user po.User
	db.First(&user) // Lấy một user bất kỳ để xóa
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := db.Delete(&user).Error; err != nil {
			b.Errorf("Failed to delete user: %v", err)
		}
	}
}
