package oracle

import (
	"go-nf/config"
	migrations "go-nf/oracle/migrations"
	"log"

	oracle "github.com/godoes/gorm-oracle"
	"gorm.io/gorm"
)

var db *gorm.DB
var serviceName = "[Oracle]"

func Connect(cfg *config.OracleConfig) {
	logPrefix := serviceName+"[Connect]"
	dsn := oracle.BuildUrl(cfg.Url, cfg.Port, cfg.ServiceName, cfg.User, cfg.Password, cfg.Options)

	dialector := oracle.New(oracle.Config{
		DSN: dsn,
		RowNumberAliasForOracle11: "ROW_NUM",
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
			log.Fatalf("failed to connect to database")
	}
	log.Println(logPrefix+": connect success")

	Migrations(db)
}

func Migrations(DB *gorm.DB) {
	logPrefix := serviceName+"[Migration]"
	log.Println(logPrefix+": start migrate...")

	migrations.AddTierAndUserEntity(DB)
	
	log.Println(logPrefix+": migrate success.")
}