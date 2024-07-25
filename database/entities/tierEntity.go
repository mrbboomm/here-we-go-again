package database

import (
	"log"

	"gorm.io/gorm"
)

type Lang struct {
	En string
	Th string
}

type Tier struct {
	gorm.Model
	Name  Lang `gorm:"serializer:json"`
}

func (Tier) TableName() string {
    return "tier"
}

// mock for get-tier
func GetTier(db *gorm.DB, id int) *Tier {
	var tier Tier
	result := db.First(&tier, id)

	if result.Error != nil {
		log.Fatalf("Error: Cound not get tier: %v", result.Error)
	}

	return &tier
}

func InsertTier(db *gorm.DB, obj *Tier) {
	result := db.Save(obj)
	if result.Error != nil {
		log.Fatalf("Error: Cound not get tier: %v", result.Error)
	}
}

func DeleteTier(db *gorm.DB, id int) {
	result := db.Unscoped().Delete(&Tier{}, id)
	if result.Error != nil {
		log.Fatalf("Error: Cound not get tier: %v", result.Error)
	}
}