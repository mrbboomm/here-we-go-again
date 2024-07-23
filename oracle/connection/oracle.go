package oracle

import (
	"fmt"
	migrations "go-nf/oracle/migrations"
	"log"

	oracle "github.com/godoes/gorm-oracle"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){
	options := map[string]string{
		"CONNECTION TIMEOUT": "90",
		"SSL":                "false",
	}
	url := oracle.BuildUrl("127.0.0.1", 1521, "godev", "godev_user", "godev_pass", options)
	dialector := oracle.New(oracle.Config{
		DSN:                     url,
		RowNumberAliasForOracle11: "ROW_NUM",
	})

	DB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
			log.Fatalf("failed to connect to database")
	}

	Migrations(DB)
}

func Migrations(DB *gorm.DB) {
	fmt.Println("migration start...")

	migrations.AddTierAndUserEntity(DB)
	
	fmt.Println("migration end.")
}