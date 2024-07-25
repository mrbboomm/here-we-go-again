package database

import (
	"go-nf/config"
	"log"

	_ "github.com/godror/godror"

	oracle "github.com/godoes/gorm-oracle"

	"gorm.io/gorm"
)

var serviceName = "[Oracle]"

func Connect(cfg *config.OracleConfig) *gorm.DB {
	logPrefix := serviceName+"[Connect]"
	log.Println(logPrefix+": connect to oracle database...")
	dsn := oracle.BuildUrl(cfg.Url, cfg.Port, cfg.ServiceName, cfg.User, cfg.Password, cfg.Options)

	dialector := oracle.New(oracle.Config{
		DSN: dsn,
		RowNumberAliasForOracle11: "ROW_NUM",
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
			log.Fatalf("failed to connect to database")
	}
	log.Println(logPrefix+": connected.")

	// db.AutoMigrate(&entities.Tier{},&entities.User{})
	// return db
	// log.Println(logPrefix+": migrating...")

	// db2, err := sql.Open("godror", "godev_user/godev_pass@localhost:1521/godev")
	// if err != nil {
	// 		log.Fatalf(logPrefix+": migration failed, can not connect to database")
	// }

	// migrations := &migrate.FileMigrationSource{
	// 		Dir: "../../migrations/oracle",
	// }

	// exec, err := migrate.Exec(db2, "godror", migrations, migrate.Up)
	// if err != nil {
	// 		log.Fatalf(logPrefix+": migration failed: %v", err)
	// }

	// fmt.Println(exec)
	// log.Printf("Applied %d migrations!", n)
	return db
}