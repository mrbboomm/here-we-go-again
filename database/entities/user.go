package database

import (
	"log"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	TierID int
	Tier Tier `gorm:foreignKey:TierID`
}

func (User) TableName() string {
  return "users"
}

// mock for get-user
func GetUser(db *gorm.DB, id int) *User {
	var user User
	result := db.Preload("Tier").First(&user, id) // relations on pre-load

	if result.Error != nil {
		log.Fatalf("Error: Cound not get user: %v", result.Error)
	}

	return &user
}

func InsertUser(db *gorm.DB, obj *User) {
	result := db.Save(obj) // clean-insert
	if result.Error != nil {
		log.Fatalf("Error: Cound not get tier: %v", result.Error)
	}
}

func DeleteUser(db *gorm.DB, id int) {
	result := db.Unscoped().Delete(&User{}, id)
	if result.Error != nil {
		log.Fatalf("Error: Cound not get tier: %v", result.Error)
	}
}