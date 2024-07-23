package oracle

import (
	"fmt"
	"go-nf/config"
	migrations "go-nf/oracle/migrations"
	"log"

	oracle "github.com/godoes/gorm-oracle"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){
	cfg := config.LocalOracleConfig()
	options := map[string]string{
		"CONNECTION TIMEOUT": "90",
		"SSL": "false",
	}

	dsn := oracle.BuildUrl(cfg.Url, cfg.Port, cfg.ServiceName, cfg.User, cfg.Password, options)

	dialector := oracle.New(oracle.Config{
		DSN: dsn,
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