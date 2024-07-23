package oracle

import (
	"log"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	TierID int
	Tier *Tier `gorm:foreignKey:TierID`
}

func (User) TableName() string {
  return "users"
}

// mock for get-user
func getUser(db *gorm.DB, id int) *User {
	var user User
	result := db.First(&user, id)

	if result.Error != nil {
		log.Fatalf("Error: Cound not get user: %v", result.Error)
	}

	return &user
}