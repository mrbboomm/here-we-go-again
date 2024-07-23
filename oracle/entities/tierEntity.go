package oracle

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
func getTier(db *gorm.DB, id int) *Tier {
	var tier Tier
	result := db.First(&tier, id)

	if result.Error != nil {
		log.Fatalf("Error: Cound not get tier: %v", result.Error)
	}

	return &tier
}