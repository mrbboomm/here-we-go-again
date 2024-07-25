package database

import (
	"database/sql"
	"go-nf/config"
	"log"

	_ "github.com/godror/godror"

	oracle "github.com/godoes/gorm-oracle"
	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/gorm"
)

var serviceName = "[Oracle]"

func Connect(cfg *config.OracleConfig) {
	logPrefix := serviceName+"[Connect]"
	log.Println(logPrefix+": connect to oracle database...")
	dsn := oracle.BuildUrl(cfg.Url, cfg.Port, cfg.ServiceName, cfg.User, cfg.Password, cfg.Options)

	dialector := oracle.New(oracle.Config{
		DSN: dsn,
		RowNumberAliasForOracle11: "ROW_NUM",
	})

	_, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
			log.Fatalf("failed to connect to database")
	}
	log.Println(logPrefix+": connected.")

	log.Println(logPrefix+": migrating...")

	db, err := sql.Open("godror", "godev_user/godev_pass@localhost:1521/godev")
	if err != nil {
			log.Fatalf(logPrefix+": migration failed, can not connect to database")
	}

	migrations := &migrate.FileMigrationSource{
			Dir: "../..migrations/oracle",
	}

	n, err := migrate.Exec(db, "godror", migrations, migrate.Up)
	if err != nil {
			log.Fatalf(logPrefix+": migration failed: %v", err)
	}

	log.Printf("Applied %d migrations!", n)
}