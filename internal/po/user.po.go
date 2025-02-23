package po

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User struct represents a user entity
type User struct {
	ID        uuid.UUID `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(100);not null" json:"name"`
	Email     string    `gorm:"column:email;type:varchar(100);unique;not null" json:"email"`
	Password  string    `gorm:"column:password;type:varchar(255);not null" json:"-"`
	CreatedAt time.Time `gorm:"column:create_at;tautoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	// ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	// Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	// Email     string    `gorm:"type:varchar(100);unique;not null" json:"email"`
	// Password  string    `gorm:"type:varchar(255);not null" json:"-"`
	// CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	// UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// BeforeCreate ensures UUID is generated before inserting a new user
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return nil
}
