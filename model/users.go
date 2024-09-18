package model

import (
	"time"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID              uint32    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Email           string    `gorm:"column:email;not null" json:"email"`
	CryptedPassword string    `gorm:"column:crypted_password;not null" json:"crypted_password"`
	Secret          string    `gorm:"column:secret" json:"secret"`
	Token           string    `gorm:"column:token" json:"token"`
	Status          int32     `gorm:"column:status" json:"status"`
	Role            string    `gorm:"column:role" json:"role"`
	CreatedAt       time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
