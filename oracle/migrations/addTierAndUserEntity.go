package oracle

import (
	"fmt"
	entities "go-nf/oracle/entities"
	"log"

	"gorm.io/gorm"
)

func AddTierAndUserEntity(DB *gorm.DB) {
	logPrefix := "AddTierAndUserEntity"
	if !DB.Migrator().HasTable("TIER") || !DB.Migrator().HasTable("USERS") {
		fmt.Println("migrate table tier...")
		err := DB.Migrator().AutoMigrate(&entities.Tier{}, &entities.User{})
		if err != nil {
			log.Fatalf("[%v] migration fail: %v", logPrefix, err)	
		}
		fmt.Println("migrate table end.")
	}
}